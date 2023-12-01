package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

type Decode struct {
	data image.Image
}

func DecodeImage() {
	RawImage := NewDecode()
	RawImage.ImageToText()
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

func (rawImage *Decode) ImageToText() {
	bounds := rawImage.data.Bounds()
	bitString := ""

	// (1,1)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// fmt.Printf("(1,%d)\n", x)
			pixelColor := rawImage.data.At(x, y)
			rgba := color.RGBAModel.Convert(pixelColor).(color.RGBA)
			// fmt.Printf("Pixel at (%d, %d): R:%d G:%d B:%d A:%d\n", x, y, rgba.R, rgba.G, rgba.B, rgba.A)

			if (rgba == color.RGBA{255, 255, 255, 255}) {
				bitString += "1"
			} else {
				bitString += "0"
			}
		}
	}

	// fmt.Println("raw Image: ", bitString)
	RawBytes := binaryToBytes(bitString)
	CreateText(RawBytes)
}

func binaryToBytes(bitString string) []byte {
	// check if divisible by 8 - adds 1 index on length
	for len(bitString)%8 != 0 {
		bitString = "0" + bitString
	}

	bytes := make([]byte, len(bitString)/8)

	for i := 0; i < len(bitString); i += 8 {
		chunk := bitString[i : i+8]
		value, _ := strconv.ParseUint(chunk, 2, 8)

		bytes[i/8] = byte(value)
	}

	return bytes
}

func CreateText(bytes []byte) {
	FileError := os.WriteFile("./output/decoded.txt", bytes, 0644)
	if FileError != nil {
		fmt.Println("Error:", FileError)
	} else {
		fmt.Println("Text file 'output.txt' created successfully.")
	}
}
