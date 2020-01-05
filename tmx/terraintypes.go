package tmx

import "encoding/xml"

// TerrainTypes structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#terraintypes
type TerrainTypes struct {
	XMLName xml.Name `xml:"terraintypes"`
	// This element defines an array of terrain types, which can be referenced from the terrain attribute of the tile
	// element.

	// Can contain: <terrain>
	Terrain []*Terrain `xml:"terrain"`
}
