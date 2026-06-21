package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	invert := flag.Bool("i", false, "invert for dark/transparent terminals")
	width := flag.Int("w", 100, "output width in characters")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: im2ci [-i] [-w width] <image>")
		os.Exit(1)
	}

	img, err := LoadImage(flag.Arg(0))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Print(ImgToAscii(ResizeImage(img, *width), *invert))
}
