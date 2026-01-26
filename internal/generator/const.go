package generator

const (
	partW = 44
	partH = 54
)

var (
	rarityTypes = map[string]RarityConfig{
		"Normal": RarityConfig{
			Path: "../../src/items/normal_",
			R:    255,
			G:    255,
			B:    255,
		},
		"Magic": RarityConfig{
			Path: "../../src/items/magic_",
			R:    135,
			G:    135,
			B:    254,
		},
		"Rare": RarityConfig{
			Path: "../../src/items/rare_",
			R:    254,
			G:    254,
			B:    118,
		},
	}
)

type RarityConfig struct {
	Path string
	R    int
	G    int
	B    int
}
