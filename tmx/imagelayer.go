package tmx

import "encoding/xml"

// ImageLayer structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#imagelayer
type ImageLayer struct {
	XMLName xml.Name `xml:"imagelayer"`
	ID      int      `xml:"id,attr"`      // Unique ID of the layer. Each layer that added to a map gets a unique id. Even if a layer is deleted, no layer ever gets the same ID. Can not be changed in Tiled. (since Tiled 1.2)
	Name    string   `xml:"name,attr"`    //  The name of the image layer.
	OffsetX int      `xml:"offsetx,attr"` // Rendering offset of the image layer in pixels. Defaults to 0. (since 0.15)
	OffsetY int      `xml:"offsety,attr"` // Rendering offset of the image layer in pixels. Defaults to 0. (since 0.15)
	X       int      `xml:"x,attr"`       // The x position of the image layer in pixels. (deprecated since 0.15)
	Y       int      `xml:"y,attr"`       // The y position of the image layer in pixels. (deprecated since 0.15)
	Opacity bool     `xml:"opacity,attr"` // The opacity of the layer as a value from 0 to 1. Defaults to 1.
	Visible bool     `xml:"visible,attr"` // Whether the layer is shown (1) or hidden (0). Defaults to 1.

	// A layer consisting of a single image.

	// Can contain: <properties>, <image>
	Properties Properties `xml:"properties"`
	Image      Image      `xml:"image"`
}
