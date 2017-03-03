# halfgone

A collection of halftoning algorithms written in Go. For the while this is not aimed at being reused!


## Original image

```go
var img, err = LoadImage("images/penguin.jpg")
```

![original](images/penguin.jpg)


## Grayscale

```go
var gray = rgbaToGray(img)
```

![grayscale](images/grayscale.png)


## Inverted grayscale

```go
var inverted = InvertGray(gray)
```

![reversed_grayscale](images/reversed_grayscale.png)


## Threshold dithering

```go
var td = ThresholdDitherer{
    thresold: 122,
}.apply(gray)
```

![threshold_dithering](images/threshold_dithering.png)


## Random threshold dithering

```go
var rtd = RandomThresholdDitherer{
    maxThreshold: 155,
    rng: rand.New(rand.NewSource(time.Now().UnixNano())),
}.apply(gray)
```

![random_threshold_dithering](images/random_threshold_dithering.png)


## Bosch and Herman’s grid-based dithering

```go
var gd = GridDitherer{
    k:     5, // Size in pixels of a side of a cell
    alpha: 3, // Minimum desired number of points in a cell
    beta:  8, // Maximum desired number of points in a cell
    rng:   rand.New(rand.NewSource(time.Now().UnixNano())),
}.apply(gray)
```

![grid_dithering](images/grid_dithering.png)


## Importance sampling

```go
var is = ImportanceSampling{
    n: 2000,
    threshold: 100,
    rng:   rand.New(rand.NewSource(time.Now().UnixNano())),
}.apply(gray)
```

![importance_sampling](images/importance_sampling.png)


## Floyd-Steinberg dithering

```go
var fsd = FloydSteinbergDitherer{}.apply(gray)
```

![floyd_steinberg_dithering](images/floyd_steinberg_dithering.png)
