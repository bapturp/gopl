package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{
	color.White,
	color.RGBA{0x19, 0x04, 0x82, 0xff},
	color.RGBA{0x77, 0x52, 0xfe, 0xFF},
	color.RGBA{0x8e, 0x8f, 0xfa, 0xFF},
	color.RGBA{0xc2, 0xd9, 0xff, 0xFF},
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	if r.Form.Has("cycles") {
		cycles, err := strconv.Atoi(r.Form.Get("cycles"))
		if err != nil {
			log.Print(err)
		}
		lissajous(w, cycles)
	} else {
		lissajous(w, 5)
	}
}

func lissajous(out io.Writer, cycles int) {
	// cycles: number of complete x oscillator revolutions
	// must be a positive integer
	if cycles < 5 {
		cycles = 5
	}
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	var colorIndex uint8 = 1

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}

		if nframes%(len(palette)-1) == 0 {
			colorIndex++
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // ignoring errors
}
