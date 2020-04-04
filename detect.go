package image_packer

import (
	"fmt"
	"image"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

func DecodeImage(f *os.File) (image.Image, string, error) {
	decode, s, err := image.Decode(f)
	if err != nil {
		return nil, "", fmt.Errorf("decode image error: %w", err)
	}

	return decode, s, nil
}
