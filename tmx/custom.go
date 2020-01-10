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

// UpdateAnimationTile updates the animation frame on each animated tile.
func (c *Custom) UpdateAnimationTile(milliseconds int64) {
	for tile, animaionTile := range c.AnimationTile {
		// Update the Milliseconds of this Particular Tile Frame
		animaionTile.FrameDuration += milliseconds

		// If the elapsed time is greater than the current frame duration, go to the next frame index and
		// reset time.
		if animaionTile.FrameDuration >
			int64(tile.Animation.Frame[animaionTile.FrameIndex].Duration) {
			animaionTile.FrameIndex++
			animaionTile.FrameDuration = 0
		}

		// If the current frame equals the last frame start over.
		if animaionTile.FrameIndex == len(tile.Animation.Frame) {
			animaionTile.FrameIndex = 0
			animaionTile.FrameDuration = 0
		}
	}
}

// Return Row, Column and GID or Return image.Rectangle
// // Get the Tileset of the current GID.
// tileset, err := layer.GIDTileset(gid, l.TMX.Map.Tileset)
// if err != nil {
// 	return err
// }

// // Get the real GID by subtracting the tileset firtst GID.
// gid -= tileset.FirstGID

// if l.TMX.Custom.TilesetTile[tileset][gid] != nil {
// 	tile := l.TMX.Custom.TilesetTile[tileset][gid]
// 	if tile.Animation != nil {
// 		gid = tile.Animation.Frame[l.TMX.Custom.AnimationTile[tile].FrameIndex].TileID
// 	}
// }

// // From the GID and Tileset calculate the row and column.
// row = int(gid / tileset.Columns)
// column = (gid % tileset.Columns)

// image.Rectangle{
// 	Min: image.Point{column * tileset.TileWidth, row * tileset.TileHeight},
// 	Max: image.Point{(column * tileset.TileWidth) + tileset.TileWidth, (row * tileset.TileHeight) + tileset.TileHeight},
// }

func (c *Custom) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Custom:\n")

	if c.Image != nil {
		fmt.Fprintf(&b, "\tImage: (%T) %v\n", c.Image, c.Image)
	}

	return b.String()
}
