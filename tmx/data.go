package tmx

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

// Data structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#data
type Data struct {
	XMLName     xml.Name `xml:"data"`
	Raw         []byte   `xml:",innerxml"`
	Encoding    string   `xml:"encoding,attr"`    // The encoding used to encode the tile layer data. When used, it can be “base64” and “csv” at the moment.
	Compression string   `xml:"compression,attr"` // The compression used to compress the tile layer data. Tiled supports “gzip” and “zlib”.

	// When no encoding or compression is given, the tiles are stored as individual XML tile elements. Next to that,
	// the easiest format to parse is the “csv” (comma separated values) format.

	// The base64-encoded and optionally compressed layer data is somewhat more complicated to parse. First you need to
	// base64-decode it, then you may need to decompress it. Now you have an array of bytes, which should be
	// interpreted as an array of unsigned 32-bit integers using little-endian byte ordering.

	// Whatever format you choose for your layer data, you will always end up with so called “global tile IDs” (gids).
	// They are global, since they may refer to a tile from any of the tilesets used by the map. In order to find out
	// from which tileset the tile is you need to find the tileset with the highest firstgid that is still lower or
	// equal than the gid. The tilesets are always stored with increasing firstgids.

	// // Can contain: <tile>, <chunk>
	// Tiles struct {
	// 	GID []int

	// 	// Not to be confused with the tile element inside a tileset, this element defines the value of a single tile on a
	// 	// tile layer. This is however the most inefficient way of storing the tile layer data, and should generally be
	// 	// avoided.
	// } `xml:"tile,omitempty"`

	//LayerTiles *LayerTiles `xml:"layertile,omitempty"`

	// https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#tmx-tilelayer-tile
	Tiles *LayerTiles `xml:"tile,omitempty"`

	// Not to be confused with the tile element inside a tileset, this element defines the value of a single tile on a
	// tile layer. This is however the most inefficient way of storing the tile layer data, and should generally be
	// avoided.

	Chunks *Chunks `xml:"chunk,omitempty"`
}

func (d *Data) decodeBase64() ([]byte, error) {
	raw := bytes.NewReader(bytes.TrimSpace(d.Raw))
	enc := base64.NewDecoder(base64.StdEncoding, raw)

	var err error
	var reader io.Reader

	switch d.Compression {
	case "gzip":
		reader, err = gzip.NewReader(enc)
		if err != nil {
			return nil, err
		}
	case "zlib":
		reader, err = zlib.NewReader(enc)
		if err != nil {
			return nil, err
		}
	case "":
		fallthrough
	default:
		err = fmt.Errorf("unrecognized '%s' compression format: %w", d.Compression, err)
		return nil, err
	}

	return ioutil.ReadAll(reader)
}

func (d *Data) decodeCSV() error {
	// Make sure there are only digits and commas in the Data.
	// https://golang.org/pkg/strings/#Map
	// Map returns a copy of the string s with all its characters modified according to the mapping function. If
	// mapping returns a negative value, the character is dropped from the string with no replacement.
	cleanRawData := strings.Map(
		func(r rune) rune {
			if (r >= '0' && r <= '9') || (r == ',') {
				return r
			}
			return -1
		},
		string(d.Raw),
	)
	splitData := strings.Split(string(cleanRawData), ",")

	d.Tiles = new(LayerTiles)
	d.Tiles.GID = make([]int, len(splitData), len(splitData))

	for i, value := range splitData {
		gid, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		d.Tiles.GID[i] = gid
	}

	return nil
}

func (d *Data) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Data:\n")
	fmt.Fprintf(&b, "\tEncoding:    (%T) %q\n", d.Encoding, d.Encoding)
	fmt.Fprintf(&b, "\tCompression: (%T) %q\n", d.Compression, d.Compression)
	fmt.Fprintf(&b, "\tTiles.GID:   (%T) len(%d) cap(%d) %v\n", d.Tiles.GID, len(d.Tiles.GID), cap(d.Tiles.GID), d.Tiles.GID)

	return b.String()
}
