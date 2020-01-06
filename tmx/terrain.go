package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Terrain structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#terrain
type Terrain struct {
	XMLName xml.Name `xml:"terrain"`
	Name    string   `xml:"name,attr"` // The name of the terrain type.
	Tile    int      `xml:"tile,attr"` // The local tile-id of the tile that represents the terrain visually.

	// Can contain: <properties>
	Properties *Properties `xml:"properties"`
}

func (t *Terrain) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Terrain (%s):\n", t.Name)
	fmt.Fprintf(&b, "\tName: (%T) %q\n", t.Name, t.Name)
	fmt.Fprintf(&b, "\tTile: (%T) %d\n", t.Tile, t.Tile)

	if t.Properties != nil {
		fmt.Fprintf(&b, t.Properties.String())
	}

	return b.String()
}
