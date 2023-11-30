package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {

	existingImageFile, err := os.Open("output.png")
	if err != nil {
		panic(err)
	}
	defer existingImageFile.Close()

	imageData, imageType, imageErr := image.Decode(existingImageFile)
	if imageErr != nil {
		panic(imageErr)
	}
	fmt.Println(imageData)
	fmt.Println(imageType)

	existingImageFile.Seek(0, 0)

	loadedImage, loadedErr := png.Decode(existingImageFile)
	if loadedErr != nil {
		panic(loadedErr)
	}

	readAndPrintPixels(loadedImage)
}

func readAndPrintPixels(img image.Image) {
	bounds := img.Bounds()
	bitString := ""

	fmt.Printf("Image Size: %dx%d\n", bounds.Dx(), bounds.Dy())

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			pixelColor := img.At(x, y)

			rgba := color.RGBAModel.Convert(pixelColor).(color.RGBA)

			fmt.Printf("Pixel at (%d, %d): R:%d G:%d B:%d A:%d\n", x, y, rgba.R, rgba.G, rgba.B, rgba.A)

			if rgba.R == 255 && rgba.G == 255 && rgba.B == 255 {
				fmt.Print(" white")
				bitString += "1"
			} else {
				fmt.Print(" black")
				bitString += "0"
			}

		}
	}

	fmt.Println("bit string: ", bitString)
}

func imageToBits(img image.Image) string {
	bounds := img.Bounds()
	bits := ""

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			pixelColor := color.GrayModel.Convert(img.At(x, y)).(color.Gray)

			bit := "0"
			if pixelColor.Y > 128 {
				bit = "1"
			}

			bits += bit
		}
		bits += "\n"
	}

	return bits
}
