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

	// FrameIndex is a custom field to keep track of the current animated frame.
	FrameIndex int `xml:"-"`
	// FrameDuration is a custom field to keep track of how long the current frame index as been drawn.
	FrameDuration int64 `xml:"-"`
}

func (a *Animation) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Animation:\n")
	fmt.Fprintf(&b, "\tFrame Index:    (%T) %d\n", a.FrameIndex, a.FrameIndex)
	fmt.Fprintf(&b, "\tFrame Duration: (%T) %d\n", a.FrameDuration, a.FrameDuration)

	for i := range a.Frame {
		fmt.Fprintf(&b, a.Frame[i].String())
	}

	return b.String()
}
