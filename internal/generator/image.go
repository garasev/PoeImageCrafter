package generator

import (
	"errors"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
)

func getStringSize(text string, face font.Face) (int, int) {
	tempDC := gg.NewContext(1, 1)
	tempDC.SetFontFace(face)
	textWidth, textHeight := tempDC.MeasureString(text)
	return int(textWidth), int(textHeight)
}

func generateHeader(srcPath string, outW int) (*gg.Context, error) {
	if outW <= partW*3 {
		return nil, errors.New("outW must be greater than 132 pixels")
	}

	left, err := gg.LoadImage(srcPath + "left.png")
	if err != nil {
		return nil, err
	}

	center, err := gg.LoadImage(srcPath + "center.png")
	if err != nil {
		return nil, err
	}

	right, err := gg.LoadImage(srcPath + "right.png")
	if err != nil {
		return nil, err
	}

	dc := gg.NewContext(outW, partH)

	dc.DrawImage(left, 0, 0)

	centerW := outW - partW*2
	dcenter := gg.NewContext(centerW, partH)
	scale := float64(centerW) / float64(partW)
	dcenter.Scale(scale, 1)
	dcenter.DrawImage(center, 0, 0)

	dc.DrawImage(dcenter.Image(), partW, 0)
	dc.DrawImage(right, outW-partW, 0)

	// dst := gg.NewContext(outW*2, partH*2)

	// dst.Push()
	// dst.Scale(2, 2)
	// dst.DrawImage(dc.Image(), 0, 0)
	// dst.Pop()

	return dc, nil
}
