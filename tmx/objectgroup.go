package tmx

import "encoding/xml"

// ObjectGroup structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#objectgroup
type ObjectGroup struct {
	XMLName   xml.Name `xml:"objectgroup"`
	ID        int      `xml:"id,attr"`        // Unique ID of the layer. Each layer that added to a map gets a unique id. Even if a layer is deleted, no layer ever gets the same ID. Can not be changed in Tiled. (since Tiled 1.2)
	Name      string   `xml:"name,attr"`      // The name of the object group.
	Color     string   `xml:"color,attr"`     // The color used to display the objects in this group.
	X         int      `xml:"x,attr"`         // The x coordinate of the object group in tiles. Defaults to 0 and can no longer be changed in Tiled.
	Y         int      `xml:"y,attr"`         // The y coordinate of the object group in tiles. Defaults to 0 and can no longer be changed in Tiled.
	Width     int      `xml:"width,attr"`     // The width of the object group in tiles. Meaningless.
	Height    int      `xml:"height,attr"`    // The height of the object group in tiles. Meaningless.
	Opacity   bool     `xml:"opacity,attr"`   // The opacity of the layer as a value from 0 to 1. Defaults to 1.
	Visible   bool     `xml:"visible,attr"`   // Whether the layer is shown (1) or hidden (0). Defaults to 1.
	OffsetX   int      `xml:"offsetx,attr"`   // Rendering offset for this object group in pixels. Defaults to 0. (since 0.14)
	OffsetY   int      `xml:"offsety,attr"`   // Rendering offset for this object group in pixels. Defaults to 0. (since 0.14)
	DrawOrder int      `xml:"draworder,attr"` // Whether the objects are drawn according to the order of appearance (“index”) or sorted by their y-coordinate (“topdown”). Defaults to “topdown”.

	// The object group is in fact a map layer, and is hence called “object layer” in Tiled.

	// Can contain: <properties>, <object>
	Properties Properties `xml:"properties"`
	Object     Object     `xml:"object"`
}
