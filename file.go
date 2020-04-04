package image_packer

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

type ImageSaveType int

const (
	ImageTypeUnknown ImageSaveType = iota
	ImageTypePng
	ImageTypeJpeg
)

func DetectType(t string) ImageSaveType {
	switch strings.ToLower(t) {
	case "png":
		return ImageTypePng
	case "jpg", "jpeg":
		return ImageTypeJpeg
	default:
		return ImageTypeUnknown
	}
}

func SaveImage(img image.Image, t ImageSaveType, dest string) error {
	f, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("open file %q error: %w", dest, err)
	}

	defer f.Close()

	switch t {
	case ImageTypePng:
		if err := png.Encode(f, img); err != nil {
			return fmt.Errorf("png encode error: %w", err)
		}
	case ImageTypeJpeg:
		if err := jpeg.Encode(f, img, &jpeg.Options{Quality: 100}); err != nil {
			return fmt.Errorf("jpeg encode error: %w", err)
		}
	default:
		return fmt.Errorf("img format (%d) not supported", t)
	}

	log.Printf("save file: %s\n", dest)

	return nil
}
