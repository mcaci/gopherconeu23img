package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strings"
)

func letter() {
	const a = ` * 
* *
***
* *`
	fmt.Println(a)
	aRect := image.Rect(0, 0, strings.Count(a, "\n")+1, strings.Count(a, "\n")+1)
	aImg := image.NewRGBA(aRect)
	var x, y int
	for i := range a {
		switch a[i] {
		case '*':
			aImg.Set(x, y, color.Black)
			x++
		case ' ':
			aImg.SetRGBA(x, y, color.RGBA{R: 255, B: 200, G: 200, A: 255})
			x++
		case '\n':
			y++
			x = 0
		default:
			log.Print(a[i], "what?")
		}
	}

	f, err := os.Create("a.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, aImg); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
