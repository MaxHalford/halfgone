# halfgone

This repository contains implementations of *digital halftoning* - also called *dithering* -  algorithms written in Go. The implementations are restricted to black and white rendering and are based on the `image` library from Go's standard library.

The implementations are quite fast but are not optimized for production where you would typically want to use bit shifting when possible. I moved the common code for error-diffusion dithering into a separate functions because it's always the same underlying algorithm, whether it be Floyd-Steinberg dithering or Stucki dithering. I did the same for ordered dithering. In production you would probably want to choose a particular dithering algorithm and avoid using generic code which makes it harder to write optimized code.

If you are interested in digital halftoning, [this web page](http://www.efg2.com/Lab/Library/ImageProcessing/DHALF.TXT) is, in my opinion, a fantastic introduction. I've also written a [blog post](https://maxhalford.github.io/blog/halftoning-1/) which goes through some of the implementations.


## Original image

```go
img := LoadImage("images/penguin.jpg")
```

![original](examples/images/penguin.jpg)


## Grayscale

```go
gray := ImageToGray(img)
```

![grayscale](examples/images/grayscale.png)


## Inverted grayscale

```go
InvertGray(gray)
```

![reversed_grayscale](examples/images/inverted_grayscale.png)


## Threshold dithering

```go
halfgone.ThresholdDitherer{Threshold: 127}.Apply(gray)
```

![threshold_dithering](examples/images/threshold_dithering.png)


## Random threshold dithering

```go
halfgone.RandomThresholdDitherer{MaxThreshold: 255, RNG: rng}.Apply(gray)
```

![random_threshold_dithering](examples/images/random_threshold_dithering.png)


## Importance sampling

```go
halfgone.ImportanceSampling{N: 4000, Threshold: 100, RNG: rng}.Apply(gray)
```

![importance_sampling](examples/images/importance_sampling.png)


## Bosch and Herman’s grid-based dithering

```go
halfgone.GridDitherer{K: 5, Alpha: 3, Beta: 8, RNG: rng}.Apply(gray)
```

![grid_dithering](examples/images/grid_dithering.png)


## Ordered dithering

### Order-2 ordered dithering

```go
halfgone.Order2OrderedDitherer{}.Apply(gray)
```

![order_2_ordered_dithering](examples/images/order_2_ordered_dithering.png)


### Order-3 ordered dithering

```go
halfgone.Order3OrderedDitherer{}.Apply(gray)
```

![order_3_ordered_dithering](examples/images/order_3_ordered_dithering.png)


### Order-4 ordered dithering

```go
halfgone.Order4OrderedDitherer{}.Apply(gray)
```

![order_4_ordered_dithering](examples/images/order_4_ordered_dithering.png)


### Order-8 ordered dithering

```go
halfgone.Order8OrderedDitherer{}.Apply(gray)
```

![order_8_ordered_dithering](examples/images/order_8_ordered_dithering.png)


## Error-diffusion dithering

### Floyd-Steinberg dithering

```go
halfgone.FloydSteinbergDitherer{}.apply(gray)
```

![floyd_steinberg_dithering](examples/images/floyd_steinberg_dithering.png)


### Jarvis-Judice-Ninke dithering

```go
halfgone.JarvisJudiceNinkeDitherer{}.Apply(gray)
```

![jarvis_judice_ninke_dithering](examples/images/jarvis_judice_ninke_dithering.png)


### Stucki dithering

```go
halfgone.StuckiDitherer{}.Apply(gray)
```

![stucki_dithering](examples/images/stucki_dithering.png)


### Atkinson dithering

```go
halfgone.AtkinsonDitherer{}.Apply(gray)
```

![atkinson_dithering](examples/images/atkinson_dithering.png)


### Burkes dithering

```go
halfgone.BurkesDitherer{}.Apply(gray)
```

![burkes_dithering](examples/images/burkes_dithering.png)


### Sierra dithering

```go
halfgone.SierraDitherer{}.Apply(gray)
```

![seria_dithering](examples/images/seria_dithering.png)


### Two-row Sierra dithering

```go
halfgone.TwoRowSierraDitherer{}.Apply(gray)
```

![two_row_seria_dithering](examples/images/two_row_seria_dithering.png)


### Sierra Lite dithering

```go
halfgone.SierraLiteDitherer{}.Apply(gray)
```

![seria_lite_dithering](examples/images/seria_lite_dithering.png)

## License

The MIT License (MIT). Please see the [license file](LICENSE) for more information.
