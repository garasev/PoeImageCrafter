package domain

const (
	Header       Type = 0
	Stats        Type = 1
	Requirements Type = 2
	Sockets      Type = 3
	ItemLevel    Type = 4
	Implicits    Type = 5
	Affixes      Type = 6
	Skip         Type = 7
)

type Type int

type Item struct {
	Blocks []*Block
}

type Info struct {
	ItemClass string
	Rarity    string
	Name      string
	BaseName  string
}

type Block struct {
	Type  Type
	Stats []string
}
