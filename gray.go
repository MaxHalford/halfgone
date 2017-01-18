package main

import (
	"image"
	"image/color"
)

// RgbaToGray create a new grayscale image from an rgba image.
func RgbaToGray(img image.Image) *image.Gray {
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

// MakeGray generates a new image.Gray with each pixel being of a given
// intensity.
func MakeGray(bounds image.Rectangle, intensity uint8) *image.Gray {
	var gray = image.NewGray(bounds)
	for i := range gray.Pix {
		gray.Pix[i] = intensity
	}
	return gray
}

func copyGray(gray *image.Gray) *image.Gray {
	var clone = image.NewGray(gray.Bounds())
	copy(clone.Pix, gray.Pix)
	return clone
}

// InvertGray inverses the scale of a grayscale image by substracting each
// intensity to 255. Thus 0 (black) becomes 255 (white) and vice versa.
func InvertGray(gray *image.Gray) *image.Gray {
	var (
		reverse = copyGray(gray)
		bounds  = reverse.Bounds()
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			reverse.SetGray(x, y, color.Gray{255 - reverse.GrayAt(x, y).Y})
		}
	}
	return reverse
}

// ExtractPoints extracts a list of points for which the intensity lies in a
// given intensity range.
func ExtractPoints(gray *image.Gray, min, max uint8) (points []image.Point) {
	var bounds = gray.Bounds()
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var intensity = gray.GrayAt(x, y).Y
			if intensity >= min && intensity <= max {
				points = append(points, image.Point{x, y})
			}
		}
	}
	return
}
