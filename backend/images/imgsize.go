//go:build ignore

package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"

	_ "golang.org/x/image/webp"
)

func main() {
	dirname := "."

	files, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		r, err := os.Open(filepath.Join(dirname, f.Name()))
		if err != nil {
			log.Printf("%s: open err: %v", f.Name(), err)
			continue
		}
		cfg, _, err := image.DecodeConfig(r)
		if err != nil {
			log.Printf("%s: decode err: %v", f.Name(), err)
			r.Close()
			continue
		}
		fmt.Printf("%s: %dx%d\n", f.Name(), cfg.Width, cfg.Height)
		r.Close()
	}
}
