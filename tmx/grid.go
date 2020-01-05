package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Grid structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#grid
type Grid struct {
	XMLName     xml.Name `xml:"grid"`
	Orientation string   `xml:"orientation,attr"` //   Orientation of the grid for the tiles in this tileset (orthogonal or isometric)
	Width       int      `xml:"width,attr"`       //   Width of a grid cell
	Height      int      `xml:"height,attr"`      //   Height of a grid cell

	// This element is only used in case of isometric orientation, and determines how tile overlays for terrain and
	// collision information are rendered.
}

func (g *Grid) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Grid:\n")
	fmt.Fprintf(&b, "\tOrientation: (%T) %q\n", g.Orientation, g.Orientation)
	fmt.Fprintf(&b, "\tWidth:       (%T) %d\n", g.Width, g.Width)
	fmt.Fprintf(&b, "\tHeight:      (%T) %d\n", g.Height, g.Height)

	return b.String()
}
