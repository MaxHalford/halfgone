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
