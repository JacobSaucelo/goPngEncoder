package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Decode struct {
	data image.Image
}

func DecodeImage() {
	RawImage := NewDecode()
	RawImage.ImageToBits()
}

func NewDecode() Decode {
	ImageFile, ImageFileErr := os.Open("./output/output.png")
	if ImageFileErr != nil {
		fmt.Println("Error: Unable to find 'output.png' on output directory.")
		panic(ImageFileErr)
	}
	defer ImageFile.Close()

	LoadImage, LoadImageErr := png.Decode(ImageFile)
	if LoadImageErr != nil {
		fmt.Println("Error: Unable to load image.")
		panic(LoadImageErr)
	}

	return Decode{
		data: LoadImage,
	}
}

func (rawImage *Decode) ImageToBits() {
	bounds := rawImage.data.Bounds()
	bitString := ""

	// (1,1)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// fmt.Printf("(1,%d)\n", x)
			pixelColor := rawImage.data.At(x, y)
			rgba := color.RGBAModel.Convert(pixelColor).(color.RGBA)
			fmt.Printf("Pixel at (%d, %d): R:%d G:%d B:%d A:%d\n", x, y, rgba.R, rgba.G, rgba.B, rgba.A)

			if (rgba == color.RGBA{255, 255, 255, 255}) {
				fmt.Print(" white")
				bitString += "1"
			} else {
				fmt.Print(" black")
				bitString += "0"

			}
		}
	}

	fmt.Println("raw Image: ", bitString)
}
