// plot renders a GameMap's MapMarkers as numbered circles on top of the
// map's image, so their (x, y) coordinates can be sanity-checked against
// the actual art instead of guessed blind.
//
// Usage:
//
//	go run ./cmd/plot -map blingdenstone
//	go run ./cmd/plot -map blingdenstone -crop 3500,350,4050,800 -out check.png
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"golang.org/x/image/webp"

	_ "modernc.org/sqlite"
)

type marker struct {
	i     int
	x, y  int
	title string
}

func main() {
	dbPath := flag.String("db", "oota.db", "path to the sqlite database")
	mapID := flag.String("map", "", "map_id to plot (required)")
	imagesDir := flag.String("images", "images", "directory containing map images (used if the img column is a relative path)")
	out := flag.String("out", "", "output PNG path (default <map_id>.png)")
	radius := flag.Int("radius", 18, "marker circle radius in source-image pixels")
	cropFlag := flag.String("crop", "", "optional crop rect \"x0,y0,x1,y1\" in source-image pixels, applied after plotting")
	maxDim := flag.Int("max-dim", 2000, "downscale output so its longest side is at most this many pixels (0 to disable)")
	flag.Parse()

	if *mapID == "" {
		log.Fatal("-map is required")
	}
	if *out == "" {
		*out = *mapID + ".png"
	}

	db, err := sql.Open("sqlite", *dbPath+"?_pragma=foreign_keys(1)")
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	var imgPath string
	if err := db.QueryRow(`SELECT img FROM GameMaps WHERE id = ?`, *mapID).Scan(&imgPath); err != nil {
		log.Fatalf("lookup map %q: %v", *mapID, err)
	}

	rows, err := db.Query(`SELECT i, x, y, title FROM MapMarkers WHERE map_id = ? ORDER BY i`, *mapID)
	if err != nil {
		log.Fatalf("query markers: %v", err)
	}
	defer rows.Close()

	var markers []marker
	for rows.Next() {
		var m marker
		if err := rows.Scan(&m.i, &m.x, &m.y, &m.title); err != nil {
			log.Fatalf("scan marker: %v", err)
		}
		markers = append(markers, m)
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("iterate markers: %v", err)
	}
	if len(markers) == 0 {
		log.Fatalf("no markers found for map %q", *mapID)
	}

	resolved := imgPath
	if !filepath.IsAbs(resolved) {
		resolved = filepath.Join(*imagesDir, filepath.Base(imgPath))
	}
	src, err := loadImage(resolved)
	if err != nil {
		log.Fatalf("load image %q: %v", resolved, err)
	}

	canvas := image.NewRGBA(src.Bounds())
	draw.Draw(canvas, canvas.Bounds(), src, src.Bounds().Min, draw.Src)

	for _, m := range markers {
		drawMarker(canvas, m, *radius)
	}

	final := image.Image(canvas)
	if *cropFlag != "" {
		r, err := parseRect(*cropFlag)
		if err != nil {
			log.Fatalf("bad -crop: %v", err)
		}
		final = cropImage(canvas, r)
	}
	if *maxDim > 0 {
		final = downscale(final, *maxDim)
	}

	f, err := os.Create(*out)
	if err != nil {
		log.Fatalf("create %q: %v", *out, err)
	}
	defer f.Close()
	if err := png.Encode(f, final); err != nil {
		log.Fatalf("encode png: %v", err)
	}
	fmt.Printf("plotted %d markers for %q -> %s\n", len(markers), *mapID, *out)
}

func loadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	switch strings.ToLower(filepath.Ext(path)) {
	case ".webp":
		return webp.Decode(f)
	case ".png":
		return png.Decode(f)
	default:
		img, _, err := image.Decode(f)
		return img, err
	}
}

func parseRect(s string) (image.Rectangle, error) {
	parts := strings.Split(s, ",")
	if len(parts) != 4 {
		return image.Rectangle{}, fmt.Errorf("expected \"x0,y0,x1,y1\", got %q", s)
	}
	vals := make([]int, 4)
	for i, p := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			return image.Rectangle{}, fmt.Errorf("parse %q: %w", p, err)
		}
		vals[i] = n
	}
	return image.Rect(vals[0], vals[1], vals[2], vals[3]), nil
}

func cropImage(img image.Image, r image.Rectangle) image.Image {
	r = r.Intersect(img.Bounds())
	out := image.NewRGBA(image.Rect(0, 0, r.Dx(), r.Dy()))
	draw.Draw(out, out.Bounds(), img, r.Min, draw.Src)
	return out
}

// downscale returns img resized (nearest-neighbor) so its longest side is
// at most maxDim pixels. It's a diagnostic image, not a deliverable, so
// nearest-neighbor is fine.
func downscale(img image.Image, maxDim int) image.Image {
	b := img.Bounds()
	w, h := b.Dx(), b.Dy()
	longest := w
	if h > longest {
		longest = h
	}
	if longest <= maxDim {
		return img
	}
	scale := float64(maxDim) / float64(longest)
	nw, nh := int(float64(w)*scale), int(float64(h)*scale)

	out := image.NewRGBA(image.Rect(0, 0, nw, nh))
	for y := 0; y < nh; y++ {
		sy := b.Min.Y + int(float64(y)/scale)
		for x := 0; x < nw; x++ {
			sx := b.Min.X + int(float64(x)/scale)
			out.Set(x, y, img.At(sx, sy))
		}
	}
	return out
}

var markerColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}

func drawMarker(canvas *image.RGBA, m marker, radius int) {
	drawCircle(canvas, m.x, m.y, radius, markerColor)
	label := fmt.Sprintf("%d %s", m.i, m.title)
	drawLabel(canvas, m.x+radius+6, m.y+radius/2, label)
}

// drawCircle draws a ring (not filled) of the given radius, ~4px thick.
func drawCircle(canvas *image.RGBA, cx, cy, radius int, col color.Color) {
	const thickness = 4
	b := canvas.Bounds()
	for dy := -radius - thickness; dy <= radius+thickness; dy++ {
		for dx := -radius - thickness; dx <= radius+thickness; dx++ {
			d2 := dx*dx + dy*dy
			if d2 < (radius-thickness)*(radius-thickness) || d2 > (radius+thickness)*(radius+thickness) {
				continue
			}
			x, y := cx+dx, cy+dy
			if x < b.Min.X || x >= b.Max.X || y < b.Min.Y || y >= b.Max.Y {
				continue
			}
			canvas.Set(x, y, col)
		}
	}
}

// drawLabel draws white-outlined black text so it reads on any background.
func drawLabel(canvas *image.RGBA, x, y int, text string) {
	face := basicfont.Face7x13
	pt := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}
	for _, off := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		drawText(canvas, face, pt.X+fixed.I(off[0]), pt.Y+fixed.I(off[1]), text, color.White)
	}
	drawText(canvas, face, pt.X, pt.Y, text, color.Black)
}

func drawText(canvas *image.RGBA, face font.Face, x, y fixed.Int26_6, text string, col color.Color) {
	d := &font.Drawer{
		Dst:  canvas,
		Src:  image.NewUniform(col),
		Face: face,
		Dot:  fixed.Point26_6{X: x, Y: y},
	}
	d.DrawString(text)
}
