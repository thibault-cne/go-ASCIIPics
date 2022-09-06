package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/png"
	"os"
)

var (
	density = flag.String("density", "nT#JCwfy325Fp6mqSghVd4EgXPGZbYkOA&8U$@KHDBWNMR0Q", "ASCII density string")
	imgPath = flag.String("image", "exemple.png", "Path to the image to display")
	r       = 0.2126
	g       = 0.7152
	b       = 0.0722
)

func main() {
	// Parse flags to retrieve flags value
	flag.Parse()

	f, err := os.Open(*imgPath)

	if err != nil {
		fmt.Printf("An error occured while opening the file : %s\n", err.Error())
		return
	}

	defer f.Close()
	img, _, err := image.Decode(f)

	if err != nil {
		fmt.Printf("An error occured while decoding the image : %s\n", err.Error())
		return
	}

	dString := *density
	dLen := float64(len(dString)) - 1

	src := CloneAsRGBA(img)

	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pix := src.At(x, y).(color.RGBA)
			gray := float64(pix.R)*r + float64(pix.G)*g + float64(pix.B)*b + 0.5

			fmt.Printf("%c", dString[int(gray*dLen/255)])
		}
		fmt.Printf("\n")
	}
}

func CloneAsRGBA(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)
	draw.Draw(newImg, bounds, img, bounds.Min, draw.Src)

	return newImg
}
