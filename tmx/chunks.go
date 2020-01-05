package tmx

import "encoding/xml"

// Chunks structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#chunk
type Chunks struct {
	XMLName xml.Name `xml:"chunk"`
	X       int      `xml:"x,attr"`      // The x coordinate of the chunk in tiles.
	Y       int      `xml:"y,attr"`      // The y coordinate of the chunk in tiles.
	Width   int      `xml:"width,attr"`  // The width of the chunk in tiles.
	Height  int      `xml:"height,attr"` // The height of the chunk in tiles.

	// Can contain: <tile>

	// https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#tmx-tilelayer-tile
	Tiles []int `xml:"tile"`

	// Not to be confused with the tile element inside a tileset, this element defines the value of a single tile on a
	// tile layer. This is however the most inefficient way of storing the tile layer data, and should generally be
	// avoided.
}
