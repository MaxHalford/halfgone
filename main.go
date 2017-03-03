package main

import (
	"math/rand"
	"time"
)

func main() {
	var (
		img, _ = LoadImage("screenshots/penguin.jpg")
		rng    = rand.New(rand.NewSource(time.Now().UnixNano()))
	)

	// Grayscale conversion
	var gray = RGBAToGray(img)
	SaveImagePNG(gray, "screenshots/grayscale.png")

	// Intensity inversion
	var inverted = InvertGray(gray)
	SaveImagePNG(inverted, "screenshots/inverted_grayscale.png")

	// Threshold dithering
	var td = ThresholdDitherer{122}.apply(gray)
	SaveImagePNG(td, "screenshots/threshold_dithering.png")

	// Random threshold dithering
	var rtd = RandomThresholdDitherer{100, rng}.apply(gray)
	SaveImagePNG(rtd, "screenshots/random_threshold_dithering.png")

	// Bosch and Herman’s grid-based dithering
	//var gd = GridDitherer{5, 3, 8, rng}.apply(gray)
	//SaveImagePNG(gd, "screenshots/grid_dithering.png")

	// Floyd-Steinberg dithering
	var fsd = FloydSteinbergDitherer{}.apply(gray)
	SaveImagePNG(fsd, "screenshots/floyd_steinberg_dithering.png")

	// Importance sampling
	var is = ImportanceSampling{n: 2000, threshold: 100, rng: rng}.apply(gray)
	SaveImagePNG(is, "screenshots/importance_sampling.png")
}
