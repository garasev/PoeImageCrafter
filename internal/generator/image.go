package generator

import (
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
)

func getStringSize(text string, face font.Face) (int, int) {
	tempDC := gg.NewContext(1, 1)
	tempDC.SetFontFace(face)
	textWidth, textHeight := tempDC.MeasureString(text)
	return int(textWidth), int(textHeight)
}
