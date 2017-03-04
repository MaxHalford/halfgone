package halfgone

import (
	"image"
	"image/color"
)

// ImageToGray converts an image.Image into an image.Gray.
func ImageToGray(img image.Image) *image.Gray {
	var (
		bounds = img.Bounds()
		gray   = image.NewGray(bounds)
	)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			gray.Set(x, y, img.At(x, y))
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

// makeGray generates a new image.Gray with each pixel being of a given
// intensity.
func makeGray(bounds image.Rectangle, intensity uint8) *image.Gray {
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
		bounds  = gray.Bounds()
		reverse = makeGray(bounds, 0)
	)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			reverse.SetGray(x, y, color.Gray{255 - gray.GrayAt(x, y).Y})
		}
	}
	return reverse
}
