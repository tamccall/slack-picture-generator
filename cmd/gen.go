package cmd

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println(ConvertFile("/Users/seph/Pictures/Kappa.png"))
}

var colors = map[int]interface{}{
	0:   nil,
	17:  nil,
	34:  nil,
	51:  nil,
	68:  nil,
	85:  nil,
	102: nil,
	119: nil,
	136: nil,
	153: nil,
	170: nil,
	187: nil,
	204: nil,
	221: nil,
	238: nil,
	255: nil,
}

func Convert(img image.Image) string {
	xBuffer := 0
	yBuffer := 0

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	totalLen := 0

	result := ""

	for y := yBuffer; y < height-yBuffer; y++ {
		for x := xBuffer; x < width-xBuffer; x++ {
			pxColor := img.At(x, y)
			r, g, b, a := pxColor.RGBA()

			af := float32(a)
			rf := float32(r)
			gf := float32(g)
			bf := float32(b)
			wf := float32(0xffff)

			ac := af / wf
			rc := rf / wf
			gc := gf / wf
			bc := bf / wf

			avg := (rc + gc + bc) / float32(3)
			gray := int(avg * float32(255))

			var char string
			if ac <= float32(.25) {
				char = ":t:"
			} else {
				char = fmt.Sprintf(":g%d:", nearestMatch(gray))
			}
			result += char
			totalLen += len(char)
		}
		result += "\n"
	}
	result = strings.TrimSuffix(result, "\n")
	return result
}

func ConvertFile(fileIn string) string {
	fileBytes, err := ioutil.ReadFile(fileIn)
	if err != nil {
		panic(err)
	}

	fileReader := bytes.NewReader(fileBytes)
	img, err := png.Decode(fileReader)
	if err != nil {
		panic(err)
	}
	return Convert(img)
}

func nearestMatch(srcGrey int) int {
	delta := 0
	var present bool
	for {
		_, present = colors[srcGrey+delta]
		if present {
			return srcGrey + delta
		}

		_, present = colors[srcGrey-delta]
		if present {
			return srcGrey - delta
		}

		delta++
	}
	return -1
}

