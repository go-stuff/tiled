package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Tile structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#tile
type Tile struct {
	XMLName xml.Name `xml:"tile"`

	// The local tile ID within its tileset.
	ID int `xml:"id,attr"`

	// The type of the tile. Refers to an object type and is used by tile objects. (optional) (since 1.0)
	Type string `xml:"type,attr"`

	// Defines the terrain type of each corner of the tile, given as comma-separated indexes in the terrain types array
	// in the order top-left, top-right, bottom-left, bottom-right. Leaving out a value means that corner has no
	// terrain. (optional)
	Terrain string `xml:"terrain,attr"`

	// A percentage indicating the probability that this tile is chosen when it competes with others while editing with
	// the terrain tool. (optional)
	Probability float32 `xml:"probability,attr"`

	// Can contain: <properties>, <image> (since 0.9), <objectgroup>, <animation>
	Properties  []*Property    `xml:"properties>property"`
	Image       *Image         `xml:"image"`
	ObjectGroup []*ObjectGroup `xml:"objectgroup"`
	Animation   *Animation     `xml:"animation"`
}

func (t *Tile) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Tile:\n")
	fmt.Fprintf(&b, "\tID:          (%T) %d\n", t.ID, t.ID)
	fmt.Fprintf(&b, "\tType:        (%T) %q\n", t.Type, t.Type)
	fmt.Fprintf(&b, "\tTerrain:     (%T) %q\n", t.Terrain, t.Terrain)
	fmt.Fprintf(&b, "\tProbability: (%T) %f\n", t.Probability, t.Probability)

	for _, property := range t.Properties {
		fmt.Fprintf(&b, property.String())
	}

	if t.Image != nil {
		fmt.Fprintf(&b, t.Image.String())
	}

	for _, objectGroup := range t.ObjectGroup {
		fmt.Fprintf(&b, objectGroup.String())
	}

	if t.Animation != nil {
		fmt.Fprintf(&b, t.Animation.String())
	}

	return b.String()
}
