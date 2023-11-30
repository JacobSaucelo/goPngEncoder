package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// Read the text file
	textData, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fileBits := bytesToBinary(textData)

	width := len(fileBits)
	height := 1
	img := image.NewGray(image.Rect(0, 0, width, height))

	white := color.Gray{255}
	black := color.Gray{0}

	for i, character := range fileBits {
		if character == '1' {
			img.SetGray(i, 0, white)
			// fmt.Printf("%c white \n", character)
		} else if character == '0' {
			img.SetGray(i, 0, black)
			// fmt.Printf("%c black \n", character)
		}
	}

	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		panic(err)
	}

	fmt.Println("Image saved to output.png")
}

func bytesToBinary(data []byte) string {
	binaryString := ""
	for _, b := range data {
		// Convert each byte to binary representation
		binaryString += fmt.Sprintf("%08b", b)
	}
	return binaryString
}
