package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Terrain structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#terrain
type Terrain struct {
	XMLName xml.Name `xml:"terrain"`

	// The name of the terrain type.
	Name string `xml:"name,attr"`

	// The local tile-id of the tile that represents the terrain visually.
	Tile int `xml:"tile,attr"`

	// Can contain: <properties>
	Properties []*Property `xml:"properties>property"`
}

func (t *Terrain) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Terrain (%s):\n", t.Name)
	fmt.Fprintf(&b, "\tName: (%T) %q\n", t.Name, t.Name)
	fmt.Fprintf(&b, "\tTile: (%T) %d\n", t.Tile, t.Tile)

	for _, property := range t.Properties {
		fmt.Fprintf(&b, property.String())
	}

	return b.String()
}
