package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Animation structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#animation
type Animation struct {
	XMLName xml.Name `xml:"animation"`

	// Contains a list of animation frames.

	// Each tile can have exactly one animation associated with it. In the future, there could be support for multiple named animations on a tile.

	// Can contain: <frame>
	Frame []*Frame `xml:"frame"`
}

func (a *Animation) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Animation:\n")

	for i := range a.Frame {
		fmt.Fprintf(&b, a.Frame[i].String())
	}

	return b.String()
}
