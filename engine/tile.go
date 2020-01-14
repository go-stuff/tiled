package engine

import "github.com/go-stuff/tiled/tmx"

// Tile needs to consolidate into to deal with tiles.
type Tile struct {
	Tileset  *tmx.Tileset
	Tile     *tmx.Tile
	Flipping *Flipping
}
