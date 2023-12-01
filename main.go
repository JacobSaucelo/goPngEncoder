package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// EncodeImage()
	// DecodeImage()

	mainMenu()
}

func mainMenu() {
	clearScreen()
	fmt.Println("pick an option: ")

	reader := bufio.NewReader(os.Stdin)

	options, _ := reader.ReadString('\n')

	switch strings.TrimSpace(options) {
	case "encode":
		fmt.Println("you picked encode")
	case "decode":
		fmt.Println("you picked decode")
	default:
		mainMenu()
	}
}

// go run main.go encode.go decode.go
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
