package engine

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-stuff/tiled/tmx"
)

// GID is returned using X and Y coordinates.
func (e *Engine) GID(layer *tmx.Layer, x, y int) (int, error) {
	if x > layer.Width {
		return -1, fmt.Errorf("x is too large")
	}
	if y > layer.Height {
		return -1, fmt.Errorf("y is too large")
	}
	//return layer.Data.Tile.GID[(y*layer.Width)+x], nil
	return 1, nil
}

// func (d *Data) decodeBase64() ([]byte, error) {
// 	raw := bytes.NewReader(bytes.TrimSpace(d.Raw))
// 	enc := base64.NewDecoder(base64.StdEncoding, raw)

// 	var err error
// 	var reader io.Reader

// 	switch d.Compression {
// 	case "gzip":
// 		reader, err = gzip.NewReader(enc)
// 		if err != nil {
// 			return nil, err
// 		}
// 	case "zlib":
// 		reader, err = zlib.NewReader(enc)
// 		if err != nil {
// 			return nil, err
// 		}
// 	case "":
// 		fallthrough
// 	default:
// 		err = fmt.Errorf("unrecognized '%s' compression format: %w", d.Compression, err)
// 		return nil, err
// 	}

// 	return ioutil.ReadAll(reader)
// }

func (e *Engine) DecodeData(data *tmx.Data) ([]int, error) {

	// No value for encoding means it is XML.
	if data.Encoding == "" {
		data.Encoding = "xml"
	}

	switch data.Encoding {
	// case "xml": TODO
	// case "base64": TODO
	case "csv":
		// continue to decode
	default:
		return []int{}, fmt.Errorf("unhandled encoding: %s", data.Encoding)
	}

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
		string(data.Raw),
	)
	splitData := strings.Split(string(cleanRawData), ",")

	// d.Tile = new(LayerTile)
	// d.Tile.GID = make([]int, len(splitData), len(splitData))
	gids := make([]int, len(splitData), len(splitData))

	for i, value := range splitData {
		gid, err := strconv.Atoi(value)
		if err != nil {
			return []int{}, err
		}
		gids[i] = gid
	}

	return gids, nil
}
