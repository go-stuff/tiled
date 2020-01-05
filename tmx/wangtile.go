package tmx

import "encoding/xml"

// WangTile structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#wangtile
type WangTile struct {

	// Defines a Wang tile, by referring to a tile in the tileset and associating it with a certain Wang ID.

	XMLName xml.Name `xml:"wangtile"`
	TileID  int      `xml:"tileid,attr"` // The tile ID.
	WangID  int      `xml:"wangid,attr"` // The Wang ID, which is a 32-bit unsigned integer stored in the format 0xCECECECE (where each C is a corner color and each E is an edge color, from right to left clockwise, starting with the top edge)
}
