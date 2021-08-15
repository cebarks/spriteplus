package main

import (
	"flag"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/cebarks/spriteplus"
)

func main() {
	in := flag.String("input", "", "location to input image")
	height := flag.Int("height", 16, "height of each sprite")
	width := flag.Int("width", 16, "width of each sprite")

	flag.Parse()

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("couldn't get working directory: %v", err)
	}

	err = os.Mkdir(filepath.Join(pwd, "out"), 0744)
	// if err != nil {
	// 	log.Fatalf("couldn't make output directory: %T: %s", err, err)
	// }

	infile, err := os.Open(filepath.Join(pwd, *in))
	if err != nil {
		log.Fatalf("Couldn't open input file: %v", err)
	}
	defer infile.Close()

	img, err := png.Decode(infile)
	if err != nil {
		log.Fatalf("Couldn't decode input file: %v", err)
	}

	for x := 0; x < img.Bounds().Max.X / *width; x++ {
		for y := 0; y < img.Bounds().Max.Y / *height; y++ {
			sub, err := spriteplus.Subimage(img, x**width, y**height, x**width+*width, y**height+*height)
			if err != nil {
				log.Fatalln(err)
			}

			outfile, err := os.Create(filepath.Join(pwd, "out", fmt.Sprintf("%v-%v.png", x, y)))
			defer outfile.Close()
			if err != nil {
				log.Fatalln(err)
			}

			err = png.Encode(outfile, sub)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
