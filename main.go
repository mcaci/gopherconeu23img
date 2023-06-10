package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func main() {
	textContent := "text"
	fontSize := 32.0
	fgColorHex := 0xff
	bgColorHex := 0xff

	img, err := generateImage(textContent, fgColorHex, bgColorHex, fontSize)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("text.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func generateImage(textContent string, fgColorHex, bgColorHex int, fontSize float64) (image.Image, error) {

	fgColor := color.RGBA{0xff, 0xff, 0xff, 0xff}
	// if len(fgColorHex) == 7 {
	// 	_, err := fmt.Sscanf(fgColorHex, "#%02x%02x%02x", &fgColor.R, &fgColor.G, &fgColor.B)
	// 	if err != nil {
	// 		log.Println(err)
	// 		fgColor = color.RGBA{0x2e, 0x34, 0x36, 0xff}
	// 	}
	// }

	bgColor := color.RGBA{0x30, 0x0a, 0x24, 0xff}
	// if len(bgColorHex) == 7 {
	// 	_, err := fmt.Sscanf(bgColorHex, "#%02x%02x%02x", &bgColor.R, &bgColor.G, &bgColor.B)
	// 	if err != nil {
	// 		log.Println(err)
	// 		bgColor = color.RGBA{0x30, 0x0a, 0x24, 0xff}
	// 	}
	// }

	loadedFont, err := loadFont()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	code := strings.Replace(textContent, "\t", "    ", -1) // convert tabs into spaces
	text := strings.Split(code, "\n")                      // split newlines into arrays

	fg := image.NewUniform(fgColor)
	bg := image.NewUniform(bgColor)
	rgba := image.NewRGBA(image.Rect(0, 0, 1200, 630))
	draw.Draw(rgba, rgba.Bounds(), bg, image.Pt(0, 0), draw.Src)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(loadedFont)
	c.SetFontSize(fontSize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	c.SetHinting(font.HintingNone)

	textXOffset := 50
	textYOffset := 10 + int(c.PointToFixed(fontSize)>>6) // Note shift/truncate 6 bits first

	pt := freetype.Pt(textXOffset, textYOffset)
	for _, s := range text {
		_, err = c.DrawString(strings.Replace(s, "\r", "", -1), pt)
		if err != nil {
			return nil, err
		}
		pt.Y += c.PointToFixed(fontSize * 1.5)
	}

	return rgba, nil
}

func loadFont() (*truetype.Font, error) {
	fontFile := "/usr/share/fonts/truetype/ubuntu/Ubuntu-R.ttf"
	fontBytes, err := os.ReadFile(fontFile)
	if err != nil {
		return nil, err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}
	return f, nil
}
