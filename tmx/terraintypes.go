package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// TerrainTypes structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#terraintypes
type TerrainTypes struct {
	XMLName xml.Name `xml:"terraintypes"`
	// This element defines an array of terrain types, which can be referenced from the terrain attribute of the tile
	// element.

	// Can contain: <terrain>
	Terrain []*Terrain `xml:"terrain"`
}

func (t *TerrainTypes) String() string {
	var b strings.Builder

	for _, terrain := range t.Terrain {
		fmt.Fprintf(&b, terrain.String())
	}

	return b.String()
}
