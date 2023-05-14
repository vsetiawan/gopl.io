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
	lissajous(w, getDefaultLissajousConfig())
}

//!-main
