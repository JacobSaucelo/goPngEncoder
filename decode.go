package main

import (
	"fmt"
	"image"
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
	fmt.Println("raw Image: ", rawImage)
}
