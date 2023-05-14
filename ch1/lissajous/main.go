// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image/color"
	"math/rand"
	"os"
	"strconv"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

//!+main

var (
	colorRed   = color.RGBA{R: 0xFF, A: 0x01}
	colorGreen = color.RGBA{G: 0xFF, A: 0x01}
	colorBlue  = color.RGBA{B: 0xFF, A: 0x01}
)

var palette = []color.Color{color.Black, colorRed, colorGreen, colorBlue}

const (
	backgroundColorIndex = 0 // next color in palette
	greenOnBlackIndex    = 2 // first color in palette
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		http.HandleFunc("/", lissajousHandler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout, getDefaultLissajousConfig())
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	lissajous(w, parseLissajousConfigFromRequest(r))
}

// i.e.: http://localhost:8000/?cycles=5&res=0.001&size=100&nframes=64&delay=8
func parseLissajousConfigFromRequest(r *http.Request) *lissajousConfig {
	if err := r.ParseForm(); err != nil {
		log.Println(err.Error())
	}
	config := &lissajousConfig{}
	for k, v := range r.Form {
		switch k {
		case "cycles":
			var err error
			if config.cycles, err = strconv.Atoi(v[0]); err != nil {
				log.Println(err.Error())
			}
		case "res":
			var err error
			if config.res, err = strconv.ParseFloat(v[0], 64); err != nil {
				log.Println(err.Error())
			}
		case "size":
			var err error
			if config.size, err = strconv.Atoi(v[0]); err != nil {
				log.Println(err.Error())
			}
		case "nframes":
			var err error
			if config.nframes, err = strconv.Atoi(v[0]); err != nil {
				log.Println(err.Error())
			}
		case "delay":
			var err error
			if config.delay, err = strconv.Atoi(v[0]); err != nil {
				log.Println(err.Error())
			}
		default:
			log.Printf("%v:%v", k, v)
		}
	}
	return config
}

//!-main
