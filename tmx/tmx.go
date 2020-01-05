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
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"

	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path"
)

// TMX structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#tmx-map-format
type TMX struct {
	Map Map
}

// LoadTMX loads the xml of a tmx file into a TMX struct.
func LoadTMX(filepath string) (*TMX, error) {
	t := new(TMX)

	t.Map.Image = make(map[string]*image.Image)

	tmxDir, _ := path.Split(filepath)

	tmxBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error reading tmx file: %w", err)
	}

	err = xml.Unmarshal(tmxBytes, &t.Map)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling tmx bytes: %w", err)
	}

	for i := range t.Map.Tileset {
		if t.Map.Tileset[i].Source != "" {
			tsxDir, tsxFile := path.Split(t.Map.Tileset[i].Source)

			fmt.Println("tsxDir: ", tsxDir, "tsxFile: ", tsxFile)

			tsxBytes, err := ioutil.ReadFile(path.Join(tmxDir, tsxDir, tsxFile))
			if err != nil {
				return nil, fmt.Errorf("error reading tsx file: %w", err)
			}

			err = xml.Unmarshal(tsxBytes, &t.Map.Tileset[i])
			if err != nil {
				return nil, fmt.Errorf("error unmarshaling tsx bytes: %w", err)
			}

			t.Map.Tileset[i].Source = path.Join(tmxDir, tsxDir, tsxFile)

			imgDir, imgFile := path.Split(t.Map.Tileset[i].Image.Source)

			t.Map.Tileset[i].Image.Source = path.Join(tmxDir, tsxDir, imgDir, imgFile)

			if t.Map.Image[t.Map.Tileset[i].Image.Source] == nil {
				imgBytes, err := ioutil.ReadFile(t.Map.Tileset[i].Image.Source)
				if err != nil {
					return nil, fmt.Errorf("error reading image file: %w", err)
				}
				pngImage, _, err := image.Decode(bytes.NewReader(imgBytes))
				if err != nil {
					return nil, fmt.Errorf("error decoding image file: %w", err)
				}
				t.Map.Image[t.Map.Tileset[i].Image.Source] = &pngImage
			}
		}
	}

	for i := range t.Map.Layer {
		err = t.Map.Layer[i].Data.decodeCSV()
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}

func (t *TMX) String() string {
	return t.Map.String()
}

// type TMX struct {
// 	Layer     [][][]int
// 	Rectangle map[int]map[int]image.Rectangle
// }

// all paths in TMX file are relative to the TMX file
// all paths in TSX file are relative to the TSX file

// // NewMap unmarshalls an xml file into a struct.
// func NewTMX(tmxPath string) (*TMX, error) {

// 	bytes, err := ioutil.ReadFile(tmxPath)
// 	if err != nil {
// 		return nil, fmt.Errorf("error reading tmx file: %w", err)
// 	}

// 	tmxMap := &Map{}
// 	err = xml.Unmarshal(bytes, &tmxMap)
// 	if err != nil {
// 		return nil, fmt.Errorf("error unmarshaling tmx bytes: %w", err)
// 	}

// 	tmxDir, tmxFile := path.Split(tmxPath)

// 	log.Println(tmxDir, tmxFile)

// 	tmx := &TMX{
// 		Layer:     make([][][]int, len(tmxMap.Layer)),
// 		Rectangle: make(map[int]map[int]image.Rectangle),
// 	}

// 	for i, tileset := range tmxMap.Tileset {

// 		// Each Tileset has a slice of rectangles.
// 		tmx.Rectangle[i] = make(map[int]image.Rectangle)

// 		// TSX File Tileset.
// 		if tileset.Source != "" {
// 			tsxPath := ""
// 			tsxDir, tsxFile := path.Split(tileset.Source)
// 			log.Println(tsxDir, tsxFile)

// 			if tsxDir == "" {
// 				tsxPath = path.Join(tmxDir, tsxFile)
// 			}
// 			log.Println(tsxPath)

// 			bytes, err := ioutil.ReadFile(tsxPath)
// 			if err != nil {
// 				return nil, fmt.Errorf("error reading tsx file: %w", err)
// 			}
// 			tsx := &Tileset{}
// 			err = xml.Unmarshal(bytes, &tsx)
// 			if err != nil {
// 				return nil, fmt.Errorf("error unmarshaling tsx bytes: %w", err)
// 			}

// 			tileset = *tsx

// 			spew.Dump(tsx.Tilewidth)
// 			spew.Dump(tsx.Tileheight)

// 			iw, err := strconv.Atoi(tileset.Image.Width)
// 			if err != nil {
// 				return nil, err
// 			}

// 			ih, err := strconv.Atoi(tileset.Image.Height)
// 			if err != nil {
// 				return nil, err
// 			}

// 			// opts := &ebiten.DrawImageOptions{}
// 			// opts.GeoM.Translate(0, 0)
// 			// opts.GeoM.Scale(1.0, 1.0)
// 			tw, err := strconv.Atoi(tileset.Tilewidth)
// 			if err != nil {
// 				return nil, fmt.Errorf("error converting tsx tileset width: %w", err)
// 			}

// 			th, err := strconv.Atoi(tileset.Tileheight)
// 			if err != nil {
// 				return nil, fmt.Errorf("error converting tsx tileset height: %w", err)
// 			}

// 			j := 0
// 			for r := 0; r < ih*th; r += th {
// 				for c := 0; c < iw*tw; c += tw {
// 					tmx.Rectangle[i][j] = image.Rectangle{
// 						Min: image.Point{c, r},
// 						Max: image.Point{c, r}.Add(image.Point{tw, th}),
// 					}
// 					j++
// 				}
// 			}

// 			engine.Sheet[tileset.Name], err = NewSheet(
// 				"/home/stevo/code/github.com/go-stuff/game/asset/rpg-overworld-tileset v1.2 (wonderdot)/Overworld_Tileset.png",
// 				image.Point{tw, th},
// 				nil)
// 			if err != nil {
// 				return nil, err
// 			}
// 		}

// 		// Embedded Tileset.
// 		if tileset.Image.Source != "" {
// 			log.Println(tileset.Image.Source)

// 			iw, err := strconv.Atoi(tileset.Image.Width)
// 			if err != nil {
// 				return nil, err
// 			}

// 			ih, err := strconv.Atoi(tileset.Image.Height)
// 			if err != nil {
// 				return nil, err
// 			}

// 			tw, err := strconv.Atoi(tileset.Tilewidth)
// 			if err != nil {
// 				return nil, err
// 			}

// 			th, err := strconv.Atoi(tileset.Tileheight)
// 			if err != nil {
// 				return nil, err
// 			}

// 			j := 0
// 			for r := 0; r < ih*th; r += th {
// 				for c := 0; c < iw*tw; c += tw {
// 					tmx.Rectangle[i][j] = image.Rectangle{
// 						Min: image.Point{c, r},
// 						Max: image.Point{c, r}.Add(image.Point{tw, th}),
// 					}
// 					j++
// 				}
// 			}

// 			engine.Sheet[tileset.Name], err = NewSheet(
// 				//tileset.Image.Source,
// 				"/home/stevo/code/github.com/go-stuff/game/asset/rpg-overworld-tileset v1.2 (wonderdot)/TropicalExtras_Tileset.png",
// 				image.Point{iw, ih},
// 				nil)
// 			if err != nil {
// 				return nil, err
// 			}
// 		}

// 	}

// 	return tmx, nil
// }
