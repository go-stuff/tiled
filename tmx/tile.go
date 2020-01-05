package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Tile structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#tile
type Tile struct {
	XMLName     xml.Name `xml:"tile"`
	ID          int      `xml:"id,attr"`          // The local tile ID within its tileset.
	Type        string   `xml:"type,attr"`        // The type of the tile. Refers to an object type and is used by tile objects. (optional) (since 1.0)
	Terrain     string   `xml:"terrain,attr"`     // Defines the terrain type of each corner of the tile, given as comma-separated indexes in the terrain types array in the order top-left, top-right, bottom-left, bottom-right. Leaving out a value means that corner has no terrain. (optional)
	Probability float32  `xml:"probability,attr"` // A percentage indicating the probability that this tile is chosen when it competes with others while editing with the terrain tool. (optional)

	// Can contain: <properties>, <image> (since 0.9), <objectgroup>, <animation>
	Properties   *Properties    `xml:"properties"`
	Image        *Image         `xml:"image"`
	ObjectGroups []*ObjectGroup `xml:"objectgroup"`
	Animation    *Animation     `xml:"animation"`
}

func (t *Tile) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Tile (%d):\n", t.ID)
	fmt.Fprintf(&b, "\tID:          %d (%T)\n", t.ID, t.ID)
	fmt.Fprintf(&b, "\tType:        %q (%T)\n", t.Type, t.Type)
	fmt.Fprintf(&b, "\tTerrain:     %q (%T)\n", t.Terrain, t.Terrain)
	fmt.Fprintf(&b, "\tProbability: %f (%T)\n", t.Probability, t.Probability)

	return b.String()
}
