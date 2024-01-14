// Lissajous generates GIF animations of random Lissajous figures.
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
	"os"
)

var palette = []color.Color{
	color.White,
	color.RGBA{0x19, 0x1d, 0x32, 0xff}, // oxford blue
	color.RGBA{0x28, 0x2f, 0x44, 0xff}, // space cadet
	color.RGBA{0x45, 0x3a, 0x49, 0xff}, // english violet
	color.RGBA{0x6d, 0x3b, 0x47, 0xff}, // wine
	color.RGBA{0xba, 0x2c, 0x73, 0xff}, // magenta dye
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	var colorIndex uint8 = 1
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
		colorIndex = uint8(int(phase)%5 + 1)

		if nframes%(len(palette)-1) == 0 {
			colorIndex++
		}
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
