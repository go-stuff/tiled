package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Wangset structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#wangset
type Wangset struct {

	// Defines a list of corner colors and a list of edge colors, and any number of Wang tiles using these colors.

	XMLName xml.Name `xml:"wangset"`

	// The name of the Wang set.
	Name string `xml:"name,attr"`

	// The tile ID of the tile representing this Wang set.
	Tile int `xml:"tile,attr"`

	// Can contain: <wangcornercolor>, <wangedgecolor>, <wangtile>
	WangCornerColor WangCornerColor `xml:"wangcornercolor"`
	WangEdgeColor   WangEdgeColor   `xml:"wangedgecolor"`
	WangTile        WangTile        `xml:"wangtile"`
}

func (w *Wangset) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Wangset:\n")
	fmt.Fprintf(&b, w.WangCornerColor.String())
	fmt.Fprintf(&b, w.WangEdgeColor.String())
	fmt.Fprintf(&b, w.WangTile.String())

	return b.String()
}
