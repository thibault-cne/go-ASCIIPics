package main

import (
	"flag"
	"fmt"
	"go-ASCIIPics/asciipics"
	"image"
	_ "image/png"
	"os"
)

var (
	density = flag.String("density", "nT#JCwfy325Fp6mqSghVd4EgXPGZbYkOA&8U$@KHDBWNMR0Q", "ASCII density string")
	imgPath = flag.String("image", "exemple.png", "Path to the image to display")
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

	asciipics.AsciiToConsole(img, dString)
}
