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
	var gray = halfgone.ImageToGray(img)
	halfgone.SaveImagePNG(gray, "images/grayscale.png")

	// Intensity inversion
	var inverted = halfgone.InvertGray(gray)
	halfgone.SaveImagePNG(inverted, "images/inverted_grayscale.png")

	// Threshold dithering
	var td = halfgone.ThresholdDitherer{Threshold: 127}.Apply(gray)
	halfgone.SaveImagePNG(td, "images/threshold_dithering.png")

	// Random threshold dithering
	var rtd = halfgone.RandomThresholdDitherer{MaxThreshold: 200, RNG: rng}.Apply(gray)
	halfgone.SaveImagePNG(rtd, "images/random_threshold_dithering.png")

	// Importance sampling
	var is = halfgone.ImportanceSampling{N: 4000, Threshold: 100, RNG: rng}.Apply(gray)
	halfgone.SaveImagePNG(is, "images/importance_sampling.png")

	// Bosch and Herman’s grid-based dithering
	var gd = halfgone.GridDitherer{K: 5, Alpha: 3, Beta: 8, RNG: rng}.Apply(gray)
	halfgone.SaveImagePNG(gd, "images/grid_dithering.png")

	// Order-2 ordered dithering
	var o2od = halfgone.Order2OrderedDitherer{}.Apply(gray)
	halfgone.SaveImagePNG(o2od, "images/order_2_ordered_dithering.png")

	// Order-3 ordered dithering
	var o3od = halfgone.Order3OrderedDitherer{}.Apply(gray)
	halfgone.SaveImagePNG(o3od, "images/order_3_ordered_dithering.png")

	// Order-4 ordered dithering
	var o4od = halfgone.Order4OrderedDitherer{}.Apply(gray)
	halfgone.SaveImagePNG(o4od, "images/order_4_ordered_dithering.png")

	// Order-8 ordered dithering
	var o8od = halfgone.Order8OrderedDitherer{}.Apply(gray)
	halfgone.SaveImagePNG(o8od, "images/order_8_ordered_dithering.png")

	// Floyd-Steinberg dithering
	var fsd = halfgone.FloydSteinbergDitherer{}.Apply(gray)
	halfgone.SaveImagePNG(fsd, "images/floyd_steinberg_dithering.png")

	// Jarvis-Judice-Ninke dithering
	var jjnd = halfgone.JarvisJudiceNinkeDitherer{}.Apply(gray)
	halfgone.SaveImagePNG(jjnd, "images/jarvis_judice_ninke_dithering.png")

	// Stucki dithering
	var sd = halfgone.StuckiDitherer{}.Apply(gray)
	halfgone.SaveImagePNG(sd, "images/stucki_dithering.png")

	// Atkinson dithering
	var ad = halfgone.AtkinsonDitherer{}.Apply(gray)
	halfgone.SaveImagePNG(ad, "images/atkinson_dithering.png")

	// Burkes dithering
	var bd = halfgone.BurkesDitherer{}.Apply(gray)
	halfgone.SaveImagePNG(bd, "images/burkes_dithering.png")

	// Sierra dithering
	var sid = halfgone.SierraDitherer{}.Apply(gray)
	halfgone.SaveImagePNG(sid, "images/seria_dithering.png")

	// Two-row Sierra dithering
	var trsd = halfgone.TwoRowSierraDitherer{}.Apply(gray)
	halfgone.SaveImagePNG(trsd, "images/two_row_seria_dithering.png")

	// Sierra Lite dithering
	var sld = halfgone.SierraLiteDitherer{}.Apply(gray)
	halfgone.SaveImagePNG(sld, "images/seria_lite_dithering.png")
}
