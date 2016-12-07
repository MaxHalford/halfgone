package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"
)

// A Tessellation is a map of disjoint sets each containing GrayPoints.
type Tessellation map[image.Point][]image.Point

// BuildVoronoi constructs a Voronoï tessalation.
func BuildVoronoi(bounds image.Rectangle, nSites int, rng *rand.Rand) Tessellation {
	var (
		tess  = make(Tessellation)
		sites = make([]image.Point, nSites)
	)
	// Generate random sites
	for i := 0; i < nSites; i++ {
		var site = randPoint(bounds, rng)
		tess[site] = []image.Point{}
		sites[i] = site
	}
	// Match each point to it's closest site
	for i := 0; i < bounds.Dx(); i++ {
		for j := 0; j < bounds.Dy(); j++ {
			var (
				point       = image.Point{i, j}
				closestSite = findClosestSite(point, sites)
			)
			tess[closestSite] = append(tess[closestSite], point)
		}
	}
	return tess
}

// CenterVoronoi centers a Voronoï tessallation with Lloyd's algorithm. It
// starts by calculating the weighted centroids of each region and then
// reassigns the points to the closest centroid. The point weights are
// determined by looking at the gray intensities of the provided grayscale
// image.
func CenterVoronoi(tess Tessellation, weights *image.Gray) Tessellation {
	var (
		centeredTess = make(Tessellation)
		centroids    []image.Point
	)
	// Determine the centroid of each region in the original tessallation
	for _, points := range tess {
		var centroid = calculateCentroid(points, weights)
		centeredTess[centroid] = []image.Point{}
		centroids = append(centroids, centroid)
	}
	// Reassign the points to the new centroids
	for _, points := range tess {
		for _, point := range points {
			var closestCentroid = findClosestSite(point, centroids)
			centeredTess[closestCentroid] = append(centeredTess[closestCentroid], point)
		}
	}
	return centeredTess
}

// calculateCentroid determines the weighted centroid of a set of points.
func calculateCentroid(points []image.Point, weights *image.Gray) (centroid image.Point) {
	var totalWeight int
	for _, point := range points {
		var weight = 255 - int(weights.GrayAt(point.X, point.Y).Y)
		centroid.X += point.X * weight
		centroid.Y += point.Y * weight
		totalWeight += weight
	}
	centroid.X /= totalWeight
	centroid.Y /= totalWeight
	return
}

func findClosestSite(point image.Point, sites []image.Point) (closestSite image.Point) {
	var min = math.Inf(1)
	for _, site := range sites {
		var dist = pointDistance(site, point)
		if dist < min {
			min = dist
			closestSite = site
		}
	}
	return
}

// DrawTessallationRegions draws tessallation regions by assigning a random
// color to each region.
func DrawTessallationRegions(tess Tessellation, bounds image.Rectangle) image.Image {
	var img = image.NewNRGBA(bounds)
	for _, points := range tess {
		var c = randColor()
		for _, point := range points {
			img.Set(point.X, point.Y, c)
		}
	}
	return img
}

// DrawTessallationSites draws tessallation sites.
func DrawTessallationSites(tess Tessellation, bounds image.Rectangle) image.Image {
	var img = image.NewNRGBA(bounds)
	for site := range tess {
		img.Set(site.X, site.Y, color.Black)
	}
	return img
}
