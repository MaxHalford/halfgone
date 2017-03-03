package main

import (
	_ "image/jpeg"
	"math/rand"
	"time"

	"github.com/MaxHalford/halfgone"
)

func main() {
	var (
		img, _ = halfgone.LoadImage("images/penguin.jpg")
		rng    = rand.New(rand.NewSource(time.Now().UnixNano()))
	)

	// Grayscale conversion
	var gray = halfgone.RGBAToGray(img)
	halfgone.SaveImagePNG(gray, "images/grayscale.png")

	// Intensity inversion
	var inverted = halfgone.InvertGray(gray)
	halfgone.SaveImagePNG(inverted, "images/inverted_grayscale.png")

	// Threshold dithering
	var td = halfgone.ThresholdDitherer{Threshold: 122}.Apply(gray)
	halfgone.SaveImagePNG(td, "images/threshold_dithering.png")

	// Random threshold dithering
	var rtd = halfgone.RandomThresholdDitherer{MaxThreshold: 100, RNG: rng}.Apply(gray)
	halfgone.SaveImagePNG(rtd, "images/random_threshold_dithering.png")

	// Importance sampling
	var is = halfgone.ImportanceSampling{N: 2000, Threshold: 100, RNG: rng}.Apply(gray)
	halfgone.SaveImagePNG(is, "images/importance_sampling.png")

	// Bosch and Herman’s grid-based dithering
	var gd = halfgone.GridDitherer{K: 5, Alpha: 3, Beta: 8, RNG: rng}.Apply(gray)
	halfgone.SaveImagePNG(gd, "images/grid_dithering.png")

	// Floyd-Steinberg dithering
	var fsd = halfgone.FloydSteinbergDitherer{}.Apply(gray)
	halfgone.SaveImagePNG(fsd, "images/floyd_steinberg_dithering.png")
}
