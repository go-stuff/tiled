package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Group structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#group
type Group struct {
	XMLName xml.Name `xml:"group"`
	ID      int      `xml:"id,attr"`      // Unique ID of the layer. Each layer that added to a map gets a unique id. Even if a layer is deleted, no layer ever gets the same ID. Can not be changed in Tiled. (since Tiled 1.2)
	Name    string   `xml:"name,attr"`    // The name of the group layer.
	OffsetX int      `xml:"offsetx,attr"` // Rendering offset of the group layer in pixels. Defaults to 0.
	OffsetY int      `xml:"offsety,attr"` // Rendering offset of the group layer in pixels. Defaults to 0.
	Opacity bool     `xml:"opacity,attr"` // The opacity of the layer as a value from 0 to 1. Defaults to 1.
	Visible bool     `xml:"visible,attr"` // Whether the layer is shown (1) or hidden (0). Defaults to 1.

	// A group layer, used to organize the layers of the map in a hierarchy. Its attributes offsetx, offsety,
	// opacity and visible recursively affect child layers.

	// Can contain: <properties>, <layer>, <objectgroup>, <imagelayer>, <group>
	Properties  *Properties    `xml:"properties"`
	Layer       []*Layer       `xml:"layer"`
	ObjectGroup []*ObjectGroup `xml:"objectgroup"`
	ImageLayer  []*ImageLayer  `xml:"imagelayer"`
	Group       []*Group       `xml:"group"`
}

func (g *Group) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Group:\n")
	fmt.Fprintf(&b, "\tID:      (%T) %q\n", g.ID, g.ID)
	fmt.Fprintf(&b, "\tName:    (%T) %q\n", g.Name, g.Name)
	fmt.Fprintf(&b, "\tOffsetX: (%T) %d\n", g.OffsetX, g.OffsetX)
	fmt.Fprintf(&b, "\tOffsetY: (%T) %d\n", g.OffsetY, g.OffsetY)
	fmt.Fprintf(&b, "\tOpacity: (%T) %t\n", g.Opacity, g.Opacity)
	fmt.Fprintf(&b, "\tVisible: (%T) %t\n", g.Visible, g.Visible)

	if g.Properties != nil {
		fmt.Fprintf(&b, g.Properties.String())
	}

	for _, layer := range g.Layer {
		fmt.Fprintf(&b, layer.String())
	}

	for _, objectGroup := range g.ObjectGroup {
		fmt.Fprintf(&b, objectGroup.String())
	}

	for _, imageLayer := range g.ImageLayer {
		fmt.Fprintf(&b, imageLayer.String())
	}

	for _, group := range g.Group {
		fmt.Fprintf(&b, group.String())
	}

	return b.String()
}
