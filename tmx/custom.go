package tmx

import (
	"fmt"
	"image"
	"strings"
)

// Custom is for use with my own game engine.
type Custom struct {
	// Image is a custom field, it is a map of tileset images, accessed by image source path.
	Image map[string]*image.Image `xml:"-"`
	// TilesetTile is a custom field, it is a map of tile, accessed by GID.
	TilesetTile map[*Tileset]map[int]*Tile `xml:"-"`
}

// NewCustom initializes a Custom structure.
func NewCustom() (*Custom, error) {
	custom := new(Custom)

	// map[Tileset.Source]*image.Image
	custom.Image = make(map[string]*image.Image)

	// map[GID]*Tile
	custom.TilesetTile = make(map[*Tileset]map[int]*Tile)

	return custom, nil
}

func (c *Custom) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Custom:\n")

	if c.Image != nil {
		fmt.Fprintf(&b, "\tImage:              (%T) %v\n", c.Image, c.Image)
	}

	return b.String()
}
