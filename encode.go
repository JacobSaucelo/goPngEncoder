package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Encode struct {
	data []byte
}

func EncodeImage() {
	EncodeData := Encode{}

	textData, textDataErr := os.ReadFile("./files/input.txt")
	if textDataErr != nil {
		fmt.Println("ERROR: file 'input.text' doesnt exists in 'files' directory.")
		panic(textDataErr)
	}
	EncodeData.data = textData

	binaryData := EncodeData.BytesToBinary()
	EncodeData.CreateImageDir()
	EncodeData.CreateImage(binaryData)
}

func (rawData *Encode) BytesToBinary() string {
	binaryString := ""
	for _, data := range rawData.data {
		binaryString += fmt.Sprintf("%08b", data)
	}
	return binaryString
}

func (rawData *Encode) CreateImageDir() {
	FolderErr := os.MkdirAll("output", os.ModePerm)
	if FolderErr != nil {
		fmt.Println("ERROR: Unable to crate folder named 'output' on the root path.")
		panic(FolderErr)
	}

}

func (rawData *Encode) CreateImage(binaryData string) {
	width := len(binaryData)
	height := 1

	img := image.NewGray(image.Rect(0, 0, width, height))
	white := color.Gray{255}
	black := color.Gray{0}

	for index, character := range binaryData {
		if character == '1' {
			img.SetGray(index, 0, white)
		} else {
			img.SetGray(index, 0, black)
		}
	}

	imageFile, imageFErr := os.Create("./output/output.png")
	if imageFErr != nil {
		fmt.Println("ERROR: Theres a problem on creating a png file.")
		panic(imageFErr)
	}
	defer imageFile.Close()

	pngImageErr := png.Encode(imageFile, img)
	if pngImageErr != nil {
		fmt.Println("ERROR: Unable to encode png image.")
		panic(pngImageErr)
	}

	fmt.Println("Successfully created output.png")
}
