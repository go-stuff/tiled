package tmx

import "encoding/xml"

// Frame structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#frame
type Frame struct {
	XMLName  xml.Name `xml:"frame"`
	TileID   int      `xml:"tileid,attr"`   // The local ID of a tile within the parent <tileset>.
	Duration int      `xml:"duration,attr"` // How long (in milliseconds) this frame should be displayed before advancing to the next frame.
}
