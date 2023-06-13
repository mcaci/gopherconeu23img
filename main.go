package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func main() {
	l := flag.Int("l", 0, "Length of the image")
	h := flag.Int("h", 0, "Height of the image")
	path := flag.String("o", "examples/text.png", "path of the output image/gif")
	bgColorHex := flag.String("bgHex", "0x4d3178", "Hexadecimal value for the background color")
	fgColorHex := flag.String("fgHex", "0xabc", "Hexadecimal value for the color of the text")
	fontPath := flag.String("fontPath", "/usr/share/fonts/truetype/ubuntu/Ubuntu-R.ttf", "path of the font to use")
	fontSize := flag.Float64("fontSize", 32.0, "font size of the output text in the image")
	xPtFactor := flag.Float64("xPtFactor", 0.5, "x size factor of one letter box")
	yPtFactor := flag.Float64("yPtFactor", 1.0, "y size factor of one letter box")
	figlet := flag.String("figlet", "banner", "name of the figlet font; see https://github.com/common-nighthawk/go-figure/tree/master/fonts for the values and http://www.figlet.org/examples.html for the actual effect")
	banner := flag.Bool("banner", false, "if true it's a banner gif, else it's a picture")
	blink := flag.Bool("blink", false, "if true it's a plain blinking gif, else it's a picture")
	alt := flag.Bool("alt", false, "if true it's a alternating colors blinking gif, else it's a picture")
	flag.Parse()

	asciiArtLines := prepareText(strings.Join(flag.Args(), " "), *figlet)
	lImg, hImg := imgBounds(asciiArtLines, *fontSize, *xPtFactor, *yPtFactor)
	if *l != 0 {
		lImg = *l
	}
	if *h != 0 {
		hImg = *h
	}

	switch {
	case *banner:
		images := makeBanner(asciiArtLines, lImg, hImg, *bgColorHex, *fgColorHex, *fontPath, *fontSize, *xPtFactor, *yPtFactor)
		writeGif(images, 5, *path)
	case *blink:
		images := makeBlink(asciiArtLines, lImg, hImg, *bgColorHex, *fgColorHex, *fontPath, *fontSize, *xPtFactor, *yPtFactor)
		writeGif(images, 75, *path)
	case *alt:
		images := makeAlt(asciiArtLines, lImg, hImg, *bgColorHex, *fgColorHex, *fontPath, *fontSize, *xPtFactor, *yPtFactor)
		writeGif(images, 75, *path)
	default:
		image := makePng(asciiArtLines, lImg, hImg, *bgColorHex, *fgColorHex, *fontPath, *fontSize, *xPtFactor, *yPtFactor)
		writePng(image, *path)
	}
}

func mustFile(name string) *os.File {
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func writeGif(images []*image.Paletted, delay int, path string) {
	f := mustFile(path)
	defer f.Close()
	delays := make([]int, len(images))
	for i := range delays {
		delays[i] = delay
	}
	err := gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func writePng(image *image.Paletted, path string) {
	f := mustFile(path)
	defer f.Close()
	err := png.Encode(f, image)
	if err != nil {
		log.Fatal(err)
	}
}

func makeBanner(asciiArtLines []string, l, h int, bgColorHex, fgColorHex string, fontPath string, fontSize, xPtFactor, yPtFactor float64) []*image.Paletted {
	var images []*image.Paletted
	d := 170
	for i := 0; i < maxLineLen(asciiArtLines)-d+1; i++ {
		img, err := setupBG(bgColorHex, l/2, h)
		if err != nil {
			log.Fatal(err)
		}
		err = drawFGText(asciiArtLines, i, i+d, img, fgColorHex, fontPath, fontSize, xPtFactor, yPtFactor)
		if err != nil {
			log.Fatal(err)
		}
		images = append(images, img)
	}
	return images
}

func makeBlink(asciiArtLines []string, l, h int, bgColorHex, fgColorHex string, fontPath string, fontSize, xPtFactor, yPtFactor float64) []*image.Paletted {
	const maxBlink = 10
	var images []*image.Paletted
	for i := 0; i < maxBlink; i++ {
		var img *image.Paletted
		var err error
		img, err = setupBG(bgColorHex, l, h)
		if err != nil {
			log.Fatal(err)
		}
		switch i % 2 {
		case 0:
			err = drawFGText(asciiArtLines, 0, 0, img, fgColorHex, fontPath, fontSize, xPtFactor, yPtFactor)
			if err != nil {
				log.Fatal(err)
			}
		default:
			// do nothing (just background)
		}
		images = append(images, img)
	}
	return images
}

func makeAlt(asciiArtLines []string, l, h int, bgColorHex, fgColorHex string, fontPath string, fontSize, xPtFactor, yPtFactor float64) []*image.Paletted {
	const maxBlink = 10
	var images []*image.Paletted
	for i := 0; i < maxBlink; i++ {
		var img *image.Paletted
		var err error
		switch i % 2 {
		case 0:
			img, err = setupBG(bgColorHex, l, h)
			if err != nil {
				log.Fatal(err)
			}
			err = drawFGText(asciiArtLines, 0, 0, img, fgColorHex, fontPath, fontSize, xPtFactor, yPtFactor)
			if err != nil {
				log.Fatal(err)
			}
		default:
			img, err = setupBG(fgColorHex, l, h)
			if err != nil {
				log.Fatal(err)
			}
			err = drawFGText(asciiArtLines, 0, 0, img, bgColorHex, fontPath, fontSize, xPtFactor, yPtFactor)
			if err != nil {
				log.Fatal(err)
			}
		}
		images = append(images, img)
	}
	return images
}

func makePng(asciiArtLines []string, l, h int, bgColorHex, fgColorHex string, fontPath string, fontSize, xPtFactor, yPtFactor float64) *image.Paletted {
	img, err := setupBG(bgColorHex, l, h)
	if err != nil {
		log.Fatal(err)
	}
	err = drawFGText(asciiArtLines, 0, 0, img, fgColorHex, fontPath, fontSize, xPtFactor, yPtFactor)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func parseHexColor(hex string) (color.RGBA, error) {
	var c color.RGBA
	var err error
	c.A = 0xff
	switch len(hex) {
	case 8:
		_, err = fmt.Sscanf(hex, "0x%02x%02x%02x", &c.R, &c.G, &c.B)
		return c, err
	case 5:
		_, err = fmt.Sscanf(hex, "0x%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
		return c, err
	default:
		return color.RGBA{}, fmt.Errorf("invalid length, must be 8 or 5")
	}
}

func imgBounds(asciiArtLines []string, fontSize, xPtFactor, yPtFactor float64) (int, int) {
	l := maxLineLen(asciiArtLines)*int(fontSize*xPtFactor) + 2*10 // +offset
	h := len(asciiArtLines)*int(fontSize*yPtFactor) + 2*30        // +offset
	return l, h
}

func maxLineLen(asciiArtLines []string) int {
	var maxLineLen int
	for i := range asciiArtLines {
		if maxLineLen >= len(asciiArtLines[i]) {
			continue
		}
		maxLineLen = len(asciiArtLines[i])
	}
	return maxLineLen
}

func prepareText(msg, figlet string) []string {
	text := figure.NewFigure(msg, figlet, true)
	text.Print()
	asciiArtLines := text.Slicify()
	for i := range asciiArtLines {
		asciiArtLines[i] = strings.Replace(asciiArtLines[i], "\t", "    ", -1) // convert tabs into spaces
	}
	return asciiArtLines
}

func loadFont(fontPath string) (*truetype.Font, error) {
	fontBytes, err := os.ReadFile(fontPath)
	if err != nil {
		return nil, err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func setupBG(bgHex string, l, h int) (*image.Paletted, error) {
	c, err := parseHexColor(bgHex)
	if err != nil {
		return nil, err
	}
	bg := image.NewUniform(c)
	finalImage := image.NewPaletted(image.Rect(0, 0, l, h), palette.Plan9)
	draw.Draw(finalImage, finalImage.Bounds(), bg, image.Pt(0, 0), draw.Src)
	return finalImage, nil
}

func drawFGText(lines []string, s, e int, bg draw.Image, fgHex, fontPath string, fontSize, xPtFactor, yPtFactor float64) error {
	c := freetype.NewContext()
	c.SetDPI(72)
	f, err := loadFont(fontPath)
	if err != nil {
		return err
	}
	c.SetFont(f)
	c.SetFontSize(fontSize)
	c.SetClip(bg.Bounds())
	c.SetDst(bg)
	fgColor, err := parseHexColor(fgHex)
	if err != nil {
		return err
	}
	fg := image.NewUniform(fgColor)
	c.SetSrc(fg)
	c.SetHinting(font.HintingNone)

	textXOffset := 10
	textYOffset := 30 + int(c.PointToFixed(fontSize)>>6) // Note shift/truncate 6 bits first

	pt := freetype.Pt(textXOffset, textYOffset)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		startX := pt.X
		switch {
		case s < e && e < len(line):
			line = line[s:e]
		case s < e && e >= len(line) && s < len(line):
			line = line[s:]
		}
		log.Print(line)
		for _, char := range line {
			_, err := c.DrawString(strings.Replace(string(char), "\r", " ", -1), pt)
			if err != nil {
				return err
			}
			pt.X += c.PointToFixed(fontSize * xPtFactor)
		}
		pt.X = startX
		pt.Y += c.PointToFixed(fontSize * yPtFactor)
	}
	return nil
}
