package main

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"math"
	"math/rand"
	"os"
)

// Read and load an image from a file path.
func loadImage(filepath string) (image.Image, error) {
	infile, err := os.Open(filepath)
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

// randColor generates a random non-transparent color.
func randColor() color.Color {
	return color.NRGBA{
		uint8(rand.Intn(256)),
		uint8(rand.Intn(256)),
		uint8(rand.Intn(256)),
		255,
	}
}

// square returns the square of an integer.
func square(x int) int {
	return x * x
}

// pointDistance calculates the L2 distance between two points.
func pointDistance(a, b image.Point) float64 {
	return math.Pow(float64(square(b.X-a.X)+square(b.Y-a.Y)), 0.5)
}
