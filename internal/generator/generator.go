package generator

import (
	"image"
	"strings"

	"github.com/fogleman/gg"
	"github.com/garasev/poe-item-generator/internal/domain"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

type Generator struct {
	rarityConfig RarityConfig
	maxWidth     int
	maxHeight    int

	font *sfnt.Font

	header *gg.Context
	stats  *gg.Context
}

func NewGenerator(fontBytes []byte) (*Generator, error) {
	fontFace, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &Generator{
		font: fontFace,
	}, nil
}

func (g *Generator) CreateItem(item *domain.Item) image.Image {
	g.getMaxWidth(item)
	for _, block := range item.Blocks {
		switch block.Type {
		case domain.Header:
			g.createHeader(block)
		case domain.Stats:
			g.createStats(block)
		}
	}
	//g.createHeader(item)
	dc := gg.NewContext(g.maxWidth, 300)
	dc.DrawImage(g.header.Image(), 0, 0)
	dc.DrawImage(g.stats.Image(), 0, g.header.Height())
	dc.SavePNG("result1.png")
	return dc.Image()
}

func (g *Generator) getMaxWidth(item *domain.Item) {
	face, err := opentype.NewFace(g.font, &opentype.FaceOptions{
		Size:    18,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return
	}
	for _, block := range item.Blocks {
		for _, stat := range block.Stats {
			w, _ := getStringSize(stat, face)
			if w > g.maxWidth {
				g.maxWidth = w
			}
		}
	}
}

func (g *Generator) createHeader(block *domain.Block) error {
	var rarityFlag bool
	var dc *gg.Context
	var err error

	name := make([]string, 0, 2)

	for _, stat := range block.Stats {
		if rarityFlag {
			name = append(name, stat)
		}
		if !strings.Contains(stat, "Rarity:") {
			continue
		}

		rarityStr := strings.Split(stat, " ")
		g.rarityConfig = rarityTypes[rarityStr[1]]

		rarityFlag = true
	}
	dc, err = generateHeader(g.rarityConfig.Path, g.maxWidth, len(name), g.rarityConfig)
	if err != nil {
		return err
	}
	face, err := opentype.NewFace(g.font, &opentype.FaceOptions{
		Size:    float64(18 * 2 / len(name)),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	dc.SetFontFace(face)
	dc.SetRGB255(g.rarityConfig.R, g.rarityConfig.G, g.rarityConfig.B)
	if len(name) == 2 {
		for i, n := range name {
			dc.DrawStringAnchored(
				n,
				float64(g.maxWidth)/2,
				float64(partH)*(0.2+float64(i)*0.45),
				0.5,
				0.5,
			)
		}
	}
	if len(name) == 1 {
		dc.DrawStringAnchored(
			name[0],
			float64(g.maxWidth)/2,
			float64(partH)*(0.35),
			0.5,
			0.5,
		)
	}

	g.header = dc
	return nil
}

func (g *Generator) createStats(block *domain.Block) error {
	face, err := opentype.NewFace(g.font, &opentype.FaceOptions{
		Size:    18,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return err
	}
	sumH := 0
	for _, stat := range block.Stats {
		_, h := getStringSize(stat, face)
		sumH += h
	}

	dc := gg.NewContext(g.maxWidth, sumH)
	borderWidth := 2.0
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	dc.SetRGB255(g.rarityConfig.R, g.rarityConfig.G, g.rarityConfig.B)
	dc.SetLineWidth(borderWidth)

	halfBorder := borderWidth / 2
	dc.DrawRectangle(halfBorder, halfBorder,
		float64(g.maxWidth)-borderWidth,
		float64(sumH)-borderWidth)
	dc.Stroke()
	g.stats = dc
	return nil
}
