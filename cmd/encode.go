package cmd

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"io/ioutil"
	"log"
	"os"

	"github.com/topherbullock/steganogify/cmd/flags"
	"github.com/topherbullock/steganogify/convert"
)

type Encode struct {
	InputGif    flags.GIFFlag `long:"input" short:"i" description:"Source gif file to use as carrier"`
	OutputFile  string        `long:"output" short:"o" default:"stegano.gif" description:"output gif file destination"`
	Message     *string       `long:"msg" short:"m" description:"plaintext message"`
	MessageFile *string       `long:"msg-file" short:"f" description:"file to read message from"`
}

func (e *Encode) Run() error {
	message, err := e.getMessage()
	if err != nil {
		log.Fatal(err)
	}

	colors := convert.PlaintextToColors(message)

	outgif := gif.GIF{
		Image: coloredImages(colors),
	}

	delays := make([]int, len(outgif.Image))
	for index, _ := range delays {
		delays[index] = 10
	}

	outgif.Delay = delays

	fmt.Print(outgif.Image[0].Bounds())

	out, err := os.Create(e.OutputFile)
	if err != nil {
		log.Fatal(err)
	}

	err = gif.EncodeAll(out, &outgif)
	if err != nil {
		fmt.Print(err)
	}

	return nil
}

func (e *Encode) getMessage() (string, error) {

	if e.Message != nil {
		return *e.Message, nil
	}

	if e.MessageFile != nil {
		bytes, err := ioutil.ReadFile(*e.MessageFile)
		return string(bytes), err
	}

	return "", errors.New("must provide a message or message file")
}

func coloredImages(colors []color.RGBA) []*image.Paletted {
	var images []*image.Paletted

	rect := image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{256, 256},
	}

	for _, colorValue := range colors {
		palette := color.Palette{colorValue}
		img := image.NewPaletted(rect, palette)
		draw.Draw(img, rect, &image.Uniform{colorValue}, image.ZP, draw.Src)
		images = append(images, img)
	}

	return images
}
