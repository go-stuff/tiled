package tmx

import "encoding/xml"

// Animation structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#animation
type Animation struct {
	XMLName xml.Name `xml:"animation"`

	// Contains a list of animation frames.

	// Each tile can have exactly one animation associated with it. In the future, there could be support for multiple named animations on a tile.

	// Can contain: <frame>
	Frame []*Frame `xml:"frame"`
}
