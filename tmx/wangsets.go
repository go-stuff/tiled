package tmx

import "encoding/xml"

// Wangsets structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#wangsets
type Wangsets struct {
	XMLName xml.Name `xml:"wangsets"`

	// Contains the list of Wang sets defined for this tileset.

	// Can contain: <wangset>
	Wangset []Wangset `xml:"wangset"`
}
