package generator

import (
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/fogleman/gg"
	"github.com/garasev/poe-item-generator/internal/domain"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Generator struct {
	dc        *gg.Context
	maxWidth  int
	maxHeight int

	font font.Face
}

func NewGenerator(fontBytes []byte) (*Generator, error) {
	fontFace, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	face, err := opentype.NewFace(fontFace, &opentype.FaceOptions{
		Size:    18,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}
	return &Generator{
		font: face,
	}, nil
}

func (g *Generator) CreateItem(item *domain.Item) image.Image {
	g.getMaxWidth(item)
	g.createHeader(item)
	dc := gg.NewContext(g.maxWidth, g.maxHeight)
	return dc.Image()
}

func (g *Generator) getMaxWidth(item *domain.Item) {
	for _, block := range item.Blocks {
		for _, stat := range block.Stats {
			w, _ := getStringSize(stat, g.font)
			if w > g.maxWidth {
				g.maxWidth = w
			}
		}
	}
}

func (g *Generator) createHeader(item *domain.Item) error {
	var header *domain.Block
	for _, block := range item.Blocks {
		if block.Type == domain.Header {
			header = block
			break
		}
	}

	var rarityFlag bool
	name := make([]string, 0, 2)
	var dc *gg.Context
	var err error

	for _, stat := range header.Stats {
		if rarityFlag {
			name = append(name, stat)
		}
		if !strings.Contains(stat, "Rarity:") {
			continue
		}

		rarityStr := strings.Split(stat, " ")
		path := rarityTypes[rarityStr[1]]
		cwd, _ := os.Getwd()
		fmt.Println(cwd)
		dc, err = generateHeader(path, g.maxWidth)
		if err != nil {
			return err
		}
		rarityFlag = true
	}
	dc.SetFontFace(g.font)
	dc.SetRGB(1, 1, 1)
	for i, n := range name {
		dc.DrawStringAnchored(
			n,
			float64(g.maxWidth)/2,
			float64(partH)*(0.25+float64(i)*0.5),
			0.5,
			0.5,
		)
	}
	dc.SavePNG("result.png")
	return nil
}
