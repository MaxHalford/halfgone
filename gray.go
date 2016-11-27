package main

import "image"

// Create a new grayscale image from an rgba image.
func rgbaToGray(img image.Image) *image.Gray {
	var (
		bounds = img.Bounds()
		gray   = image.NewGray(bounds)
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var rgba = img.At(x, y)
			gray.Set(x, y, rgba)
		}
	}
	return gray
}

// Calculate the average intensity of a gray image.
// TODO: make this faster
func avgIntensity(gray *image.Gray) float64 {
	var sum float64
	for _, pix := range gray.Pix {
		sum += float64(pix)
	}
	return sum / float64(len(gray.Pix)*256)
}

func colorWhite(gray *image.Gray) *image.Gray {
	for i := range gray.Pix {
		gray.Pix[i] = 255
	}
	return gray
}
