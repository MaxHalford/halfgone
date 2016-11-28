package main

import (
	"image"
	"image/color"
)

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
func avgIntensity(gray *image.Gray) float64 {
	var sum float64
	for _, pix := range gray.Pix {
		sum += float64(pix)
	}
	return sum / float64(len(gray.Pix)*256)
}

func newWhite(bounds image.Rectangle) *image.Gray {
	var white = image.NewGray(bounds)
	for i := range white.Pix {
		white.Pix[i] = 255
	}
	return white
}

func blackOrWhite(g color.Gray) color.Gray {
	if g.Y < 123 {
		return color.Gray{0} // Black
	}
	return color.Gray{255} // White
}

func copyGray(gray *image.Gray) *image.Gray {
	var clone = image.NewGray(gray.Bounds())
	copy(clone.Pix, gray.Pix)
	return clone
}
