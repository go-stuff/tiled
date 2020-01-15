package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Wangsets structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#wangsets
type Wangsets struct {
	XMLName xml.Name `xml:"wangsets"`

	// Contains the list of Wang sets defined for this tileset.

	// Can contain: <wangset>
	Wangset []Wangset `xml:"wangset"`
}

func (w *Wangsets) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Wangsets:\n")

	for _, wangset := range w.Wangset {
		fmt.Fprintf(&b, wangset.String())
	}

	return b.String()
}
