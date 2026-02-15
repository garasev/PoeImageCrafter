package render

import (
	"image"
	"image/color"
	"log"
	"strings"

	"github.com/fogleman/gg"
	"github.com/garasev/poe-item-generator/internal/parser"
)

const (
	headerFontSize = 18
	fontSize       = 14
	padding        = 5
	miniPadding    = 4
	fontPath       = "../../src/fontin/FontinSans_Cyrillic_SC_46b.ttf"
)

var (
	propertyColor = color.RGBA{
		R: 127,
		G: 127,
		B: 127,
		A: 255,
	}
	augmentedPropertyColor = color.RGBA{
		R: 136,
		G: 136,
		B: 255,
		A: 255,
	}

	levelColor = color.RGBA{255, 255, 255, 255}
	strColor   = color.RGBA{255, 80, 80, 255}
	dexColor   = color.RGBA{80, 255, 80, 255}
	intColor   = color.RGBA{80, 160, 255, 255}

	fractureColor = color.RGBA{162, 145, 96, 255}
)

var rarityConfigs = map[string]RarityConfig{
	"Normal": RarityConfig{
		path: "../../src/items/normal_",
		color: color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
	},
	"Magic": RarityConfig{
		path: "../../src/items/magic_",
		color: color.RGBA{
			R: 135,
			G: 135,
			B: 255,
			A: 255,
		},
	},
	"Rare": RarityConfig{
		path: "../../src/items/rare_",
		color: color.RGBA{
			R: 255,
			G: 255,
			B: 118,
			A: 255,
		},
	},
	"Unique": RarityConfig{
		path: "../../src/items/uniq_",
		color: color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
	},
}

type RarityConfig struct {
	color color.RGBA
	path  string
}

func RenderPoB2(item parser.Item) image.Image {
	baseWidth := 400
	rarityConfig := rarityConfigs[item.Rarity]

	header := createHeader(item.Name, item.BaseType, baseWidth, rarityConfig)
	property := createProperties(baseWidth, item.Properties)
	requirements := createRequires(baseWidth, item.Requirements)
	sockets := createSockets(baseWidth, item.Sockets)
	itemLevel := createItemLevel(baseWidth, item.ItemLevel)
	implicits := createModLines(baseWidth, item.Implicits)
	enchants := createModLines(baseWidth, item.Enchants)
	mods := createModLines(baseWidth, item.Mods)

	sepLine := createSeparationLine(baseWidth, rarityConfig)

	height := header.Height() + padding*5

	if len(item.Properties) > 0 {
		height += property.Height() + sepLine.Height()
	}

	if len(item.Requirements) > 0 {
		height += requirements.Height()
	}

	if item.Sockets != "" {
		height += sockets.Height()
	}

	if item.ItemLevel != "" {
		height += itemLevel.Height() + sepLine.Height()
	}

	if len(item.Enchants) > 0 {
		height += enchants.Height() + sepLine.Height()
	}

	if len(item.Implicits) > 0 {
		height += implicits.Height() + sepLine.Height()
	}

	if len(item.Mods) > 0 {
		height += mods.Height()
	}

	resultImage := createBaseContext(baseWidth, height, fontSize)

	y := 0
	resultImage.DrawImage(header.Image(), 0, y)
	y += header.Height()

	// PROPERTY
	if len(item.Properties) > 0 {
		resultImage.DrawImage(property.Image(), 0, y)
		y += property.Height()
	}

	// SOCKETS
	if item.Sockets != "" {
		resultImage.DrawImage(sockets.Image(), 0, y)
		y += sockets.Height()
	}

	// LINE
	if len(item.Properties) > 0 {
		resultImage.DrawImage(sepLine.Image(), 0, y)
		y += sepLine.Height()
	}

	// ITEM LEVEL
	if item.ItemLevel != "" {
		resultImage.DrawImage(itemLevel.Image(), 0, y)
		y += itemLevel.Height()
	}

	if len(item.Requirements) > 0 {
		// REQUIREMENTS
		resultImage.DrawImage(requirements.Image(), 0, y)
		y += requirements.Height()
	}

	// LINE
	resultImage.DrawImage(sepLine.Image(), 0, y)
	y += sepLine.Height()

	// ENCHANTS
	if len(item.Enchants) > 0 {
		resultImage.DrawImage(enchants.Image(), 0, y)
		y += enchants.Height()
		// LINE
		resultImage.DrawImage(sepLine.Image(), 0, y)
		y += sepLine.Height()
	}

	// IMPLICITS
	if len(item.Implicits) > 0 {
		resultImage.DrawImage(implicits.Image(), 0, y)
		y += implicits.Height()
		// LINE
		resultImage.DrawImage(sepLine.Image(), 0, y)
		y += sepLine.Height()
	}

	// MODS
	if len(item.Mods) > 0 {
		resultImage.DrawImage(mods.Image(), 0, y)
	}

	resultImage.SavePNG("last.png")
	return resultImage.Image()
}

func createHeader(name, base string, width int, rarityConfig RarityConfig) *gg.Context {
	height := 54
	if base == "" {
		height = 27
	}

	header := gg.NewContext(width, height)
	if err := header.LoadFontFace(fontPath, headerFontSize); err != nil {
		log.Fatal(err)
	}

	left, err := gg.LoadImage(rarityConfig.path + "left.png")
	if err != nil {
		log.Fatal(err)
	}

	center, err := gg.LoadImage(rarityConfig.path + "center.png")
	if err != nil {
		log.Fatal(err)
	}

	right, err := gg.LoadImage(rarityConfig.path + "right.png")
	if err != nil {
		log.Fatal(err)
	}

	scale := float64(height) / float64(left.Bounds().Max.Y)
	s := float64(width) / float64(left.Bounds().Max.X)

	header.Push()
	header.Scale(s, scale)
	header.DrawImage(center, 0, 0)
	header.Pop()

	header.Push()
	header.Scale(scale, scale)
	header.DrawImage(left, 0, 0)
	header.Pop()

	header.Push()
	header.Scale(scale, scale)
	header.DrawImage(right, int(float64(width)/scale)-left.Bounds().Max.X, 0)
	header.Pop()

	header.SetColor(rarityConfig.color)

	if base == "" {
		header.DrawStringAnchored(name, float64(width)/2, float64(height)/2, 0.5, 0.5)

		return header
	}

	header.DrawStringAnchored(name, float64(width)/2, float64(height)*0.25, 0.5, 0.5)
	header.DrawStringAnchored(base, float64(width)/2, float64(height)*0.7, 0.5, 0.5)

	return header
}

func createProperties(width int, properties []string) *gg.Context {
	height := (len(properties))*(fontSize+miniPadding) + padding
	propertiesContext := createBaseContext(width, height, fontSize)

	for i, property := range properties {
		augmented := false
		if strings.Contains(property, " (augmented)") {
			property = strings.ReplaceAll(property, " (augmented)", "")
			augmented = true
		}
		if augmented && !strings.Contains(property, "Consumes") {
			w, _ := getStringSize(property, fontSize)
			p := strings.Split(property, ": ")
			w1, _ := getStringSize(p[0]+": ", fontSize)
			propertiesContext.SetColor(propertyColor)
			propertiesContext.DrawStringAnchored(p[0]+": ", float64(width)/2-float64(w)/2+float64(w1)/2, float64((miniPadding+fontSize)*(i+1)), 0.5, 0)
			propertiesContext.SetColor(augmentedPropertyColor)
			propertiesContext.DrawStringAnchored(p[1], float64(width)/2+float64(w)/2-float64(w-w1)/2, float64((miniPadding+fontSize)*(i+1)), 0.5, 0)
			continue
		}
		propertiesContext.SetColor(propertyColor)
		propertiesContext.DrawStringAnchored(property, float64(width)/2, float64((miniPadding+fontSize)*(i+1)), 0.5, 0)
	}

	return propertiesContext
}

func createBaseContext(width, height int, fontSize float64) *gg.Context {
	context := gg.NewContext(width, height)
	if err := context.LoadFontFace(fontPath, fontSize); err != nil {
		log.Fatal(err)
	}
	context.SetRGB(0, 0, 0)
	context.Clear()

	return context
}

func getStringSize(text string, fontSize float64) (int, int) {
	tempDC := createBaseContext(1, 1, fontSize)
	textWidth, textHeight := tempDC.MeasureString(text)
	return int(textWidth), int(textHeight)
}

func createSeparationLine(width int, rarityConfig RarityConfig) *gg.Context {
	height := 1 + padding*2
	separationLine := createBaseContext(width, height, fontSize)

	sepImage, err := gg.LoadImage(rarityConfig.path + "sep.png")
	if err != nil {
		log.Fatal(err)
	}

	separationLine.DrawImage(sepImage, width/2-sepImage.Bounds().Max.X/2, padding)

	return separationLine
}

func createRequires(width int, requirements []string) *gg.Context {
	height := fontSize + padding
	requirementsContext := createBaseContext(width, height, fontSize)

	var lvl, str, dex, int string
	for _, req := range requirements {
		s := strings.Split(req, ": ")
		if len(s) < 2 {
			continue
		}
		switch s[0] {
		case "Level":
			lvl = s[1]
		case "Str":
			str = s[1]
		case "Dex":
			dex = s[1]
		case "Int":
			int = s[1]
		}
	}
	x := float64(width) / 2
	y := float64(fontSize)

	parts := []struct {
		text  string
		color color.Color
	}{
		{"Requires Level ", propertyColor},
		{lvl, levelColor},
	}

	if str != "" {
		parts = append(parts,
			struct {
				text  string
				color color.Color
			}{", ", propertyColor},
			struct {
				text  string
				color color.Color
			}{str + " Str", strColor},
		)
	}
	if dex != "" {
		parts = append(parts,
			struct {
				text  string
				color color.Color
			}{", ", propertyColor},
			struct {
				text  string
				color color.Color
			}{dex + " Dex", dexColor},
		)
	}
	if int != "" {
		parts = append(parts,
			struct {
				text  string
				color color.Color
			}{", ", propertyColor},
			struct {
				text  string
				color color.Color
			}{int + " Int", intColor},
		)
	}

	totalWidth := 0.0
	for _, p := range parts {
		w, _ := requirementsContext.MeasureString(p.text)
		totalWidth += w
	}

	curX := x - totalWidth/2

	for _, p := range parts {
		requirementsContext.SetColor(p.color)
		requirementsContext.DrawString(p.text, curX, y)
		w, _ := requirementsContext.MeasureString(p.text)
		curX += w
	}

	return requirementsContext
}

func createSockets(width int, sockets string) *gg.Context {
	height := fontSize + padding
	socketsContext := createBaseContext(width, height, fontSize)

	x := float64(width) / 2
	y := float64(fontSize)

	parts := []struct {
		text  string
		color color.Color
	}{{"Sockets: ", propertyColor}}

	for _, symbol := range sockets {
		var socketColor color.Color
		switch symbol {
		case 'R':
			socketColor = strColor
		case 'G':
			socketColor = dexColor
		case 'B':
			socketColor = intColor
		default:
			socketColor = propertyColor
		}
		parts = append(parts,
			struct {
				text  string
				color color.Color
			}{string(symbol), socketColor},
		)
	}

	totalWidth := 0.0
	for _, p := range parts {
		w, _ := socketsContext.MeasureString(p.text)
		totalWidth += w
	}

	curX := x - totalWidth/2

	for _, p := range parts {
		socketsContext.SetColor(p.color)
		socketsContext.DrawString(p.text, curX, y)
		w, _ := socketsContext.MeasureString(p.text)
		curX += w
	}

	return socketsContext
}

func createItemLevel(width int, itemLevel string) *gg.Context {
	height := fontSize + padding
	itemLevelContext := createBaseContext(width, height, fontSize)

	parts := []struct {
		text  string
		color color.Color
	}{{"Item Level: ", propertyColor}, {itemLevel, levelColor}}

	x := float64(width) / 2
	y := float64(fontSize)
	totalWidth := 0.0
	for _, p := range parts {
		w, _ := itemLevelContext.MeasureString(p.text)
		totalWidth += w
	}

	curX := x - totalWidth/2

	for _, p := range parts {
		itemLevelContext.SetColor(p.color)
		itemLevelContext.DrawString(p.text, curX, y)
		w, _ := itemLevelContext.MeasureString(p.text)
		curX += w
	}
	return itemLevelContext
}

func createModLines(width int, mods []string) *gg.Context {
	height := (len(mods))*(fontSize+miniPadding) + padding
	modsContext := createBaseContext(width, height, fontSize)

	for i, mod := range mods {
		flag := false
		mod = strings.TrimSuffix(mod, " (implicit)")

		if strings.Contains(mod, "(fractured)") {
			mod = strings.TrimSuffix(mod, " (fractured)")
			flag = true
			modsContext.SetColor(fractureColor)
		}

		if strings.Contains(mod, "(enchant)") || strings.Contains(mod, "(crafted)") {
			mod = strings.TrimSuffix(mod, " (enchant)")
			mod = strings.TrimSuffix(mod, " (crafted)")
			flag = true
			modsContext.SetColor(levelColor)
		}

		if !flag {
			modsContext.SetColor(augmentedPropertyColor)
		}

		modsContext.DrawStringAnchored(mod, float64(width)/2, float64((miniPadding+fontSize)*(i+1)), 0.5, 0)
	}

	return modsContext
}
