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
	var gray = RgbaToGray(img)
	//SaveImagePNG(gray, "screenshots/grayscale.png")

	// Intensity inversion
	//var inverted = InvertGray(gray)
	//SaveImagePNG(inverted, "screenshots/inverted_grayscale.png")

	// Threshold dithering
	//var td = ThresholdDitherer{122}.apply(gray)
	//SaveImagePNG(td, "screenshots/threshold_dithering.png")

	// Random threshold dithering
	var rtd = RandomThresholdDitherer{155, rng}.apply(gray)
	SaveImagePNG(rtd, "screenshots/random_threshold_dithering.png")

	// Bosch and Herman’s grid-based dithering
	//var gd = GridDitherer{5, 3, 8, rng}.apply(gray)
	//SaveImagePNG(gd, "screenshots/grid_dithering.png")

	// Floyd-Steinberg dithering
	var fsd = FloydSteinbergDitherer{}.apply(gray)
	SaveImagePNG(fsd, "screenshots/floyd_steinberg_dithering.png")

	// var sites = ExtractPoints(rtd, 0, 0)
	// var voronoi = BuildVoronoi(sites, gray.Bounds())

	// var vr = DrawTessallationRegions(voronoi, gray.Bounds())
	// SaveImagePNG(vr, "screenshots/voronoi_regions.png")

	// var vs = DrawTessallationSites(voronoi, gray.Bounds())
	// SaveImagePNG(vs, "screenshots/voronoi_sites.png")

	// var weights = gray
	// for i := 0; i < 15; i++ {
	// 	var centroids = CalculateCentroids(voronoi, weights)
	// 	voronoi = BuildVoronoi(centroids, gray.Bounds())
	// }

	// var cvr = DrawTessallationRegions(voronoi, gray.Bounds())
	// SaveImagePNG(cvr, "screenshots/centroidal_voronoi_regions.png")

	// var cvs = DrawTessallationSites(voronoi, gray.Bounds())
	// SaveImagePNG(cvs, "screenshots/centroidal_voronoi_sites.png")
}
