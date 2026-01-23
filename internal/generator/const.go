package generator

const (
	partW = 44
	partH = 54
)

var (
	rarityTypes = map[string]string{
		"Normal": "../../src/items/normal_",
		"Magic":  "../../src/items/magic_",
		"Rare":   "../../src/items/rare_",
	}

	rarityColor = map[string]string{}
)
