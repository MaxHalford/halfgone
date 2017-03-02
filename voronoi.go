package main

import (
	"image"
	"image/color"
	"math"
)

// A Tessellation is a map of disjoint sets each containing GrayPoints.
type Tessellation map[image.Point][]image.Point

// BuildVoronoi builds a Voronoi tessellation from a set of generating points.
func BuildVoronoi(sites []image.Point, bounds image.Rectangle) Tessellation {
	var tess = make(Tessellation)
	for _, site := range sites {
		tess[site] = []image.Point{}
	}
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			var (
				point       = image.Point{x, y}
				closestSite = findClosestPoint(point, sites)
			)
			tess[closestSite] = append(tess[closestSite], point)
		}
	}
	return tess
}

func findClosestPoint(point image.Point, points []image.Point) (closest image.Point) {
	var min = math.Inf(1)
	for _, p := range points {
		var dist = euclideanDistance(point, p)
		if dist < min {
			min = dist
			closest = p
		}
	}
	return
}

// CalculateCentroids returns a slice containing the centroid of each region in
// a Voronoi tesselation.
func CalculateCentroids(tess Tessellation, weights *image.Gray) []image.Point {
	var (
		centroids = make([]image.Point, len(tess))
		i         int
	)
	for _, points := range tess {
		centroids[i] = calculateCentroid(points, weights)
		i++
	}
	return centroids
}

func calculateCentroid(points []image.Point, weights *image.Gray) (centroid image.Point) {
	var (
		totalWeight float64
		cx          float64
		cy          float64
	)
	for _, point := range points {
		var weight = float64(weights.GrayAt(point.X, point.Y).Y)
		cx += (float64(point.X) + 0.5) * weight
		cy += (float64(point.Y) + 0.5) * weight
		totalWeight += weight
	}
	cx /= totalWeight
	cy /= totalWeight
	centroid.X = int(cx)
	centroid.Y = int(cy)
	return
}

// DrawTessallationRegions draws tessallation regions by assigning a random
// color to each region.
func DrawTessallationRegions(tess Tessellation, bounds image.Rectangle) image.Image {
	var img = image.NewNRGBA(bounds)
	for site, region := range tess {
		var col = randColor()
		for _, point := range region {
			img.Set(point.X, point.Y, col)
		}
		img.Set(site.X, site.Y, color.Black)
	}
	return img
}

// DrawTessallationSites draws tessallation sites.
func DrawTessallationSites(tess Tessellation, bounds image.Rectangle) image.Image {
	var img = MakeGray(bounds, 255)
	for site := range tess {
		img.Set(site.X, site.Y, color.Black)
	}
	return img
}
