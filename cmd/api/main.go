package main

import (
	"bytes"
	"image"
	"image/png"

	"golang.design/x/clipboard"

	"github.com/garasev/poe-item-generator/internal/parser"
	"github.com/garasev/poe-item-generator/internal/render"
)

const (
	a = ``
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	text := clipboard.Read(clipboard.FmtText)

	item := parser.ParseItem(string(text))
	image := render.RenderPoB2(item)
	copyImageToClipboard(image)
}

func copyImageToClipboard(img image.Image) error {
	err := clipboard.Init()
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		return err
	}

	clipboard.Write(clipboard.FmtImage, buf.Bytes())
	return nil
}
