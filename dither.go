package halfgone

import (
	"image"
	"image/color"
	"math"
	"math/rand"
)

// A Ditherer convert a grayscale image with intensities going from 0 (black) to
// 255 (white) into a black and white image.
type Ditherer interface {
	Apply(gray *image.Gray) *image.Gray
}

// ThresholdDitherer converts each pixel in a grayscale image to black or white
// depending on the intensity of the pixel. If a pixel's intensity is above the
// given threshold then the pixel becomes white, else it becomes black.
type ThresholdDitherer struct {
	Threshold uint8
}

// Apply threshold dithering.
func (td ThresholdDitherer) Apply(gray *image.Gray) *image.Gray {
	var (
		bounds   = gray.Bounds()
		dithered = image.NewGray(bounds)
	)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			if gray.GrayAt(x, y).Y > td.Threshold {
				dithered.SetGray(x, y, color.Gray{255}) // White
			} else {
				dithered.SetGray(x, y, color.Gray{0}) // Black
			}
		}
	}
	return dithered
}

// RandomThresholdDitherer works the same way as ThresholdDitherer except that
// the threshold is randomly sampled for each pixel. This way some pixels are
// white when they would have been actually black.
type RandomThresholdDitherer struct {
	MaxThreshold int
	RNG          *rand.Rand
}

// Apply random threshold dithering.
func (rtd RandomThresholdDitherer) Apply(gray *image.Gray) *image.Gray {
	var (
		bounds   = gray.Bounds()
		dithered = image.NewGray(bounds)
	)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			var threshold = uint8(rtd.RNG.Intn(rtd.MaxThreshold))
			if gray.GrayAt(x, y).Y > threshold {
				dithered.SetGray(x, y, color.Gray{255}) // White
			} else {
				dithered.SetGray(x, y, color.Gray{0}) // Black
			}
		}
	}
	return dithered
}

// ImportanceSampling implements importance sampling.
type ImportanceSampling struct {
	N         int   // Number of points to sample
	Threshold uint8 // Threshold after which intensities are ignored (the threshold is not ignored)
	RNG       *rand.Rand
}

// Apply importance sampling.
func (is ImportanceSampling) Apply(gray *image.Gray) *image.Gray {
	var (
		bounds    = gray.Bounds()
		dithered  = makeGray(bounds, 255)
		histogram = make(map[uint8][]image.Point)
	)
	// Build histogram
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			var intensity = gray.GrayAt(x, y).Y
			if intensity <= is.Threshold {
				histogram[intensity] = append(histogram[intensity], image.Point{x, y})
			}
		}
	}
	// Build roulette wheel
	var roulette = make([]int, is.Threshold+1)
	roulette[0] = 255 * len(histogram[0])
	for i := 1; i < len(roulette); i++ {
		roulette[i] = roulette[i-1] + (255-i)*len(histogram[uint8(i)])
	}
	// Run the wheel
	for i := 0; i < is.N; i++ {
		var bin = uint8(binarySearchInt(is.RNG.Intn(roulette[is.Threshold]), roulette))
		var point = histogram[bin][is.RNG.Intn(len(histogram[bin]))]
		dithered.SetGray(point.X, point.Y, color.Gray{0})
	}
	return dithered
}

// GridDitherer implements Bosch and Herman's grid-based method.
type GridDitherer struct {
	K     int     // Size in pixels of a side of a cell
	Alpha float64 // Minimum desired number of points in a cell
	Beta  float64 // Maximum desired number of points in a cell
	RNG   *rand.Rand
}

// Apply random grid dithering.
func (gd GridDitherer) Apply(gray *image.Gray) *image.Gray {
	var (
		bounds   = gray.Bounds()
		dithered = makeGray(bounds, 255)
		width    = bounds.Dx()
		height   = bounds.Dy()
	)
	for x := 0; x < width; x += gd.K {
		for y := 0; y < height; y += gd.K {
			var (
				cell = RGBAToGray(gray.SubImage(image.Rect(x, y, x+gd.K, y+gd.K)))
				mu   = avgIntensity(cell)              // Mean grayscale value of the cell
				n    = math.Pow((1-mu)*gd.Beta, 2) / 3 // Number of points to sample
			)
			if n < gd.Alpha {
				n = 0
			}
			// Sample n random points in belonging to the cell
			for i := 0; i < int(n); i++ {
				var (
					xx = randInt(x, min(x+gd.K, width), gd.RNG)
					yy = randInt(y, min(y+gd.K, height), gd.RNG)
				)
				dithered.SetGray(xx, yy, color.Gray{0})
			}
		}
	}
	return dithered
}

// FloydSteinbergDitherer implements the Floyd-Steingberg algorithm, which is a
// variation of the error diffusion method.
type FloydSteinbergDitherer struct{}

// Apply Floyd-Steinberg dithering.
func (fsd FloydSteinbergDitherer) Apply(gray *image.Gray) *image.Gray {
	var (
		bounds   = gray.Bounds()
		dithered = copyGray(gray)
	)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ { // Top to bottom
		for x := bounds.Min.X; x < bounds.Max.X; x++ { // Left to right
			var oldPixel = dithered.GrayAt(x, y)
			// Set the pixel to black or white
			var newPixel = color.Gray{0} // Black
			if oldPixel.Y > 127 {
				newPixel = color.Gray{255} // White
			}
			dithered.SetGray(x, y, newPixel)
			// Determine the quantization error
			var quant = (int16(oldPixel.Y) - int16(newPixel.Y)) / 16
			// Spread the quantization error
			dithered.SetGray(x+1, y, color.Gray{i16ToUI8(int16(dithered.GrayAt(x+1, y).Y) + 7*quant)})
			dithered.SetGray(x-1, y+1, color.Gray{i16ToUI8(int16(dithered.GrayAt(x-1, y+1).Y) + 3*quant)})
			dithered.SetGray(x, y+1, color.Gray{i16ToUI8(int16(dithered.GrayAt(x, y+1).Y) + 5*quant)})
			dithered.SetGray(x+1, y+1, color.Gray{i16ToUI8(int16(dithered.GrayAt(x+1, y+1).Y) + quant)})
		}
	}
	return dithered
}
