package halfgone

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"sort"
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
			var threshold = uint8(rtd.RNG.Intn(rtd.MaxThreshold + 1))
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
		sampled   = make(map[image.Point]bool)
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
	roulette[0] = 256 * len(histogram[0])
	for i := 1; i < len(roulette); i++ {
		roulette[i] = roulette[i-1] + (256-i)*len(histogram[uint8(i)])
	}
	// Run the wheel
	var i int
	for i < is.N {
		var (
			ball  = is.RNG.Intn(roulette[len(roulette)-1])
			bin   = uint8(sort.SearchInts(roulette, ball))
			point = histogram[bin][is.RNG.Intn(len(histogram[bin]))]
		)
		// Add the point if it hasn't been already sampled
		if !sampled[point] {
			dithered.SetGray(point.X, point.Y, color.Gray{0})
			sampled[point] = true
			i++
		}
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
	)
	for x := bounds.Min.X; x < bounds.Max.X; x += gd.K {
		for y := bounds.Min.Y; y < bounds.Max.Y; y += gd.K {
			var (
				cell = ImageToGray(gray.SubImage(image.Rect(x, y, x+gd.K, y+gd.K)))
				mu   = avgIntensity(cell)
				n    = math.Pow((1-mu)*gd.Beta, 2) / 3
			)
			if n < gd.Alpha {
				n = 0
			}
			for k := 0; k < int(n); k++ {
				var (
					xx = randInt(x, min(x+gd.K, bounds.Max.X), gd.RNG)
					yy = randInt(y, min(y+gd.K, bounds.Max.Y), gd.RNG)
				)
				dithered.SetGray(xx, yy, color.Gray{0})
			}
		}
	}
	return dithered
}

// A Pattern is a matrix of threshold values used for ordered dithering.
type Pattern [][]uint8

func applyOrderedDithering(gray *image.Gray, pattern Pattern) *image.Gray {
	var (
		order    = len(pattern)
		bounds   = gray.Bounds()
		dithered = makeGray(bounds, 255)
	)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			var threshold = pattern[x%order][y%order]
			if gray.GrayAt(x, y).Y > threshold {
				dithered.SetGray(x, y, color.Gray{255}) // White
			} else {
				dithered.SetGray(x, y, color.Gray{0}) // Black
			}
		}
	}
	return dithered
}

// Order2OrderedDitherer implements order-2 ordered dithering.
type Order2OrderedDitherer struct{}

// Apply order-2 ordered dithering dithering.
func (o2od Order2OrderedDitherer) Apply(gray *image.Gray) *image.Gray {
	var pattern = Pattern{
		{0, 170},
		{255, 85},
	}
	return applyOrderedDithering(gray, pattern)
}

// Order3OrderedDitherer implements order-3 ordered dithering.
type Order3OrderedDitherer struct{}

// Apply order-3 ordered dithering dithering.
func (o3od Order3OrderedDitherer) Apply(gray *image.Gray) *image.Gray {
	var pattern = Pattern{
		{0, 223, 95},
		{191, 159, 63},
		{127, 31, 255},
	}
	return applyOrderedDithering(gray, pattern)
}

// Order4OrderedDitherer implements order-4 ordered dithering.
type Order4OrderedDitherer struct{}

// Apply order-4 ordered dithering dithering.
func (o4od Order4OrderedDitherer) Apply(gray *image.Gray) *image.Gray {
	var pattern = Pattern{
		{0, 136, 34, 170},
		{204, 68, 238, 102},
		{51, 187, 17, 153},
		{255, 119, 221, 85},
	}
	return applyOrderedDithering(gray, pattern)
}

// Order8OrderedDitherer implements order-8 ordered dithering.
type Order8OrderedDitherer struct{}

// Apply order-8 ordered dithering dithering.
func (o8od Order8OrderedDitherer) Apply(gray *image.Gray) *image.Gray {
	var pattern = Pattern{
		{0, 194, 48, 242, 12, 206, 60, 255},
		{129, 64, 178, 113, 141, 76, 190, 125},
		{32, 226, 16, 210, 44, 238, 28, 222},
		{161, 97, 145, 80, 174, 109, 157, 93},
		{8, 202, 56, 250, 4, 198, 52, 246},
		{137, 72, 186, 121, 133, 68, 182, 117},
		{40, 234, 24, 218, 36, 230, 20, 214},
		{170, 105, 153, 89, 165, 101, 149, 85},
	}
	return applyOrderedDithering(gray, pattern)
}

// A DiffusionCell indicates a relative position and a diffusion intensity used for error diffusion.
type DiffusionCell struct {
	x int
	y int
	m int16
}

// A DiffusionMask contains a slice of DiffusionCells and a normalization divisor.
type DiffusionMask struct {
	divisor int16
	cells   []DiffusionCell
}

func applyErrorDiffusion(gray *image.Gray, mask DiffusionMask) *image.Gray {
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
			var quant = (int16(oldPixel.Y) - int16(newPixel.Y)) / mask.divisor
			// Spread the quantization error
			for _, c := range mask.cells {
				var newIntensity = int16(dithered.GrayAt(x+c.x, y+c.y).Y) + int16(c.m*quant)
				dithered.SetGray(x+c.x, y+c.y, color.Gray{i16ToUI8(newIntensity)})
			}
		}
	}
	return dithered
}

// FloydSteinbergDitherer implements Floyd-Steingberg dithering.
type FloydSteinbergDitherer struct{}

// Apply Floyd-Steinberg dithering.
func (fsd FloydSteinbergDitherer) Apply(gray *image.Gray) *image.Gray {
	var mask = DiffusionMask{
		divisor: 16,
		cells: []DiffusionCell{
			{1, 0, 7},
			{-1, 1, 3},
			{0, 1, 5},
			{1, 1, 1},
		},
	}
	return applyErrorDiffusion(gray, mask)
}

// JarvisJudiceNinkeDitherer implements Jarvis-Judice-Ninke dithering.
type JarvisJudiceNinkeDitherer struct{}

// Apply Jarvis-Judice-Ninke dithering.
func (jjnd JarvisJudiceNinkeDitherer) Apply(gray *image.Gray) *image.Gray {
	var mask = DiffusionMask{
		divisor: 48,
		cells: []DiffusionCell{
			{1, 0, 7},
			{2, 0, 5},
			{-2, 1, 3},
			{-1, 1, 5},
			{0, 1, 7},
			{1, 1, 5},
			{2, 1, 3},
			{-2, 2, 1},
			{-1, 2, 3},
			{0, 2, 5},
			{1, 2, 3},
			{2, 2, 1},
		},
	}
	return applyErrorDiffusion(gray, mask)
}

// StuckiDitherer implements Stucki dithering.
type StuckiDitherer struct{}

// Apply Stucki dithering.
func (sd StuckiDitherer) Apply(gray *image.Gray) *image.Gray {
	var mask = DiffusionMask{
		divisor: 42,
		cells: []DiffusionCell{
			{1, 0, 8},
			{2, 0, 4},
			{-2, 1, 2},
			{-1, 1, 4},
			{0, 1, 8},
			{1, 1, 4},
			{2, 1, 2},
		},
	}
	return applyErrorDiffusion(gray, mask)
}

// AtkinsonDitherer implements Atkinson dithering.
type AtkinsonDitherer struct{}

// Apply Atkinson dithering.
func (ad AtkinsonDitherer) Apply(gray *image.Gray) *image.Gray {
	var mask = DiffusionMask{
		divisor: 8,
		cells: []DiffusionCell{
			{1, 0, 1},
			{2, 0, 1},
			{-1, 1, 1},
			{0, 1, 1},
			{1, 1, 1},
			{0, 2, 1},
		},
	}
	return applyErrorDiffusion(gray, mask)
}

// BurkesDitherer implements Burkes dithering.
type BurkesDitherer struct{}

// Apply Burkes dithering.
func (ad BurkesDitherer) Apply(gray *image.Gray) *image.Gray {
	var mask = DiffusionMask{
		divisor: 32,
		cells: []DiffusionCell{
			{1, 0, 8},
			{2, 0, 4},
			{-2, 1, 2},
			{-1, 1, 4},
			{0, 1, 8},
			{1, 1, 4},
			{2, 1, 2},
		},
	}
	return applyErrorDiffusion(gray, mask)
}

// SierraDitherer implements Sierra dithering.
type SierraDitherer struct{}

// Apply Sierra dithering.
func (sd SierraDitherer) Apply(gray *image.Gray) *image.Gray {
	var mask = DiffusionMask{
		divisor: 32,
		cells: []DiffusionCell{
			{1, 0, 5},
			{2, 0, 3},
			{-2, 1, 2},
			{-1, 1, 4},
			{0, 1, 5},
			{1, 1, 4},
			{2, 1, 2},
			{-1, 2, 2},
			{0, 2, 3},
			{1, 2, 2},
		},
	}
	return applyErrorDiffusion(gray, mask)
}

// TwoRowSierraDitherer implements Two-row Sierra dithering.
type TwoRowSierraDitherer struct{}

// Apply Two-row Sierra dithering.
func (sd TwoRowSierraDitherer) Apply(gray *image.Gray) *image.Gray {
	var mask = DiffusionMask{
		divisor: 16,
		cells: []DiffusionCell{
			{1, 0, 4},
			{2, 0, 3},
			{-2, 1, 2},
			{-1, 1, 2},
			{0, 1, 3},
			{1, 1, 2},
			{2, 1, 1},
		},
	}
	return applyErrorDiffusion(gray, mask)
}

// SierraLiteDitherer implements Sierra Lite dithering.
type SierraLiteDitherer struct{}

// Apply Sierra Lite dithering.
func (sd SierraLiteDitherer) Apply(gray *image.Gray) *image.Gray {
	var mask = DiffusionMask{
		divisor: 4,
		cells: []DiffusionCell{
			{1, 0, 1},
			{-1, 1, 1},
			{0, 1, 1},
		},
	}
	return applyErrorDiffusion(gray, mask)
}
