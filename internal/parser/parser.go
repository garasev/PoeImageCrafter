package parser

import (
	"strings"
)

type Item struct {
	Rarity    string
	Name      string
	BaseType  string
	ItemClass string

	Sockets      string
	Properties   []string
	Implicits    []string
	Enchants     []string
	Requirements []string
	Mods         []string
	Description  []string

	ItemLevel string
}

var (
	blacklistMods = []string{
		"Searing Exarch",
		"Eater of Worlds",
	}
)

func ParseItem(text string) Item {
	lines := splitAndTrim(text)
	blocks := splitBlocks(lines)

	item := Item{}

	// ===== HEADER =====
	header := blocks[0]
	var afterRarity []string

	for _, line := range header {
		if strings.HasPrefix(line, "Item Class:") {
			item.ItemClass = strings.TrimSpace(strings.TrimPrefix(line, "Item Class:"))
			continue
		}
		if strings.HasPrefix(line, "Rarity:") {
			item.Rarity = strings.TrimSpace(strings.TrimPrefix(line, "Rarity:"))
			continue
		}
		afterRarity = append(afterRarity, line)
	}

	// name + base type
	if len(afterRarity) >= 1 {
		item.Name = afterRarity[0]
	}
	if len(afterRarity) >= 2 {
		item.BaseType = afterRarity[1]
	}

	// ===== PROPERTIES (ТОЛЬКО второй блок) =====
	blockIndex := 1
	if len(blocks) > 1 {
		if !isRequirementsBlock(blocks[1]) && !isItemLevelBlock(blocks[1]) {
			item.Properties = blocks[1]
			blockIndex = 2
		}
	}

	for i := blockIndex; i < len(blocks); i++ {
		block := blocks[i]

		if isRequirementsBlock(block) {
			item.Requirements = block[1:]
			continue
		}

		if isImplicitsBlock(block) {
			item.Implicits = block
			continue
		}

		if isEnchantsBlock(block) {
			item.Enchants = block
			continue
		}

		if isItemLevelBlock(block) {
			item.ItemLevel = strings.TrimSpace(strings.TrimPrefix(block[0], "Item Level:"))
			continue
		}

		if isSocketsBlock(block) {
			item.Sockets = strings.TrimSpace(strings.TrimPrefix(block[0], "Sockets:"))
			continue
		}

		if isDescriptionBlock(block) {
			item.Description = block
			continue
		}

		item.Mods = append(item.Mods, filterMods(block)...)
	}

	return item
}

func isRequirementsBlock(block []string) bool {
	return len(block) > 0 && block[0] == "Requirements:"
}

func isImplicitsBlock(block []string) bool {
	return len(block) > 0 && strings.Contains(block[0], "implicit")
}

func isEnchantsBlock(block []string) bool {
	return len(block) > 0 && strings.Contains(block[0], "enchant")
}

func isItemLevelBlock(block []string) bool {
	return len(block) == 1 && strings.HasPrefix(block[0], "Item Level:")
}

func splitAndTrim(text string) []string {
	raw := strings.Split(text, "\n")
	var lines []string
	for _, l := range raw {
		l = strings.TrimSpace(l)
		if l != "" {
			lines = append(lines, l)
		}
	}
	return lines
}

func splitBlocks(lines []string) [][]string {
	var blocks [][]string
	var current []string

	for _, line := range lines {
		if line == "--------" {
			blocks = append(blocks, current)
			current = []string{}
		} else {
			current = append(current, line)
		}
	}
	if len(current) > 0 {
		blocks = append(blocks, current)
	}
	return blocks
}

func isDescriptionBlock(block []string) bool {
	if len(block) > 1 {
		return strings.Contains(block[0], "Right click") ||
			strings.Contains(block[0], "Can only") ||
			strings.Contains(block[0], "Fractured Item") ||
			strings.Contains(block[0], "Searing Exarch") ||
			strings.Contains(block[0], "Eater of Worlds") ||
			strings.Contains(block[0], "Refills")
	}
	return false
}

func filterMods(mods []string) []string {
	res := make([]string, 0, len(mods))
	for _, mod := range mods {
		flag := false
		for _, blackMod := range blacklistMods {
			if strings.Contains(mod, blackMod) {
				flag = true
				break
			}
		}

		if !flag {
			res = append(res, mod)
		}
	}

	return res
}

func isSocketsBlock(block []string) bool {
	return len(block) > 0 && strings.Contains(block[0], "Sockets:")
}
