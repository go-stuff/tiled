package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// WangTile structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#wangtile
type WangTile struct {

	// Defines a Wang tile, by referring to a tile in the tileset and associating it with a certain Wang ID.

	XMLName xml.Name `xml:"wangtile"`

	// The tile ID.
	TileID int `xml:"tileid,attr"`

	// The Wang ID, which is a 32-bit unsigned integer stored in the format 0xCECECECE (where each C is a corner color
	// and each E is an edge color, from right to left clockwise, starting with the top edge)
	WangID int `xml:"wangid,attr"`
}

func (w *WangTile) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "WangTile:\n")
	fmt.Fprintf(&b, "\tTileID: (%T) %d\n", w.TileID, w.TileID)
	fmt.Fprintf(&b, "\tWangID: (%T) %d\n", w.WangID, w.WangID)

	return b.String()
}
