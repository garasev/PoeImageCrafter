package parser

import (
	"strings"

	"github.com/garasev/poe-item-generator/internal/domain"
)

type Parser struct{}

func ParseText(text string) *domain.Item {
	blocks := strings.Split(text, blockSplitter)
	if len(blocks) <= 1 {
		blocks = strings.Split(text, blockSplitter2)
	}
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
	if len(stats) <= 1 {
		stats = strings.Split(block, "\r")
	}
	res := make([]string, 0, len(stats))
	for _, stat := range stats {
		stat = strings.TrimSuffix(stat, "\r")
		stat = strings.TrimSuffix(stat, "\n")
		if stat == "" {
			continue
		}
		res = append(res, stat)
	}
	return &domain.Block{Type: resType, Stats: res}
}
