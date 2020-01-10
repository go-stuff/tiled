package tmx

import (
	"fmt"
	"image"
	"strings"
)

// AnimationTile has fields to keep track of animations.
type AnimationTile struct {
	// FrameIndex keeps track of the current animated frame.
	FrameIndex int `xml:"-"`
	// FrameDuration keeps track of how long the current frame has been drawn.
	FrameDuration int64 `xml:"-"`
}

// Custom is for use with my own game engine.
type Custom struct {
	// Image is a custom field, it is a map of tileset images, accessed by image source path.
	Image map[string]*image.Image `xml:"-"`

	// TilesetTile is a custom field, it is a map of tile, accessed by GID.
	TilesetTile map[*Tileset]map[int]*Tile `xml:"-"`

	// AnimationTile is map to keep track of animation frames.
	AnimationTile map[*Tile]*AnimationTile `xml:"-"`
}

// NewCustom initializes a Custom structure.
func NewCustom() (*Custom, error) {
	custom := new(Custom)

	// map[Tileset.Source]*image.Image
	custom.Image = make(map[string]*image.Image)

	// map[GID]*Tile
	custom.TilesetTile = make(map[*Tileset]map[int]*Tile)

	// map[*Tile]*AnimationTile
	custom.AnimationTile = make(map[*Tile]*AnimationTile)

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
