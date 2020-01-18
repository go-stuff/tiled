package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Data structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#data
type Data struct {
	XMLName  xml.Name `xml:"data"`
	InnerXML string   `xml:",innerxml"`

	// The encoding used to encode the tile layer data. When used, it can be “base64” and “csv” at the moment.
	Encoding string `xml:"encoding,attr"`

	// The compression used to compress the tile layer data. Tiled supports “gzip” and “zlib”.
	Compression string `xml:"compression,attr"`

	// When no encoding or compression is given, the tiles are stored as individual XML tile elements. Next to that,
	// the easiest format to parse is the “csv” (comma separated values) format.

	// The base64-encoded and optionally compressed layer data is somewhat more complicated to parse. First you need to
	// base64-decode it, then you may need to decompress it. Now you have an array of bytes, which should be
	// interpreted as an array of unsigned 32-bit integers using little-endian byte ordering.

	// Whatever format you choose for your layer data, you will always end up with so called “global tile IDs” (gids).
	// They are global, since they may refer to a tile from any of the tilesets used by the map. In order to find out
	// from which tileset the tile is you need to find the tileset with the highest firstgid that is still lower or
	// equal than the gid. The tilesets are always stored with increasing firstgids.

	// Can contain: <tile>, <chunk>
	Tile  []*LayerTile `xml:"tile"`
	Chunk []*Chunk     `xml:"chunk"`
}

func (d *Data) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Data:\n")
	fmt.Fprintf(&b, "\tEncoding:    (%T) %q\n", d.Encoding, d.Encoding)
	fmt.Fprintf(&b, "\tCompression: (%T) %q\n", d.Compression, d.Compression)
	fmt.Fprintf(&b, "\tInnerXML:    (%T) %q\n", d.InnerXML, d.InnerXML)

	for _, tile := range d.Tile {
		fmt.Fprintf(&b, tile.String())
	}

	for _, chunk := range d.Chunk {
		fmt.Fprintf(&b, chunk.String())
	}

	return b.String()
}
