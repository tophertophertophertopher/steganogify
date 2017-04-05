package convert

import (
	"encoding/hex"
	"image/color"
	"log"
	"math"
)

// PlaintextToColors hex encodes the colors and converts it to an ordered  slice
// of RGBA colors
func PlaintextToColors(text string) (colors []color.RGBA) {
	chunks := chunked(hexEncode(text))

	for i, chunk := range chunks {
		chunk = chunks[i]
		var encodedColor color.RGBA
		switch len(chunk) {
		case 1:
			encodedColor = color.RGBA{chunk[0], 0, 0, 255}
		case 2:
			encodedColor = color.RGBA{chunk[0], chunk[1], 0, 255}
		default:
			encodedColor = color.RGBA{chunk[0], chunk[1], chunk[2], 255}
		}

		colors = append(colors, encodedColor)
	}

	return colors
}

func ColorsToPlaintext(colors []color.Color) string {
	bytes := []byte{}

	for _, colorValue := range colors {
		r, g, b, _ := colorValue.RGBA()
		if r != 0 {
			bytes = append(bytes, byte(r))
		} else {
			break
		}
		if g != 0 {
			bytes = append(bytes, byte(g))
		} else {
			break
		}
		if b != 0 {
			bytes = append(bytes, byte(b))
		} else {
			break
		}
	}

	return hexDecode(bytes)
}

func hexEncode(plaintext string) []byte {
	text := []byte(plaintext)
	hexEncoded := make([]byte, hex.EncodedLen(len(text)))
	hex.Encode(hexEncoded, text)
	return hexEncoded
}

func hexDecode(src []byte) string {
	s := string(src)
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return string(decoded)
}

func chunked(source []byte) [][]byte {
	var chunk []byte

	var chunks [][]byte
	var remainder = len(source)
	var chunkCount = int(math.Floor(float64(remainder / 3)))

	for i := 0; i < chunkCount; i++ {
		chunk, source = source[0:3], source[3:remainder]
		remainder -= 3
		chunks = append(chunks, chunk)
	}

	//append remainder
	if len(source) > 0 {
		chunks = append(chunks, source)
	}

	return chunks
}
