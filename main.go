package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
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
	bgColorHex := flag.String("bgHex", "0x4d3178", "Hexadecimal value for the background color")
	l := flag.Int("l", 2400, "Length of the image")
	h := flag.Int("h", 300, "Height of the image")
	fgColorHex := flag.String("fgHex", "0xabc", "Hexadecimal value for the background color")
	fontPath := flag.String("fontPath", "/usr/share/fonts/truetype/ubuntu/Ubuntu-R.ttf", "path for the font to use")
	fontSize := flag.Float64("fontSize", 32.0, "font size of the output text in the image")
	xPtFactor := flag.Float64("xPtFactor", 0.5, "x size factor of one letter box")
	yPtFactor := flag.Float64("yPtFactor", 1.0, "y size factor of one letter box")
	imageName := flag.String("o", "examples/text.png", "name of the output image")
	figlet := flag.String("figlet", "banner", "name of the figlet font; see https://github.com/common-nighthawk/go-figure/tree/master/fonts for the values and http://www.figlet.org/examples.html for the actual effect")
	flag.Parse()

	asciiArtLines := prepareText(strings.Join(flag.Args(), " "), *figlet)

	img, err := setupBG(*bgColorHex, *l, *h)
	if err != nil {
		log.Fatal(err)
	}

	err = drawFGText(asciiArtLines, img, *fgColorHex, *fontPath, *fontSize, *xPtFactor, *yPtFactor)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(*imageName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
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

func setupBG(bgHex string, l, h int) (draw.Image, error) {
	c, err := parseHexColor(bgHex)
	if err != nil {
		return nil, err
	}
	bg := image.NewUniform(c)
	finalImage := image.NewRGBA(image.Rect(0, 0, l, h))
	draw.Draw(finalImage, finalImage.Bounds(), bg, image.Pt(0, 0), draw.Src)
	return finalImage, nil
}

func drawFGText(lines []string, bg draw.Image, fgHex, fontPath string, fontSize, xPtFactor, yPtFactor float64) error {
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
		startX := pt.X
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
