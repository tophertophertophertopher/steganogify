package cmd

import (
	"fmt"
	"image/color"

	"github.com/topherbullock/steganogify/cmd/flags"
	"github.com/topherbullock/steganogify/convert"
)

type Decode struct {
	InputGif   flags.GIFFlag `long:"input" short:"i" description:"Source file containing randart"`
	OutputFile string        `long:"output" short:"o" default:"randart.jpg" description:"output file destination"`
}

func (e *Decode) Run() error {
	gif := e.InputGif
	colors := make([]color.Color, len(gif.Image))

	for index, img := range gif.Image {
		cid := img.ColorIndexAt(0, 0)
		colors[index] = img.Palette[cid]
	}

	fmt.Println(colors)

	str := convert.ColorsToPlaintext(colors)
	fmt.Println(str)

	return nil

}
