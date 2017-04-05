package flags

import (
	"fmt"
	"image/gif"
	"os"
	"path/filepath"
)

type GIFFlag gif.GIF

func (f *GIFFlag) UnmarshalFlag(path string) error {
	var img *gif.GIF

	abs, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	fd, err := os.Open(abs)
	if err != nil {
		return fmt.Errorf("couldn't open file at path: %s", abs)
	}

	img, err = gif.DecodeAll(fd)
	if err != nil {
		return fmt.Errorf("error parsing file as gif image %s", err.Error())
	}

	*f = GIFFlag(*img)

	return nil

}
