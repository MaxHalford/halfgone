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
			dithered.SetGray(i, j, blackOrWhite(gray.GrayAt(i, j)))
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
		dithered = newWhite(bounds)
		width    = bounds.Dx()
		height   = bounds.Dy()
	)
	for i := 0; i < width; i += dith.k {
		for j := 0; j < height; j += dith.k {
			var (
				cell = rgbaToGray(gray.SubImage(image.Rect(i, j, i+dith.k, j+dith.k)))
				mu   = avgIntensity(cell)                // Mean grayscale value of the cell
				n    = math.Pow((1-mu)*dith.beta, 2) / 3 // Number of points to sample
			)
			if n < dith.alpha {
				n = 0
			}
			for k := 0; k < int(n); k++ {
				// Sample n random points in belonging to the cell
				var (
					x = randInt(i, min(i+dith.k, width), dith.rng)
					y = randInt(j, min(j+dith.k, height), dith.rng)
				)
				dithered.SetGray(x, y, color.Gray{0})
			}
		}
	}
	return dithered
}

// FloydSteinbergDitherer implements the Floyd-Steingberg algorithm, which is a
// variation of the error diffusion method.
type FloydSteinbergDitherer struct{}

func (dith FloydSteinbergDitherer) apply(gray *image.Gray) *image.Gray {
	var (
		bounds   = gray.Bounds()
		width    = bounds.Dx()
		height   = bounds.Dy()
		dithered = copyGray(gray)
	)
	for j := 0; j < height; j++ { // Top to bottom
		for i := 0; i < width; i++ { // Left to right
			var oldPixel = dithered.GrayAt(i, j)
			// Set the pixel to black or white
			var newPixel = blackOrWhite(oldPixel)
			dithered.SetGray(i, j, newPixel)
			// Determine the quantization error
			var quant = (int16(oldPixel.Y) - int16(newPixel.Y)) / 16
			// Spread the quantization error
			dithered.SetGray(i+1, j, color.Gray{i16ToUI8(int16(dithered.GrayAt(i+1, j).Y) + 7*quant)})
			dithered.SetGray(i-1, j+1, color.Gray{i16ToUI8(int16(dithered.GrayAt(i-1, j+1).Y) + 3*quant)})
			dithered.SetGray(i, j+1, color.Gray{i16ToUI8(int16(dithered.GrayAt(i, j+1).Y) + 5*quant)})
			dithered.SetGray(i+1, j+1, color.Gray{i16ToUI8(int16(dithered.GrayAt(i+1, j+1).Y) + quant)})
		}
	}
	return dithered
}
