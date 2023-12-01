package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
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

	bytes := binaryToBytes(bitString)

	// Write bytes to a text file
	err := writeBytesToFile(bytes, "outputTest.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Text file 'output.txt' created successfully.")
	}
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

func binaryToBytes(binaryString string) []byte {
	// Ensure the binary string length is a multiple of 8
	for len(binaryString)%8 != 0 {
		binaryString = "0" + binaryString
	}

	// Initialize a byte slice
	bytes := make([]byte, len(binaryString)/8)

	// Parse the binary string and populate the byte slice
	for i := 0; i < len(binaryString); i += 8 {
		// Extract each 8-bit chunk
		chunk := binaryString[i : i+8]

		// Parse the chunk as an integer
		value, _ := strconv.ParseUint(chunk, 2, 8)

		// Convert the integer to a byte and assign it to the byte slice
		bytes[i/8] = byte(value)
	}

	return bytes
}

func writeBytesToFile(bytes []byte, filename string) error {
	return ioutil.WriteFile(filename, bytes, 0644)
}
