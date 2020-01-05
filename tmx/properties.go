package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Properties structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#properties
type Properties struct {
	XMLName xml.Name `xml:"properties"`

	// Can contain: <property>
	Property []*Property `xml:"property"`

	// Wraps any number of custom properties. Can be used as a child of the map, tileset, tile (when part of a
	// tileset), terrain, layer, objectgroup, object, imagelayer and group elements.
}

func (p *Properties) String() string {
	var b strings.Builder

	for i := range p.Property {
		fmt.Fprintf(&b, p.Property[i].String())
		fmt.Fprintf(&b, "\n")
	}

	return b.String()
}
