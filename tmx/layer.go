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

	// Unique ID of the layer. Each layer that added to a map gets a unique id. Even if a layer is deleted, no layer
	// ever gets the same ID. Can not be changed in Tiled. (since Tiled 1.2)
	ID int `xml:"id,attr"`

	// The name of the layer.
	Name string `xml:"name,attr"`

	//  The x coordinate of the layer in tiles. Defaults to 0 and can not be changed in Tiled.
	X int `xml:"x,attr"`

	// The y coordinate of the layer in tiles. Defaults to 0 and can not be changed in Tiled.
	Y int `xml:"y,attr"`

	// The width of the layer in tiles. Always the same as the map width for fixed-size maps.
	Width int `xml:"width,attr"`

	// The height of the layer in tiles. Always the same as the map height for fixed-size maps.
	Height int `xml:"height,attr"`

	// The opacity of the layer as a value from 0 to 1. Defaults to 1.
	Opacity float32 `xml:"opacity,attr"`

	// Whether the layer is shown (1) or hidden (0). Defaults to 1.
	Visible bool `xml:"visible,attr"`

	// Rendering offset for this layer in pixels. Defaults to 0. (since 0.14)
	OffsetX float32 `xml:"offsetx,attr"`

	// Rendering offset for this layer in pixels. Defaults to 0. (since 0.14)
	OffsetY float32 `xml:"offsety,attr"`

	// Can contain: <properties>, <data>
	Properties []*Property `xml:"properties>property"`
	Data       *Data       `xml:"data"`
}

func (l *Layer) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Layer:\n")
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

	for _, property := range l.Properties {
		fmt.Fprintf(&b, property.String())
	}

	if l.Data != nil {
		fmt.Fprintf(&b, l.Data.String())
	}

	return b.String()
}
