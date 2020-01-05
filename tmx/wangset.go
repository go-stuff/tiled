package tmx

import "encoding/xml"

// Wangset structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#wangset
type Wangset struct {

	// Defines a list of corner colors and a list of edge colors, and any number of Wang tiles using these colors.

	XMLName xml.Name `xml:"wangset"`
	Name    string   `xml:"name,attr"` // The name of the Wang set.
	Tile    int      `xml:"tile,attr"` // The tile ID of the tile representing this Wang set.

	// Can contain: <wangcornercolor>, <wangedgecolor>, <wangtile>
	WangCornerColor WangCornerColor `xml:"wangcornercolor"`
	WangEdgeColor   WangEdgeColor   `xml:"wangedgecolor"`
	WangTile        WangTile        `xml:"wangtile"`
}
