package parser

import (
	"strings"

	"github.com/garasev/poe-item-generator/internal/domain"
)

type Parser struct{}

func ParseText(text string) *domain.Item {
	blocks := strings.Split(text, blockSplitter)
	itemBlocks := make([]*domain.Block, 0, len(blocks))
	for _, block := range blocks {
		itemBlocks = append(itemBlocks, parseBlock(block))
	}

	return &domain.Item{Blocks: itemBlocks}
}

func parseBlock(block string) *domain.Block {
	resType := domain.Affixes
	for subString, blockType := range blockTypes {
		if strings.Contains(block, subString) {
			resType = blockType
			break
		}
	}

	stats := strings.Split(block, "\n")
	if stats[len(stats)-1] == "" {
		stats = stats[:len(stats)-1]
	}
	return &domain.Block{Type: resType, Stats: stats}
}
