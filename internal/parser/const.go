package parser

import "github.com/garasev/poe-item-generator/internal/domain"

const (
	blockSplitter = "--------\n"
)

var (
	blockTypes = map[string]domain.Type{
		"Item Class":           domain.Header,
		"Quality":              domain.Stats,
		"Currently has":        domain.Stats,
		"augmented":            domain.Stats,
		"Requirements":         domain.Requirements,
		"Item Level:":          domain.ItemLevel,
		"Sockets":              domain.Sockets,
		"implicit":             domain.Implicits,
		"Right click to drink": domain.Skip,
	}
)
