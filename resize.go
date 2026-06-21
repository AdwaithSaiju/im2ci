package main

import (
	"image"
	"image/color"
)

func ResizeImage(img image.Image, newWidth int) image.Image {
	bounds := img.Bounds()
	oldWidth := bounds.Dx()
	oldHeight := bounds.Dy()

	newHeight := oldHeight * newWidth / oldWidth
	newHeight /= 2

	if newHeight < 1 {
		newHeight = 1
	}

	resized := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	xRatio := float64(oldWidth) / float64(newWidth)
	yRatio := float64(oldHeight) / float64(newHeight)

	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			srcX := int(float64(x) * xRatio)
			srcY := int(float64(y) * yRatio)

			c := img.At(bounds.Min.X+srcX, bounds.Min.Y+srcY)
			resized.Set(x, y, color.RGBAModel.Convert(c))
		}
	}

	return resized
}
