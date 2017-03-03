# halfgone

A collection of halftoning algorithms written in Go. For the while this is not aimed at being reused!


## Original image

```go
var img, err = LoadImage("images/penguin.jpg")
```

![original](examples/images/penguin.jpg)


## Grayscale

```go
var gray = rgbaToGray(img)
```

![grayscale](examples/images/grayscale.png)


## Inverted grayscale

```go
var inverted = InvertGray(gray)
```

![reversed_grayscale](examples/images/inverted_grayscale.png)


## Threshold dithering

```go
var td = halfgone.ThresholdDitherer{Threshold: 122}.Apply(gray)
```

![threshold_dithering](examples/images/threshold_dithering.png)


## Random threshold dithering

```go
var rtd = halfgone.RandomThresholdDitherer{MaxThreshold: 100, RNG: rng}.Apply(gray)
```

![random_threshold_dithering](examples/images/random_threshold_dithering.png)


## Importance sampling

```go
var is = halfgone.ImportanceSampling{N: 2000, Threshold: 100, RNG: rng}.Apply(gray)
```

![importance_sampling](examples/images/importance_sampling.png)


## Bosch and Herman’s grid-based dithering

```go
var gd = halfgone.GridDitherer{K: 5, Alpha: 3, Beta: 8, RNG: rng}.Apply(gray)
```

![grid_dithering](examples/images/grid_dithering.png)


## Floyd-Steinberg dithering

```go
var fsd = halfgone.FloydSteinbergDitherer{}.apply(gray)
```

![floyd_steinberg_dithering](examples/images/floyd_steinberg_dithering.png)
