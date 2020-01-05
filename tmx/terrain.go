package tmx

import "encoding/xml"

// Terrain structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#terrain
type Terrain struct {
	XMLName xml.Name `xml:"terrain"`
	Name    string   `xml:"name,attr"` // The name of the terrain type.
	Tile    int      `xml:"tile,attr"` // The local tile-id of the tile that represents the terrain visually.

	// Can contain: <properties>
	Properties Properties `xml:"properties"`
}
