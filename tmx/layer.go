package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Layer structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#layer
type Layer struct {
	// All <tileset> tags shall occur before the first <layer> tag so that parsers may rely on having the tilesets
	// before needing to resolve tiles.

	XMLName xml.Name `xml:"layer"`
	ID      int      `xml:"id,attr"`      // Unique ID of the layer. Each layer that added to a map gets a unique id. Even if a layer is deleted, no layer ever gets the same ID. Can not be changed in Tiled. (since Tiled 1.2)
	Name    string   `xml:"name,attr"`    // The name of the layer.
	X       int      `xml:"x,attr"`       //  The x coordinate of the layer in tiles. Defaults to 0 and can not be changed in Tiled.
	Y       int      `xml:"y,attr"`       // The y coordinate of the layer in tiles. Defaults to 0 and can not be changed in Tiled.
	Width   int      `xml:"width,attr"`   // The width of the layer in tiles. Always the same as the map width for fixed-size maps.
	Height  int      `xml:"height,attr"`  // The height of the layer in tiles. Always the same as the map height for fixed-size maps.
	Opacity float32  `xml:"opacity,attr"` // The opacity of the layer as a value from 0 to 1. Defaults to 1.
	Visible bool     `xml:"visible,attr"` // Whether the layer is shown (1) or hidden (0). Defaults to 1.
	OffsetX float32  `xml:"offsetx,attr"` // Rendering offset for this layer in pixels. Defaults to 0. (since 0.14)
	OffsetY float32  `xml:"offsety,attr"` // Rendering offset for this layer in pixels. Defaults to 0. (since 0.14)

	// Can contain: <properties>, <data>
	Properties *Properties `xml:"properties"`
	Data       *Data       `xml:"data"`
}

// GID is returned using X and Y coordinates.
func (l *Layer) GID(x, y int) (int, error) {
	if x > l.Width {
		return -1, fmt.Errorf("x is too large")
	}
	if y > l.Height {
		return -1, fmt.Errorf("y is too large")
	}
	return l.Data.Tile.GID[(y*l.Width)+x], nil
}

func (l *Layer) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Layer (%d):\n", l.ID)
	fmt.Fprintf(&b, "\tID:      (%T) %d\n", l.ID, l.ID)
	fmt.Fprintf(&b, "\tName:    (%T) %q\n", l.Name, l.Name)
	fmt.Fprintf(&b, "\tX:       (%T) %d\n", l.X, l.X)
	fmt.Fprintf(&b, "\tY:       (%T) %d\n", l.Y, l.Y)
	fmt.Fprintf(&b, "\tWidth:   (%T) %d\n", l.Width, l.Width)
	fmt.Fprintf(&b, "\tHeight:  (%T) %d\n", l.Height, l.Height)
	fmt.Fprintf(&b, "\tOpacity: (%T) %f\n", l.Opacity, l.Opacity)
	fmt.Fprintf(&b, "\tVisible: (%T) %t\n", l.Visible, l.Visible)
	fmt.Fprintf(&b, "\tOffsetX: (%T) %f\n", l.OffsetX, l.OffsetX)
	fmt.Fprintf(&b, "\tOffsetY: (%T) %f\n", l.OffsetY, l.OffsetY)

	if l.Properties != nil {
		fmt.Fprintf(&b, "Layer (%d) Properties:\n%v", l.ID, l.Properties.String())
	}

	if l.Data != nil {
		fmt.Fprintf(&b, "Layer (%d) Data:\n%v", l.ID, l.Data.String())
	}

	return b.String()
}
