package main

import (
	"image"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

const (
	cycles  = 5     // number of complete x oscillator revolutions
	res     = 0.001 // angular resolution
	size    = 100   // image canvas covers [-size..+size]
	nframes = 64    // number of animation frames
	delay   = 8     // delay between frames in 10ms units
)

type lissajousConfig struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

func getDefaultLissajousConfig() *lissajousConfig {
	return &lissajousConfig{
		cycles:  cycles,
		res:     res,
		size:    size,
		nframes: nframes,
		delay:   delay,
	}
}

func lissajous(out io.Writer, config *lissajousConfig) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: config.nframes}
	phase := 0.0 // phase difference
	for i := 0; i < config.nframes; i++ {
		currentColorIndex := uint8(i % len(palette)) // Exercise 1.6
		//currentColorIndex := uint8(greenOnBlackIndex) // Exercise 1.5

		rect := image.Rect(0, 0, 2*config.size+1, 2*config.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(config.cycles)*2*math.Pi; t += config.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(config.size+int(x*float64(config.size)+0.5), config.size+int(y*float64(config.size)+0.5),
				currentColorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, config.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
