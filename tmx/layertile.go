package tmx

import (
	"fmt"
	"strings"
)

// LayerTile structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#tmx-tilelayer-tile
type LayerTile struct {
	GID []int // The global tile ID (default: 0).

	// Not to be confused with the tile element inside a tileset, this element defines the value of a single tile on a
	// tile layer. This is however the most inefficient way of storing the tile layer data, and should generally be
	// avoided.
}

func (t *LayerTile) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "LayerTile:\n")
	fmt.Fprintf(&b, "\tGID: (%T) %q\n", t.GID, t.GID)

	return b.String()
}
