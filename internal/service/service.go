package service

import (
	"github.com/garasev/poe-item-generator/internal/generator"
	"github.com/garasev/poe-item-generator/internal/parser"
)

type Service struct {
	generator generator.Generator
	parser    parser.Parser
}
