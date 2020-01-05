package tmx

import "encoding/xml"

// Chunk structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#chunk
type Chunk struct {
	XMLName xml.Name `xml:"chunk"`
	X       int      `xml:"x,attr"`      // The x coordinate of the chunk in tiles.
	Y       int      `xml:"y,attr"`      // The y coordinate of the chunk in tiles.
	Width   int      `xml:"width,attr"`  // The width of the chunk in tiles.
	Height  int      `xml:"height,attr"` // The height of the chunk in tiles.

	// Can contain: <tile>

	Tile *LayerTile `xml:"tile"`
}
