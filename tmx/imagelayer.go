package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// ImageLayer structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#imagelayer
type ImageLayer struct {
	XMLName xml.Name `xml:"imagelayer"`

	// Unique ID of the layer. Each layer that added to a map gets a unique id. Even if a layer is deleted, no layer
	// ever gets the same ID. Can not be changed in Tiled. (since Tiled 1.2)
	ID int `xml:"id,attr"`

	// The name of the image layer.
	Name string `xml:"name,attr"`

	// Rendering offset of the image layer in pixels. Defaults to 0. (since 0.15)
	OffsetX int `xml:"offsetx,attr,omitempty"`

	// Rendering offset of the image layer in pixels. Defaults to 0. (since 0.15)
	OffsetY int `xml:"offsety,attr,omitempty"`

	// The x position of the image layer in pixels. (deprecated since 0.15)
	// X       int      `xml:"x,attr"`

	// The y position of the image layer in pixels. (deprecated since 0.15)
	// Y       int      `xml:"y,attr"`

	// The opacity of the layer as a value from 0 to 1. Defaults to 1.
	Opacity int `xml:"opacity,attr,omitempty"`

	// Whether the layer is shown (1) or hidden (0). Defaults to 1.
	Visible int `xml:"visible,attr"`

	Locked int `xml:"locked,attr"`

	// A layer consisting of a single image.

	// Can contain: <properties>, <image>
	Properties *Properties `xml:"properties,omitempty"`
	Image      *Image      `xml:"image,omitempty"`
}

func (i *ImageLayer) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "ImageLayer:\n")
	fmt.Fprintf(&b, "\tID:      (%T) %d\n", i.ID, i.ID)
	fmt.Fprintf(&b, "\tName:    (%T) %q\n", i.Name, i.Name)
	fmt.Fprintf(&b, "\tOffsetX: (%T) %d\n", i.OffsetX, i.OffsetX)
	fmt.Fprintf(&b, "\tOffsetY: (%T) %d\n", i.OffsetY, i.OffsetY)
	fmt.Fprintf(&b, "\tOpacity: (%T) %d\n", i.Opacity, i.Opacity)
	fmt.Fprintf(&b, "\tVisible: (%T) %d\n", i.Visible, i.Visible)

	// for _, property := range i.Properties. {
	// 	fmt.Fprintf(&b, property.String())
	// }

	if i.Properties != nil {
		fmt.Fprintf(&b, i.Properties.String())
	}

	if i.Image != nil {
		fmt.Fprintf(&b, i.Image.String())
	}

	return b.String()
}
