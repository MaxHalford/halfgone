package main

import (
	"math/rand"
	"time"
)

func main() {
	var (
		img, _ = LoadImage("images/penguin.jpg")
		rng    = rand.New(rand.NewSource(time.Now().UnixNano()))
	)

	// Grayscale conversion
	var gray = RGBAToGray(img)
	SaveImagePNG(gray, "images/grayscale.png")

	// Intensity inversion
	var inverted = InvertGray(gray)
	SaveImagePNG(inverted, "images/inverted_grayscale.png")

	// Threshold dithering
	var td = ThresholdDitherer{122}.apply(gray)
	SaveImagePNG(td, "images/threshold_dithering.png")

	// Random threshold dithering
	var rtd = RandomThresholdDitherer{100, rng}.apply(gray)
	SaveImagePNG(rtd, "images/random_threshold_dithering.png")

	// Importance sampling
	var is = ImportanceSampling{n: 2000, threshold: 100, rng: rng}.apply(gray)
	SaveImagePNG(is, "images/importance_sampling.png")

	// Bosch and Herman’s grid-based dithering
	var gd = GridDitherer{5, 3, 8, rng}.apply(gray)
	SaveImagePNG(gd, "images/grid_dithering.png")

	// Floyd-Steinberg dithering
	var fsd = FloydSteinbergDitherer{}.apply(gray)
	SaveImagePNG(fsd, "images/floyd_steinberg_dithering.png")
}
