package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Property structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#property
type Property struct {
	XMLName  xml.Name    `xml:"property"`
	Name     string      `xml:"name,attr"`  // The name of the property.
	Type     string      `xml:"type,attr"`  // The type of the property. Can be string (default), int, float, bool, color or file (since 0.16, with color and file added in 0.17).
	ValueXML string      `xml:"value,attr"` // The value of the property.
	Value    interface{} `xml:"-"`

	// Boolean properties have a value of either “true” or “false”.
	//
	// Color properties are stored in the format #AARRGGBB.
	//
	// File properties are stored as paths relative from the location of the map file.
	//
	// When a string property contains newlines, the current version of Tiled will write out the value as characters
	// contained inside the property element rather than as the value attribute. It is possible that a future
	// version of the TMX format will switch to always saving property values inside the element rather than as an
	// attribute.
}

func (p *Property) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Property:\n")
	fmt.Fprintf(&b, "\tName:     (%T) %q\n", p.Name, p.Name)
	fmt.Fprintf(&b, "\tType:     (%T) %q\n", p.Type, p.Type)
	fmt.Fprintf(&b, "\tValueXML: (%T) %q\n", p.ValueXML, p.ValueXML)
	fmt.Fprintf(&b, "\tValue:    (%T) %v\n", p.Value, p.Value)

	return b.String()
}
