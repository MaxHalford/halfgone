package main

import (
	"image"
	_ "image/jpeg"
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

func randInt(min, max int, rng *rand.Rand) int {
	return rng.Intn(max-min) + min
}

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
