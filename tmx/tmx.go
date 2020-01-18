// TMX Map Format https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#tmx-map-format

// Version 1.1

// The TMX (Tile Map XML) map format used by Tiled is a flexible way to describe a tile based map. It can describe maps
// with any tile size, any amount of layers, any number of tile sets and it allows custom properties to be set on most
// elements. Beside tile layers, it can also contain groups of objects that can be placed freely.

// Note that there are many libraries and frameworks available that can work with TMX maps.

// In this document we’ll go through each element found in this map format. The elements are mentioned in the headers
// and the list of attributes of the elements are listed right below, followed by a short explanation. Attributes or
// elements that are deprecated or unsupported by the current version of Tiled are formatted in italics.

// Have a look at the changelog when you’re interested in what changed between Tiled versions.

// A DTD-file (Document Type Definition) is served at http://mapeditor.org/dtd/1.0/map.dtd. This file is not up-to-date
// but might be useful for XML-namespacing anyway.

package tmx

import (
	"path/filepath"

	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// TMX structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#tmx-map-format
type TMX struct {
	Map *Map
}

// LoadTMX loads the xml of a tmx file into a TMX struct.
func LoadTMX(source string) (*TMX, error) {
	var err error

	t := new(TMX)

	// Get tmx directory and filename and create a safe path.
	tmxDir, tmxFile := filepath.Split(source)
	tmxPath := filepath.Join(tmxDir, tmxFile)

	// Unmarshal the tmx path.
	tmxBytes, err := ioutil.ReadFile(tmxPath)
	if err != nil {
		return nil, fmt.Errorf("error reading tmx file: %w", err)
	}
	err = xml.Unmarshal(tmxBytes, &t.Map)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling tmx bytes: %w", err)
	}

	// Process each tileset in the tmx file.
	for _, tileset := range t.Map.Tileset {

		// Update any image sources that are embedded the tmx file.
		if tileset.Image != nil {

			// Get image directory and filename and create a safe path.
			imgDir, imgFile := filepath.Split(tileset.Image.Source)
			imgPath := filepath.Join(tmxDir, imgDir, imgFile)

			// Update image source with a safe path.
			tileset.Image.Source = imgPath

		}

		// External Tileset
		// Update any tileset and image sources that come from external tsx files.
		if tileset.Source != "" {

			// Get tsx directory and filename and create a safe path.
			tsxDir, tsxFile := filepath.Split(tileset.Source)
			tsxPath := filepath.Join(tmxDir, tsxDir, tsxFile)

			// Update tileset source with a safe path.
			tileset.Source = tsxPath

			// Unmarshal a tsx path.
			tsxBytes, err := ioutil.ReadFile(tsxPath)
			if err != nil {
				return nil, fmt.Errorf("error reading tsx file: %w", err)
			}
			err = xml.Unmarshal(tsxBytes, &tileset)
			if err != nil {
				return nil, fmt.Errorf("error unmarshaling tsx bytes: %w", err)
			}

			// Get image directory and filename and create a safe path.
			imgDir, imgFile := filepath.Split(tileset.Image.Source)
			imgPath := filepath.Join(tmxDir, tsxDir, imgDir, imgFile)

			// Update image source with a safe path.
			tileset.Image.Source = imgPath

		}
	}

	return t, nil
}

func (t *TMX) String() string {
	return t.Map.String()
}
