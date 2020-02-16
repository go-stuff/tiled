package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// ObjectGroup structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#objectgroup
type ObjectGroup struct {
	XMLName xml.Name `xml:"objectgroup"`

	// Unique ID of the layer. Each layer that added to a map gets a unique id. Even if a layer is deleted, no layer
	// ever gets the same ID. Can not be changed in Tiled. (since Tiled 1.2)
	ID int `xml:"id,attr"`

	// The name of the object group.
	Name string `xml:"name,attr"`

	// The color used to display the objects in this group.
	Color string `xml:"color,attr,omitempty"`

	// The x coordinate of the object group in tiles. Defaults to 0 and can no longer be changed in Tiled.
	// X      float32  `xml:"x,attr"`

	// The y coordinate of the object group in tiles. Defaults to 0 and can no longer be changed in Tiled.
	// Y      float32  `xml:"y,attr"`

	// The width of the object group in tiles. Meaningless.
	// Width  float32  `xml:"width,attr"`

	// The height of the object group in tiles. Meaningless.
	// Height float32  `xml:"height,attr"`

	// The opacity of the layer as a value from 0 to 1. Defaults to 1.
	Opacity int `xml:"opacity,attr,omitempty"`

	// Whether the layer is shown (1) or hidden (0). Defaults to 1.
	Visible int `xml:"visible,attr"`

	// Rendering offset for this object group in pixels. Defaults to 0. (since 0.14)
	OffsetX int `xml:"offsetx,attr,omitempty"`

	// Rendering offset for this object group in pixels. Defaults to 0. (since 0.14)
	OffsetY int `xml:"offsety,attr,omitempty"`

	// Whether the objects are drawn according to the order of appearance (“index”) or sorted by their y-coordinate
	// (“topdown”). Defaults to “topdown”.
	DrawOrder string `xml:"draworder,attr,omitempty"`

	// The object group is in fact a map layer, and is hence called “object layer” in Tiled.

	// Can contain: <properties>, <object>
	Properties *Properties `xml:"properties,omitempty"`
	Object     []*Object   `xml:"object"`
}

func (o *ObjectGroup) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "ObjectGroup:\n")
	fmt.Fprintf(&b, "\tID:        (%T) %d\n", o.ID, o.ID)
	fmt.Fprintf(&b, "\tName:      (%T) %q\n", o.Name, o.Name)
	fmt.Fprintf(&b, "\tColor:     (%T) %q\n", o.Color, o.Color)
	fmt.Fprintf(&b, "\tOpacity:   (%T) %d\n", o.Opacity, o.Opacity)
	fmt.Fprintf(&b, "\tVisible:   (%T) %d\n", o.Visible, o.Visible)
	fmt.Fprintf(&b, "\tOffsetX:   (%T) %d\n", o.OffsetX, o.OffsetX)
	fmt.Fprintf(&b, "\tOffsetY:   (%T) %d\n", o.OffsetY, o.OffsetY)
	fmt.Fprintf(&b, "\tDrawOrder: (%T) %q\n", o.DrawOrder, o.DrawOrder)

	// for _, property := range o.Properties {
	// 	fmt.Fprintf(&b, property.String())
	// }

	if o.Properties != nil {
		fmt.Fprintf(&b, o.Properties.String())
	}

	for _, object := range o.Object {
		fmt.Fprintf(&b, object.String())
	}

	return b.String()
}
