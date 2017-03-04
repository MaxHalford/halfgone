# halfgone

A collection of halftoning algorithms written in Go. These are not optimized for production where you would typically want to use bit shifting.


## Original image

```go
LoadImage("images/penguin.jpg")
```

![original](examples/images/penguin.jpg)


## Grayscale

```go
ImageToGray(img)
```

![grayscale](examples/images/grayscale.png)


## Inverted grayscale

```go
InvertGray(gray)
```

![reversed_grayscale](examples/images/inverted_grayscale.png)


## Threshold dithering

```go
halfgone.ThresholdDitherer{Threshold: 122}.Apply(gray)
```

![threshold_dithering](examples/images/threshold_dithering.png)


## Random threshold dithering

```go
halfgone.RandomThresholdDitherer{MaxThreshold: 100, RNG: rng}.Apply(gray)
```

![random_threshold_dithering](examples/images/random_threshold_dithering.png)


## Importance sampling

```go
halfgone.ImportanceSampling{N: 2000, Threshold: 100, RNG: rng}.Apply(gray)
```

![importance_sampling](examples/images/importance_sampling.png)


## Bosch and Herman’s grid-based dithering

```go
halfgone.GridDitherer{K: 5, Alpha: 3, Beta: 8, RNG: rng}.Apply(gray)
```

![grid_dithering](examples/images/grid_dithering.png)


## Floyd-Steinberg dithering

```go
halfgone.FloydSteinbergDitherer{}.apply(gray)
```

![floyd_steinberg_dithering](examples/images/floyd_steinberg_dithering.png)


## Jarvis-Judice-Ninke dithering

```go
halfgone.JarvisJudiceNinkeDitherer{}.Apply(gray)
```

![jarvis_judice_ninke_dithering](examples/images/jarvis_judice_ninke_dithering.png)


## Stucki dithering

```go
halfgone.StuckiDitherer{}.Apply(gray)
```

![stucki_dithering](examples/images/stucki_dithering.png)


## Atkinson dithering

```go
halfgone.AtkinsonDitherer{}.Apply(gray)
```

![atkinson_dithering](examples/images/atkinson_dithering.png)


## Burkes dithering

```go
halfgone.BurkesDitherer{}.Apply(gray)
```

![burkes_dithering](examples/images/burkes_dithering.png)


## Sierra dithering

```go
halfgone.SierraDitherer{}.Apply(gray)
```

![seria_dithering](examples/images/seria_dithering.png)


## Two-row Sierra dithering

```go
halfgone.TwoRowSierraDitherer{}.Apply(gray)
```

![two_row_seria_dithering](examples/images/two_row_seria_dithering.png)


## Sierra Lite dithering

```go
halfgone.SierraLiteDitherer{}.Apply(gray)
```

![seria_lite_dithering](examples/images/seria_lite_dithering.png)

