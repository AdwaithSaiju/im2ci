package main

import (
	"image"
	"math"
	"strings"
)

var chars = []byte("@%#*+=-:. ")

func luminance(r, g, b uint32) float64 {
	return 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
}

func ImgToAscii(img image.Image, invert bool) string {
	var builder strings.Builder
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	grid := make([][]float64, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]float64, width)
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(bounds.Min.X+x, bounds.Min.Y+y).RGBA()
			grid[y][x] = luminance(r, g, b)
		}
	}

	maxIdx := float64(len(chars) - 1)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			oldpixel := grid[y][x]

			idx := oldpixel / 65535.0 * maxIdx
			if invert {
				idx = maxIdx - idx
			}

			i := int(math.Round(idx))
			if i < 0 {
				i = 0
			}
			if i >= len(chars) {
				i = len(chars) - 1
			}

			newpixel := float64(i) / maxIdx * 65535.0
			if invert {
				newpixel = (maxIdx - float64(i)) / maxIdx * 65535.0
			}
			quantError := oldpixel - newpixel

			builder.WriteByte(chars[i])

			if x+1 < width {
				grid[y][x+1] += quantError * 7 / 16
			}
			if y+1 < height {
				if x > 0 {
					grid[y+1][x-1] += quantError * 3 / 16
				}
				grid[y+1][x] += quantError * 5 / 16
				if x+1 < width {
					grid[y+1][x+1] += quantError * 1 / 16
				}
			}
		}

		builder.WriteByte('\n')
	}

	return builder.String()
}
