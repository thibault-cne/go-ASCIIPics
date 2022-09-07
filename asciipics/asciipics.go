package asciipics

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"
)

var (
	r = 0.2126
	g = 0.7152
	b = 0.0722
)

func CloneAsRGBA(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)
	draw.Draw(newImg, bounds, img, bounds.Min, draw.Src)

	return newImg
}

func AsciiToConsole(img image.Image, dString string) {
	dLen := float64(len(dString)) - 1

	src := CloneAsRGBA(img)

	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pix := src.At(x, y).(color.RGBA)
			gray := float64(pix.R)*r + float64(pix.G)*g + float64(pix.B)*b + 0.5

			char := dString[int(gray*dLen/255)]
			asciiChar := fmt.Sprintf("%c%c", char, char)

			fmt.Printf("%s", asciiChar)
		}
		fmt.Printf("\n")
	}
}

func AsciiToFile(img image.Image, dString string, filename string) {
	dLen := float64(len(dString)) - 1

	// Create a new file
	outFilename := fmt.Sprintf("%s_%s.txt", "output", filename)
	f, _ := os.Create(outFilename)
	defer f.Close()

	src := CloneAsRGBA(img)

	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pix := src.At(x, y).(color.RGBA)
			gray := float64(pix.R)*r + float64(pix.G)*g + float64(pix.B)*b + 0.5

			char := dString[int(gray*dLen/255)]
			asciiChar := fmt.Sprintf("%c%c", char, char)

			f.WriteString(asciiChar)
		}
		f.WriteString("\n")
	}
}
