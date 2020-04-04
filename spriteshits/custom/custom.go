package custom

import (
	"encoding/json"
	"fmt"
	"image"
	"os"
	"path"

	image_packer "github.com/d7561985/image-packer"
	"github.com/d7561985/image-packer/spriteshits"
)

type shit struct {
	Manifest model
	Image    image.Image
	Type     image_packer.ImageSaveType

	outputDir string
}

func New(img, manifest *os.File, outputDir string) (spriteshits.Interface, error) {
	m := model{}
	r := json.NewDecoder(manifest)

	if err := r.Decode(&m); err != nil {
		return nil, fmt.Errorf("decode manifest file error: %w", err)
	}

	if len(m.Frames) == 0 {
		return nil, fmt.Errorf("manifest file has not frames or has bad format. File's meta: %v", m)
	}

	i, _t, err := image_packer.DecodeImage(img)
	if err != nil {
		return nil, fmt.Errorf("decode image file error: %w", err)
	}

	t := image_packer.DetectType(_t)
	if t == image_packer.ImageTypeUnknown {
		return nil, fmt.Errorf("image type %s not supported", _t)
	}

	return &shit{Manifest: m, Image: i, Type: t, outputDir: outputDir}, nil
}

func (s *shit) Process() error {
	for fileName, instance := range s.Manifest.Frames {
		if err := s.processUnit(fileName, instance); err != nil {
			return fmt.Errorf("frame %q has processing error: %w", fileName, err)
		}
	}

	return nil
}

func (s *shit) processUnit(fileName string, m frameInstance) error {
	img := image.NewRGBA(image.Rect(m.SpriteSourceSize.X, m.SpriteSourceSize.Y, m.SpriteSourceSize.W, m.SpriteSourceSize.H))

	if m.SourceSize != m.SpriteSourceSize.wh {
		return fmt.Errorf("source size and their sprite size are different. I guess there is some scate option. Not supported")
	}

	// i don'y check any transformations
	for x := 0; x <= m.SourceSize.W; x++ {
		for y := 0; y <= m.SourceSize.H; y++ {
			img.Set(x, y, s.Image.At(m.Frame.X+x, m.Frame.Y+y))
		}
	}

	return image_packer.SaveImage(img, s.Type, path.Join(s.outputDir, fileName))
}
