package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Frame structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#frame
type Frame struct {
	XMLName xml.Name `xml:"frame"`

	// The local ID of a tile within the parent <tileset>.
	TileID int `xml:"tileid,attr"`

	// How long (in milliseconds) this frame should be displayed before advancing to the next frame.
	Duration int `xml:"duration,attr"`
}

func (f *Frame) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Frame:\n")
	fmt.Fprintf(&b, "\tTileID:   (%T) %d\n", f.TileID, f.TileID)
	fmt.Fprintf(&b, "\tDuration: (%T) %d\n", f.Duration, f.Duration)

	return b.String()
}
