package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"
)

// A Ditherer convert a grayscale image with intensities going from 0 (black) to
// 255 (white) into a black and white image.
type Ditherer interface {
	apply(gray *image.Gray) *image.Gray
}

// ThresholdDitherer converts each pixel in a grayscale image to black or white
// depending on the intensity of the pixel.
type ThresholdDitherer struct {
	threshold uint8
}

func (dith ThresholdDitherer) apply(gray *image.Gray) *image.Gray {
	var (
		bounds   = gray.Bounds()
		dithered = image.NewGray(bounds)
		width    = bounds.Dx()
		height   = bounds.Dy()
	)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if gray.GrayAt(i, j).Y > dith.threshold {
				dithered.SetGray(i, j, color.Gray{255})
			} else {
				dithered.SetGray(i, j, color.Gray{0})
			}
		}
	}
	return dithered
}

// GridDitherer implements Bosch and Herman's grid-based method.
type GridDitherer struct {
	k     int     // Size in pixels of a side of a cell
	alpha float64 // Minimum desired number of points in a cell
	beta  float64 // Maximum desired number of points in a cell
	rng   *rand.Rand
}

func (dith GridDitherer) apply(gray *image.Gray) *image.Gray {
	var (
		bounds   = gray.Bounds()
		dithered = colorWhite(image.NewGray(bounds))
		width    = bounds.Dx()
		height   = bounds.Dy()
	)
	for i := 0; i < width; i += dith.k {
		for j := 0; j < height; j += dith.k {
			var (
				cell = rgbaToGray(gray.SubImage(image.Rect(i, j, i+dith.k, j+dith.k)))
				mu   = avgIntensity(cell)                                 // Mean grayscale value of the cell
				n    = math.Pow(dith.beta-mu*dith.beta, 2)/3 - dith.alpha // Number of points to sample
			)
			for k := 0; k < int(n); k++ {
				// Sample n random points in belonging to the cell
				var (
					x = randInt(i, min(i+dith.k, width), dith.rng)
					y = randInt(j, min(j+dith.k, height), dith.rng)
				)
				dithered.Set(x, y, color.Gray{0})
			}
		}
	}
	return dithered
}
