package generator

import (
	"image"

	"github.com/fogleman/gg"
	"github.com/garasev/poe-item-generator/internal/domain"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Generator struct {
	dc        *gg.Context
	maxWidth  int
	maxHeight int

	Font font.Face
}

func NewGenerator(fontBytes []byte) (*Generator, error) {
	fontFace, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	face, err := opentype.NewFace(fontFace, &opentype.FaceOptions{
		Size:    48,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}
	return &Generator{
		Font: face,
	}, nil
}

func (g *Generator) CreateItem(item *domain.Item) image.Image {
	g.getMaxWidth(item)

	dc := gg.NewContext(g.maxWidth, g.maxHeight)
	return dc.Image()
}

func (g *Generator) getMaxWidth(item *domain.Item) {
	for _, block := range item.Blocks {
		for _, stat := range block.Stats {
			w, _ := getStringSize(stat, g.Font)
			if w > g.maxWidth {
				g.maxWidth = w
			}
		}
	}
}

func (g *Generator) createHeader(item *domain.Item) {
	for _, block := range item.Blocks {
		if block.Type != domain.Header {
			continue
		}
		dc := gg.NewContext(g.maxWidth, g.maxHeight)
		break
	}
}
