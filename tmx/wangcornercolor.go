package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// WangCornerColor structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#wangcornercolor
type WangCornerColor struct {

	// A color that can be used to define the corner of a Wang tile.

	XMLName xml.Name `xml:"wangcornercolor"`

	// The name of this color.
	Name string `xml:"name,attr"`

	// The color in #RRGGBB format (example: #c17d11).
	Color string `xml:"color,attr"`

	// The tile ID of the tile representing this color.
	Tile int `xml:"tile,attr"`

	// The relative probability that this color is chosen over others in case of multiple options.
	Probability string `xml:"probability,attr"`
}

func (w *WangCornerColor) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "WangCornerColor:\n")
	fmt.Fprintf(&b, "\tName:        (%T) %q\n", w.Name, w.Name)
	fmt.Fprintf(&b, "\tColor:       (%T) %q\n", w.Color, w.Color)
	fmt.Fprintf(&b, "\tTile:        (%T) %d\n", w.Tile, w.Tile)
	fmt.Fprintf(&b, "\tProbability: (%T) %q\n", w.Probability, w.Probability)

	return b.String()
}
