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

	tmxDir, tmxFile := filepath.Split(source)

	fmt.Println("tmxDir: ", tmxDir, ", tmxFile: ", tmxFile)

	tmxBytes, err := ioutil.ReadFile(source)
	if err != nil {
		return nil, fmt.Errorf("error reading tmx file: %w", err)
	}

	err = xml.Unmarshal(tmxBytes, &t.Map)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling tmx bytes: %w", err)
	}

	for _, tileset := range t.Map.Tileset {

		// Internal Tileset
		if tileset.Image != nil {
			imgDir, imgFile := filepath.Split(tileset.Image.Source)
			imgPath := filepath.Join(tmxDir, imgDir, imgFile)

			fmt.Println("imgDir: ", imgDir, ", imgFile: ", imgFile)
			fmt.Println("imgPath: ", imgPath)

			tileset.Image.Source = imgPath
		}

		// External Tileset
		if tileset.Source != "" {
			fmt.Println("Source: ", tileset.Source)

			tsxDir, tsxFile := filepath.Split(tileset.Source)
			tsxPath := filepath.Join(tmxDir, tsxDir, tsxFile)

			fmt.Println("tsxDir: ", tsxDir, "tsxFile: ", tsxFile)
			fmt.Println("tsxPath: ", tsxPath)

			tsxBytes, err := ioutil.ReadFile(tsxPath)
			if err != nil {
				return nil, fmt.Errorf("error reading tsx file: %w", err)
			}

			err = xml.Unmarshal(tsxBytes, &tileset)
			if err != nil {
				return nil, fmt.Errorf("error unmarshaling tsx bytes: %w", err)
			}

			tileset.Source = tsxPath

			imgDir, imgFile := filepath.Split(tileset.Image.Source)
			imgPath := filepath.Join(tmxDir, tsxDir, imgDir, imgFile)

			fmt.Println("imgDir: ", imgDir, ", imgFile: ", imgFile)
			fmt.Println("imgPath: ", imgPath)

			tileset.Image.Source = imgPath
		}

	}

	// for _, objectGroup := range t.Map.ObjectGroup {

	// 	for _, object := range objectGroup.Object {

	// 		for _, polygon := range object.Polygon {

	// 			err = polygon.decodePoints()
	// 			if err != nil {
	// 				return nil, err
	// 			}

	// 		}

	// 	}
	// }

	return t, nil
}

func (t *TMX) String() string {
	return t.Map.String()
}
