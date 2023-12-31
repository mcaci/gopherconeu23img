package main

import (
	"bufio"
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
	"golang.org/x/image/font"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	l := flag.Int("l", 0, "Length of the image")
	h := flag.Int("h", 0, "Height of the image")
	path := flag.String("o", "", "path of the output image/gif")
	bgColorHex := flag.String("bgHex", "0x4d3178", "Hexadecimal value for the background color")
	fgColorHex := flag.String("fgHex", "0xabc", "Hexadecimal value for the color of the text")
	fontPath := flag.String("fontPath", "fonts/Ubuntu-R.ttf", "path of the font to use")
	fontSize := flag.Float64("fontSize", 32.0, "font size of the output text in the image")
	xPtFactor := flag.Float64("xPtFactor", 0.5, "x size factor of one letter box")
	yPtFactor := flag.Float64("yPtFactor", 1.0, "y size factor of one letter box")
	figlet := flag.String("figlet", "banner", "name of the figlet font; see https://github.com/common-nighthawk/go-figure/tree/master/fonts for the values and http://www.figlet.org/examples.html for the actual effect")
	banner := flag.Bool("banner", false, "if true it's a banner gif, else it's a picture")
	blink := flag.Bool("blink", false, "if true it's a plain blinking gif, else it's a picture")
	alt := flag.Bool("alt", false, "if true it's a alternating colors blinking gif, else it's a picture")
	delay := flag.Int("delay", 0, "used with '-banner, '-blink' or '-alt', it indicates the delay between each frame of the gif")
	inFile := flag.String("inputFile", "", "Experimental: path of the file containing a list of strings for which to create the image/gif")

	flag.Parse()

	if *inFile != "" {
		f, _ := os.Open(*inFile)
		s := bufio.NewScanner(f)
		// s.Split(bufio.ScanLines)
		const n = 24
		var i int
		for s.Scan() {
			text := s.Text()
			asciiArtLines := prepareText(text, *figlet)

			text = cases.Title(language.English).String(text)
			path := defaultOutFile(text)
			log.Printf("%d/%d Working on %s", i, n, text)
			i++

			lImg := prepareSide(*l, len(asciiArtLines[0]), 2*30, *fontSize, *xPtFactor)
			hImg := prepareSide(*h, len(asciiArtLines), 2*30, *fontSize, *yPtFactor)
			switch {
			case *banner:
				images := makeBanner(asciiArtLines, lImg, hImg, *bgColorHex, *fgColorHex, *fontPath, *fontSize, *xPtFactor, *yPtFactor)
				writeGif(images, path+".gif", *delay, 5)
			case *blink:
				images := makeBlink(asciiArtLines, lImg, hImg, *bgColorHex, *fgColorHex, *fontPath, *fontSize, *xPtFactor, *yPtFactor)
				writeGif(images, path+".gif", *delay, 75)
			case *alt:
				images := makeAlt(asciiArtLines, lImg, hImg, *bgColorHex, *fgColorHex, *fontPath, *fontSize, *xPtFactor, *yPtFactor)
				writeGif(images, path+".gif", *delay, 100)
			default:
				image := makePng(asciiArtLines, lImg, hImg, *bgColorHex, *fgColorHex, *fontPath, *fontSize, *xPtFactor, *yPtFactor)
				writePng(image, path+".png")
			}
		}
		return
	}

	text := strings.Join(flag.Args(), " ")
	outPath := *path
	if *path == "" {
		outPath = defaultOutFile(text)
	}
	asciiArtLines := prepareText(text, *figlet)
	lImg := prepareSide(*l, len(asciiArtLines[0]), 2*30, *fontSize, *xPtFactor)
	hImg := prepareSide(*h, len(asciiArtLines), 2*30, *fontSize, *yPtFactor)
	switch {
	case *banner:
		images := makeBanner(asciiArtLines, lImg, hImg, *bgColorHex, *fgColorHex, *fontPath, *fontSize, *xPtFactor, *yPtFactor)
		writeGif(images, outPath+".gif", *delay, 10)
	case *blink:
		images := makeBlink(asciiArtLines, lImg, hImg, *bgColorHex, *fgColorHex, *fontPath, *fontSize, *xPtFactor, *yPtFactor)
		writeGif(images, outPath+".gif", *delay, 75)
	case *alt:
		images := makeAlt(asciiArtLines, lImg, hImg, *bgColorHex, *fgColorHex, *fontPath, *fontSize, *xPtFactor, *yPtFactor)
		writeGif(images, outPath+".gif", *delay, 100)
	default:
		image := makePng(asciiArtLines, lImg, hImg, *bgColorHex, *fgColorHex, *fontPath, *fontSize, *xPtFactor, *yPtFactor)
		writePng(image, outPath+".png")
	}
}

func defaultOutFile(text string) string {
	text = cases.Title(language.English).String(text)
	return "out/" + strings.Replace(text, " ", "", -1)
}

func makeBanner(asciiArtLines []string, l, h int, bgColorHex, fgColorHex string, fontPath string, fontSize, xPtFactor, yPtFactor float64) []*image.Paletted {
	d := int(float64(l) / fontSize)
	nFrames := len(asciiArtLines[0])
	var images []*image.Paletted
	for i := 0; i < nFrames; i += 2 {
		img, err := setupBG(bgColorHex, l/2, h)
		if err != nil {
			log.Fatal(err)
		}
		err = drawFG(asciiArtLines, i, i+d, img, fgColorHex, fontPath, fontSize, xPtFactor, yPtFactor)
		if err != nil {
			log.Fatal(err)
		}
		images = append(images, img)
	}
	return images
}

func makeBlink(asciiArtLines []string, l, h int, bgColorHex, fgColorHex string, fontPath string, fontSize, xPtFactor, yPtFactor float64) []*image.Paletted {
	const nFrames = 10
	var images []*image.Paletted
	for i := 0; i < nFrames; i++ {
		img, err := setupBG(bgColorHex, l, h)
		if err != nil {
			log.Fatal(err)
		}
		switch i % 2 {
		case 0:
			err = drawFG(asciiArtLines, 0, 0, img, fgColorHex, fontPath, fontSize, xPtFactor, yPtFactor)
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
	const nFrames = 10
	var images []*image.Paletted
	for i := 0; i < nFrames; i++ {
		var bgColor0x, fgColor0x string
		switch i % 2 {
		case 0:
			bgColor0x, fgColor0x = bgColorHex, fgColorHex // same as params
		default:
			bgColor0x, fgColor0x = fgColorHex, bgColorHex // switch back and front colors
		}
		img, err := setupBG(bgColor0x, l, h)
		if err != nil {
			log.Fatal(err)
		}
		err = drawFG(asciiArtLines, 0, 0, img, fgColor0x, fontPath, fontSize, xPtFactor, yPtFactor)
		if err != nil {
			log.Fatal(err)
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
	err = drawFG(asciiArtLines, 0, 0, img, fgColorHex, fontPath, fontSize, xPtFactor, yPtFactor)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func writeGif(images []*image.Paletted, path string, delay, defaultDelay int) {
	if delay == 0 {
		delay = defaultDelay
	}
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

func mustFile(name string) *os.File {
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func prepareText(text, figlet string) []string {
	fig := figure.NewFigure(text, figlet, true)
	asciiArtLines := fig.Slicify()
	var maxLineLen int
	for i := range asciiArtLines {
		if maxLineLen >= len(asciiArtLines[i]) {
			continue
		}
		maxLineLen = len(asciiArtLines[i])
	}
	for i := range asciiArtLines {
		asciiArtLines[i] += strings.Repeat(" ", maxLineLen-len(asciiArtLines[i]))
	}
	return asciiArtLines
}

func prepareSide(side, n, offset int, fontSize, ptFactor float64) int {
	if side != 0 {
		return side
	}
	return n*int(fontSize*ptFactor) + offset
}

func setupBG(bgHex string, l, h int) (*image.Paletted, error) {
	c, err := parseHexColor(bgHex)
	if err != nil {
		return nil, err
	}
	bg := image.NewPaletted(image.Rect(0, 0, l, h), palette.Plan9)
	draw.Draw(bg, bg.Bounds(), image.NewUniform(c), image.Pt(0, 0), draw.Src)
	return bg, nil
}

func drawFG(lines []string, s, e int, bg draw.Image, fgHex, fontPath string, fontSize, xPtFactor, yPtFactor float64) error {
	c, err := fgContext(bg, fgHex, fontPath, fontSize)
	if err != nil {
		return err
	}
	textXOffset := 10
	textYOffset := 30 + int(c.PointToFixed(fontSize)>>6) // Note shift/truncate 6 bits first

	pt := freetype.Pt(textXOffset, textYOffset)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		switch {
		case s < e && e < len(line):
			line = line[s:e]
		case s < e && e >= len(line):
			line = line[s:]
		}
		// log.Print(line)
		startX := pt.X
		for _, char := range line {
			_, err := c.DrawString(string(char), pt)
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

func fgContext(bg draw.Image, fgColorHex, fontPath string, fontSize float64) (*freetype.Context, error) {
	c := freetype.NewContext()
	fontBytes, err := os.ReadFile(fontPath)
	if err != nil {
		return nil, err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}
	c.SetFont(f)
	c.SetDPI(72)
	c.SetFontSize(fontSize)
	c.SetClip(bg.Bounds())
	c.SetDst(bg)
	fgColor, err := parseHexColor(fgColorHex)
	if err != nil {
		return nil, err
	}
	c.SetSrc(image.NewUniform(fgColor))
	c.SetHinting(font.HintingNone)
	return c, nil
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
