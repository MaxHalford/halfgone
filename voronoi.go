package main

import (
	"image"
	"image/color"
	"math"
)

// A Tessellation is a map of disjoint sets each containing GrayPoints.
type Tessellation map[image.Point][]image.Point

// BuildVoronoi builds a Voronoi tesselation from a set of generating points.
func BuildVoronoi(sites []image.Point, bounds image.Rectangle) Tessellation {
	var tess = make(Tessellation)
	for _, site := range sites {
		tess[site] = []image.Point{}
	}
	for i := 0; i < bounds.Dx(); i++ {
		for j := 0; j < bounds.Dy(); j++ {
			var (
				point       = image.Point{i, j}
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
		var dist = pointDistance(point, p)
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
	var totalWeight float64
	var cx float64
	var cy float64
	for _, point := range points {
		var weight = float64(weights.GrayAt(point.X, point.Y).Y) + 1
		cx += float64(point.X) * weight
		cy += float64(point.Y) * weight
		totalWeight += weight
	}
	cx /= totalWeight
	cy /= totalWeight
	if cx-math.Floor(cx) < 0.5 {
		centroid.X = int(cx)
	} else {
		centroid.X = int(cx) + 1
	}
	if cy-math.Floor(cy) < 0.5 {
		centroid.Y = int(cy)
	} else {
		centroid.Y = int(cy) + 1
	}
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
	var img = image.NewNRGBA(bounds)
	for site := range tess {
		img.Set(site.X, site.Y, color.Black)
	}
	return img
}
