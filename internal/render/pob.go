package render

import (
	"image/color"
	"log"
	"strings"

	"github.com/fogleman/gg"
	"github.com/garasev/poe-item-generator/internal/parser"
)

// ================== НАСТРОЙКИ ==================

const (
	Width = 600

	Padding = 20

	FontSize = 22

	LineSpacing   = 4
	BlockSpacing  = 10
	HeaderSpacing = 2

	LineHeight = FontSize + LineSpacing
	MaxTextW   = Width - Padding*2
)

// ================== ЦВЕТА ==================

var rarityColors = map[string]color.RGBA{
	"Normal": {200, 200, 200, 255},
	"Magic":  {120, 170, 255, 255},
	"Rare":   {255, 215, 0, 255},
	"Unique": {175, 96, 37, 255},
}

// ================== ОСНОВНАЯ ФУНКЦИЯ ==================

func RenderPoB(item parser.Item, output string) {
	// временный контекст для измерений
	tmp := gg.NewContext(Width, 10)

	fontPath := "../../src/fontin/FontinSans_Cyrillic_SC_46b.ttf"

	if err := tmp.LoadFontFace(fontPath, FontSize); err != nil {
		log.Fatal(err)
	}

	height := int(measureHeight(tmp, item))
	dc := gg.NewContext(Width, height)

	if err := dc.LoadFontFace(fontPath, FontSize); err != nil {
		log.Fatal(err)
	}

	// фон
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	// рамка
	col := rarityColors[item.Rarity]
	dc.SetColor(col)
	dc.SetLineWidth(3)
	dc.DrawRectangle(2, 2, float64(Width-4), float64(height-4))
	dc.Stroke()

	y := float64(Padding)

	// ===== NAME =====
	if item.Name != "" {
		dc.SetColor(col)
		y = drawWrapped(dc, item.Name, y)
	}

	y += HeaderSpacing

	// ===== BASE TYPE =====
	if item.BaseType != "" {
		dc.SetRGB(0.8, 0.8, 0.8)
		y = drawWrapped(dc, item.BaseType, y)
	}

	y = drawSeparator(dc, y)

	// ===== PROPERTIES =====
	for _, line := range item.Properties {
		dc.SetRGB(0.7, 0.7, 0.7)
		y = drawWrapped(dc, line, y)
	}
	if len(item.Properties) > 0 {
		y = drawSeparator(dc, y)
	}

	// ===== REQUIREMENTS =====
	for _, line := range item.Requirements {
		dc.SetRGB(0.7, 0.7, 0.7)
		y = drawWrapped(dc, line, y)
	}
	if len(item.Requirements) > 0 {
		y = drawSeparator(dc, y)
	}

	// ===== ITEM LEVEL =====
	if item.ItemLevel != "" {
		dc.SetRGB(0.7, 0.7, 0.7)
		y = drawWrapped(dc, "Item Level: "+item.ItemLevel, y)

		y = drawSeparator(dc, y)
	}

	// ===== MODS =====
	for _, line := range item.Mods {
		dc.SetRGB(0.4, 0.7, 1.0)
		y = drawWrapped(dc, line, y)
	}

	// ===== DESCRIPTION =====
	if len(item.Description) > 0 {
		y = drawSeparator(dc, y)

		for _, line := range item.Description {
			dc.SetRGB(0.6, 0.6, 0.6)
			y = drawWrapped(dc, line, y)
		}
	}

	dc.SavePNG(output)
}

// ================== ВСПОМОГАТЕЛЬНЫЕ ==================

func drawWrapped(dc *gg.Context, text string, y float64) float64 {
	lines := wrapText(dc, text, float64(MaxTextW))
	for _, line := range lines {
		dc.DrawStringAnchored(line, float64(Width)/2, y, 0.5, 0)
		y += LineHeight
	}
	return y
}

func drawSeparator(dc *gg.Context, y float64) float64 {
	dc.SetRGB(0.6, 0.6, 0.6)
	dc.SetLineWidth(1)

	// маленький отступ перед линией
	y += BlockSpacing / 2

	// рисуем линию
	dc.DrawLine(
		float64(Padding),
		y+0.5, // субпиксель для чёткости
		float64(Width-Padding),
		y+0.5,
	)
	dc.Stroke()

	// маленький отступ после линии
	y += BlockSpacing / 2

	return y
}

func wrapText(dc *gg.Context, text string, maxWidth float64) []string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{text}
	}

	var lines []string
	current := words[0]

	for _, word := range words[1:] {
		test := current + " " + word
		w, _ := dc.MeasureString(test)
		if w <= maxWidth {
			current = test
		} else {
			lines = append(lines, current)
			current = word
		}
	}
	lines = append(lines, current)
	return lines
}

func measureHeight(dc *gg.Context, item parser.Item) float64 {
	y := float64(Padding)

	addBlock := func(lines []string) {
		for _, line := range lines {
			wrapped := wrapText(dc, line, float64(MaxTextW))
			y += float64(len(wrapped)) * LineHeight
		}
		y += BlockSpacing
	}

	if item.Name != "" {
		addBlock([]string{item.Name})
	}
	if item.BaseType != "" {
		y += HeaderSpacing
		addBlock([]string{item.BaseType})
	}

	addBlock(item.Properties)
	addBlock(item.Requirements)

	if item.ItemLevel != "" {
		addBlock([]string{"Item Level: " + item.ItemLevel})
	}

	addBlock(item.Mods)
	addBlock(item.Description)

	return y + Padding
}
