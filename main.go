package main

import (
	"image/png"
	"math/rand"
	"os"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	var img, _ = loadImage("penguin.jpg")
	var gray = rgbaToGray(img)
	//var dithered = ThresholdDitherer{122}.apply(gray)
	//var dithered = GridDitherer{5, 3, 8, rng}.apply(gray)
	//var dithered = FloydSteinbergDitherer{}.apply(gray)

	var voronoi = BuildVoronoi(gray.Bounds(), 1000, rng)
	for i := 0; i < 100; i++ {
		voronoi = CenterVoronoi(voronoi, gray)
	}
	var dithered = DrawTessallation(voronoi, gray.Bounds())

	// Save as out.png
	f, _ := os.Create("out.png")
	defer f.Close()
	png.Encode(f, dithered)
}
