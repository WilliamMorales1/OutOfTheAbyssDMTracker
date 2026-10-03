//go:build e2e

package e2e

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

const baseURL = "http://localhost:8080"

func TestBrowserUI(t *testing.T) {
	chromium, err := exec.LookPath("chromium")
	if err != nil {
		t.Skip("chromium is not installed")
	}

	backendDir, err := filepath.Abs("..")
	if err != nil {
		t.Fatal(err)
	}
	binary := filepath.Join(t.TempDir(), "oota")
	build := exec.Command("go", "build", "-o", binary, "./cmd/oota")
	build.Dir = backendDir
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build server: %v\n%s", err, output)
	}

	server := exec.Command(binary)
	server.Dir = backendDir
	if err := server.Start(); err != nil {
		t.Fatalf("start server: %v", err)
	}
	t.Cleanup(func() {
		_ = server.Process.Kill()
		_ = server.Wait()
	})
	waitForServer(t)

	allocatorOptions := append([]chromedp.ExecAllocatorOption{}, chromedp.DefaultExecAllocatorOptions[:]...)
	allocatorOptions = append(allocatorOptions,
		chromedp.ExecPath(chromium),
		chromedp.Headless,
		chromedp.DisableGPU,
		chromedp.Flag("disable-dev-shm-usage", true),
	)
	allocator, cancelAllocator := chromedp.NewExecAllocator(context.Background(), allocatorOptions...)
	defer cancelAllocator()
	browser, cancelBrowser := chromedp.NewContext(allocator)
	defer cancelBrowser()
	browser, cancelTimeout := context.WithTimeout(browser, 90*time.Second)
	defer cancelTimeout()

	var heading string
	if err := chromedp.Run(browser,
		chromedp.EmulateViewport(1440, 1000),
		chromedp.Navigate(baseURL),
		chromedp.WaitVisible("h1"),
		chromedp.Text("h1", &heading),
	); err != nil {
		t.Fatalf("load homepage: %v", err)
	}
	if strings.TrimSpace(heading) != "Out of the Abyss" {
		t.Fatalf("unexpected heading: %q", heading)
	}

	tabs := []struct {
		path   string
		marker string
	}{
		{"sessions", "Sessions"},
		{"notes", "Select a note"},
		{"monsters", "Monsters"},
		{"spells", "Spells"},
		{"maps", "Maps"},
		{"initiative", "Initiative"},
		{"soundboard", "Soundboard"},
		{"chat", "Ask Agent"},
		{"search", "Lore Search"},
		{"references", "References"},
	}

	for _, tab := range tabs {
		t.Run(tab.path, func(t *testing.T) {
			selector := fmt.Sprintf(`button[data-path="%s"]`, tab.path)
			if err := chromedp.Run(browser, chromedp.Click(selector)); err != nil {
				t.Fatalf("click tab: %v", err)
			}
			waitForPanel(t, browser, selector, tab.marker)
			if tab.path == "search" {
				waitForSearchResults(t, browser)
			}
			assertScreenshot(t, browser, tab.path)
		})
	}

	// Notes must use local vendored Monaco assets, never a CDN.
	var editorAvailable bool
	notesSelector := `button[data-path="notes"]`
	if err := chromedp.Run(browser, chromedp.Click(notesSelector)); err != nil {
		t.Fatalf("open notes panel: %v", err)
	}
	waitForPanel(t, browser, notesSelector, "Select a note")
	if err := chromedp.Run(browser,
		chromedp.Evaluate(`Boolean(window.monaco || document.querySelector('textarea'))`, &editorAvailable),
	); err != nil {
		t.Fatalf("check notes editor: %v", err)
	}
	if !editorAvailable {
		t.Fatal("notes editor did not load")
	}

	var previewAvailable bool
	if err := chromedp.Run(browser, chromedp.Evaluate(`(() => {
		const button = [...document.querySelectorAll('button')].find((item) => item.textContent === 'Preview');
		if (!button) return false;
		button.click();
		return true;
	})()`, &previewAvailable)); err != nil {
		t.Fatalf("open notes preview: %v", err)
	}
	if !previewAvailable {
		t.Fatal("notes Preview button missing")
	}
	waitForVisible(t, browser, ".markdown-preview:not(.hidden)")
	assertScreenshot(t, browser, "notes-preview")
}

func assertScreenshot(t *testing.T, browser context.Context, name string) {
	t.Helper()
	var actualPNG []byte
	if err := chromedp.Run(browser, chromedp.CaptureScreenshot(&actualPNG)); err != nil {
		t.Fatalf("capture %s screenshot: %v", name, err)
	}

	_, sourceFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("locate screenshot test directory")
	}
	goldenPath := filepath.Join(filepath.Dir(sourceFile), "testdata", "screenshots", name+".png")
	if os.Getenv("UPDATE_GOLDENS") == "1" {
		if err := os.MkdirAll(filepath.Dir(goldenPath), 0o755); err != nil {
			t.Fatalf("create screenshot directory: %v", err)
		}
		if err := os.WriteFile(goldenPath, actualPNG, 0o644); err != nil {
			t.Fatalf("write %s screenshot: %v", name, err)
		}
		return
	}

	goldenPNG, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("read %s golden: %v; run UPDATE_GOLDENS=1 make test-ui", name, err)
	}
	actual, err := png.Decode(bytes.NewReader(actualPNG))
	if err != nil {
		t.Fatalf("decode %s screenshot: %v", name, err)
	}
	golden, err := png.Decode(bytes.NewReader(goldenPNG))
	if err != nil {
		t.Fatalf("decode %s golden: %v", name, err)
	}
	if actual.Bounds() != golden.Bounds() {
		writeScreenshotFailure(t, name, actualPNG)
		t.Fatalf("%s screenshot dimensions changed: got %v, want %v", name, actual.Bounds(), golden.Bounds())
	}

	bounds := actual.Bounds()
	diff := image.NewRGBA(bounds)
	mismatches := 0
	const channelThreshold = 10
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			ar, ag, ab, _ := actual.At(x, y).RGBA()
			gr, gg, gb, _ := golden.At(x, y).RGBA()
			different := absInt(int(ar)-int(gr)) > channelThreshold*257 ||
				absInt(int(ag)-int(gg)) > channelThreshold*257 ||
				absInt(int(ab)-int(gb)) > channelThreshold*257
			if different {
				mismatches++
				diff.Set(x, y, color.RGBA{R: 255, A: 255})
			} else {
				diff.Set(x, y, color.RGBA{R: uint8(ar >> 8), G: uint8(ag >> 8), B: uint8(ab >> 8), A: 255})
			}
		}
	}

	allowed := maxInt(20, bounds.Dx()*bounds.Dy()/500)
	if mismatches > allowed {
		artifactDir := t.TempDir()
		actualPath, diffPath := filepath.Join(artifactDir, "actual.png"), filepath.Join(artifactDir, "diff.png")
		writePNG(t, actualPath, actual)
		writePNG(t, diffPath, diff)
		t.Fatalf("%s screenshot changed: %d/%d pixels differ; actual=%s diff=%s", name, mismatches, bounds.Dx()*bounds.Dy(), actualPath, diffPath)
	}
}

func writeScreenshotFailure(t *testing.T, name string, actualPNG []byte) {
	t.Helper()
	path := filepath.Join(t.TempDir(), name+"-actual.png")
	if err := os.WriteFile(path, actualPNG, 0o644); err != nil {
		t.Logf("write screenshot failure: %v", err)
		return
	}
	t.Logf("actual screenshot: %s", path)
}

func writePNG(t *testing.T, path string, picture image.Image) {
	t.Helper()
	file, err := os.Create(path)
	if err != nil {
		t.Logf("create screenshot artifact: %v", err)
		return
	}
	defer file.Close()
	if err := png.Encode(file, picture); err != nil {
		t.Logf("encode screenshot artifact: %v", err)
	}
}

func absInt(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func waitForServer(t *testing.T) {
	t.Helper()
	client := &http.Client{Timeout: 500 * time.Millisecond}
	deadline := time.Now().Add(15 * time.Second)
	for time.Now().Before(deadline) {
		response, err := client.Get(baseURL + "/")
		if err == nil {
			response.Body.Close()
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	t.Fatal("server did not become ready")
}

func waitForPanel(t *testing.T, browser context.Context, selector, marker string) {
	t.Helper()
	deadline := time.Now().Add(15 * time.Second)
	var lastState struct {
		Active bool   `json:"active"`
		Body   string `json:"body"`
		Error  string `json:"error"`
	}
	for time.Now().Before(deadline) {
		var state struct {
			Active bool   `json:"active"`
			Body   string `json:"body"`
			Error  string `json:"error"`
		}
		err := chromedp.Run(browser, chromedp.Evaluate(fmt.Sprintf(`({
			active: document.querySelector(%q)?.classList.contains('active') ?? false,
			body: document.body.innerText,
			error: document.querySelector('.text-red-500')?.textContent ?? '',
		})`, selector), &state))
		lastState = state
		if err == nil && state.Active && !strings.Contains(state.Body, "Loading...") {
			if state.Error != "" {
				t.Fatalf("panel error: %s", state.Error)
			}
			if !strings.Contains(state.Body, marker) {
				t.Fatalf("panel missing %q", marker)
			}
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	t.Fatalf("panel %q did not become ready: active=%v error=%q body=%q", selector, lastState.Active, lastState.Error, lastState.Body)
}

func waitForSearchResults(t *testing.T, browser context.Context) {
	t.Helper()
	deadline := time.Now().Add(15 * time.Second)
	for time.Now().Before(deadline) {
		var ready bool
		if err := chromedp.Run(browser, chromedp.Evaluate(`Boolean(document.querySelector('.card') || document.body.innerText.includes('No results found.'))`, &ready)); err == nil && ready {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	t.Fatal("search results did not become ready")
}

func waitForVisible(t *testing.T, browser context.Context, selector string) {
	t.Helper()
	deadline := time.Now().Add(15 * time.Second)
	for time.Now().Before(deadline) {
		var visible bool
		if err := chromedp.Run(browser, chromedp.Evaluate(fmt.Sprintf(`(() => {
			const element = document.querySelector(%q);
			return Boolean(element && !element.classList.contains('hidden'));
		})()`, selector), &visible)); err == nil && visible {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	t.Fatalf("element %q did not become visible", selector)
}
