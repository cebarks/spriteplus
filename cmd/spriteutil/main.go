package main

import (
	"flag"
	"image"
	"log"
	"os"
	"path/filepath"

	"github.com/cebarks/spriteplus"
)

func main() {
	in := flag.String("input", "", "location to input image")
	height := flag.Int("height", 16, "height of each sprite")
	width := flag.Int("width", 16, "width of each sprite")

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("couldn't get working directory: %v", err)
	}

	err = os.Mkdir(filepath.Join(pwd, "out"), 0744)
	if err != nil {
		log.Fatalf("couldn't make output directory: %v", err)
	}

	file, err := os.Open(*in)
	if err != nil {
		log.Fatalf("Couldn't make output directory: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Couldn't make output directory: %v", err)
	}

	for x := 0; x < img.Bounds().Max.X / *width; x++ {
		for y := 0; y < img.Bounds().Max.Y / *height; y++ {
			sub, err := spriteplus.Subimage(img)
			if err != nil {
				return nil, err
			}
		}
	}

}
