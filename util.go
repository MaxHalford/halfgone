package halfgone

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

// LoadImage reads and loads an image from a file path.
func LoadImage(path string) (image.Image, error) {
	infile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer infile.Close()
	img, _, err := image.Decode(infile)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// SaveImagePNG save an image to a PNG file.
func SaveImagePNG(img image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	png.Encode(f, img)
	return nil
}

// min returns the smallest of two ints.
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// Convert a uint16 to a int8 taking into account bounds.
func i16ToUI8(x int16) uint8 {
	switch {
	case x < 1:
		return uint8(0)
	case x > 254:
		return uint8(255)
	}
	return uint8(x)
}

// randInt generates a random int in range [min, max).
func randInt(min, max int, rng *rand.Rand) int {
	return rng.Intn(max-min) + min
}

// randPoint generates a point with random coordinates from some given bounds.
func randPoint(bounds image.Rectangle, rng *rand.Rand) image.Point {
	return image.Point{
		X: randInt(bounds.Min.X, bounds.Max.X, rng),
		Y: randInt(bounds.Min.Y, bounds.Max.Y, rng),
	}
}

// randPoints generates n points with random coordinates from some given bounds.
func randPoints(n int, bounds image.Rectangle, rng *rand.Rand) []image.Point {
	var points = make([]image.Point, n)
	for i := range points {
		points[i] = randPoint(bounds, rng)
	}
	return points
}

// randColor generates a random non-transparent color.
func randColor() color.Color {
	return color.NRGBA{
		uint8(rand.Intn(256)),
		uint8(rand.Intn(256)),
		uint8(rand.Intn(256)),
		255,
	}
}
