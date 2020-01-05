package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// TileOffset structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#tileoffset
type TileOffset struct {
	XMLName xml.Name `xml:"tileoffset"`
	X       int      `xml:"x,attr"` // Horizontal offset in pixels
	Y       int      `xml:"y,attr"` // Vertical offset in pixels (positive is down)

	// This element is used to specify an offset in pixels, to be applied when drawing a tile from the related
	// tileset. When not present, no offset is applied.
}

func (t *TileOffset) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "TileOffset:\n")
	fmt.Fprintf(&b, "\tX:(%T) %d\n", t.X, t.X)
	fmt.Fprintf(&b, "\tY:(%T) %d\n", t.Y, t.Y)

	return b.String()
}
