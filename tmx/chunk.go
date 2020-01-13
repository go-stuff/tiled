package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Chunk structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#chunk
type Chunk struct {
	XMLName xml.Name `xml:"chunk"`
	X       int      `xml:"x,attr"`      // The x coordinate of the chunk in tiles.
	Y       int      `xml:"y,attr"`      // The y coordinate of the chunk in tiles.
	Width   int      `xml:"width,attr"`  // The width of the chunk in tiles.
	Height  int      `xml:"height,attr"` // The height of the chunk in tiles.

	// Can contain: <tile>
	Tile []*LayerTile `xml:"tile"`
}

func (c *Chunk) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Chunk:\n")
	fmt.Fprintf(&b, "\tX:        (%T) %q\n", c.X, c.X)
	fmt.Fprintf(&b, "\tY:        (%T) %q\n", c.Y, c.Y)
	fmt.Fprintf(&b, "\tWidth:    (%T) %q\n", c.Width, c.Width)
	fmt.Fprintf(&b, "\tHeight:   (%T) %q\n", c.Height, c.Height)

	for _, tile := range c.Tile {
		fmt.Fprintf(&b, tile.String())
	}

	return b.String()
}
