package engine

type Tile struct {
	kind string
	//isVisible bool
}

func (t *Tile) GetChar() string {
	return string(terrainTable[t.kind].char)
}

func (t *Tile) IsPassable() bool {
	return terrainTable[t.kind].isPassable
}

type Board struct {
	Tiles [100][100]Tile
}

func (b *Board) Initialize() {
	for i, _ := range b.Tiles {
		for j, _ := range b.Tiles[i] {
			b.Tiles[i][j].kind = "void"
		}
	}
}

type terrain struct {
	char       rune
	isPassable bool
}

var terrainTable = map[string]terrain{
	"void": terrain{
		char:       ' ',
		isPassable: false,
	},
	"earth": terrain{
		char:       '.',
		isPassable: true,
	},
	"wall": terrain{
		char:       '\u2588',
		isPassable: false,
	},
}
