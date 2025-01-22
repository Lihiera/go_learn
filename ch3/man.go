// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

func man(out io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		unit                   = (xmax - xmin) / float64(width)
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z1 := complex(x-unit/2, y-unit/2)
			z2 := complex(x+unit/2, y-unit/2)
			z3 := complex(x+unit/2, y+unit/2)
			z4 := complex(x-unit/2, y+unit/2)
			z := color.RGBA{B: (mandelbrot(z1).B/4.0 + mandelbrot(z2).B/4.0 + mandelbrot(z3).B/4.0 + mandelbrot(z4).B/4.0), A: 255}
			// Image point (px, py) represents complex value z.
			// z = color.RGBA{0, 0, 255, 255}
			img.Set(px, py, z)
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{0, 0, 255 - contrast*n, 255}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}
