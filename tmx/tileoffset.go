package tmx

import "encoding/xml"

// TileOffset structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#tileoffset
type TileOffset struct {
	XMLName xml.Name `xml:"tileoffset"`
	X       int      `xml:"x,attr"` // Horizontal offset in pixels
	Y       int      `xml:"y,attr"` // Vertical offset in pixels (positive is down)

	// This element is used to specify an offset in pixels, to be applied when drawing a tile from the related
	// tileset. When not present, no offset is applied.
}
