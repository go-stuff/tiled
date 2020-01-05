package tmx

import "encoding/xml"

// WangCornerColor structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#wangcornercolor
type WangCornerColor struct {

	// A color that can be used to define the corner of a Wang tile.

	XMLName     xml.Name `xml:"wangcornercolor"`
	Name        string   `xml:"name,attr"`        // The name of this color.
	Color       string   `xml:"color,attr"`       // The color in #RRGGBB format (example: #c17d11).
	Tile        int      `xml:"tile,attr"`        // The tile ID of the tile representing this color.
	Probability string   `xml:"probability,attr"` // The relative probability that this color is chosen over others in case of multiple options.
}
