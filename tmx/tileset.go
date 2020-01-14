package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Tileset structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#tileset
type Tileset struct {
	XMLName    xml.Name `xml:"tileset"`
	FirstGID   int      `xml:"firstgid,attr"`   // The first global tile ID of this tileset (this global ID maps to the first tile in this tileset).
	Source     string   `xml:"source,attr"`     // If this tileset is stored in an external TSX (Tile Set XML) file, this attribute refers to that file. That TSX file has the same structure as the <tileset> element described here. (There is the firstgid attribute missing and this source attribute is also not there. These two attributes are kept in the TMX map, since they are map specific.)
	Name       string   `xml:"name,attr"`       // The name of this tileset.
	TileWidth  int      `xml:"tilewidth,attr"`  // The (maximum) width of the tiles in this tileset.
	TileHeight int      `xml:"tileheight,attr"` // The (maximum) height of the tiles in this tileset.
	Spacing    int      `xml:"spacing,attr"`    // The spacing in pixels between the tiles in this tileset (applies to the tileset image).
	Margin     int      `xml:"margin,attr"`     // The margin around the tiles in this tileset (applies to the tileset image).
	TileCount  int      `xml:"tilecount,attr"`  // The number of tiles in this tileset (since 0.13)
	Columns    int      `xml:"columns,attr"`    // The number of tile columns in the tileset. For image collection tilesets it is editable and is used when displaying the tileset. (since 0.15)

	// If there are multiple <tileset> elements, they are in ascending order of their firstgid attribute. The first
	// tileset always has a firstgid value of 1. Since Tiled 0.15, image collection tilesets do not necessarily
	// number their tiles consecutively since gaps can occur when removing tiles.

	// Can contain: <tileoffset>, <grid> (since 1.0), <properties>, <image>, <terraintypes>, <tile>, <wangsets>
	// (since 1.1)
	TileOffset   *TileOffset   `xml:"tileoffset"`
	Grid         *Grid         `xml:"grid"`
	Properties   []*Property   `xml:"properties>property"`
	Image        *Image        `xml:"image"`
	TerrainTypes *TerrainTypes `xml:"terraintypes"`
	Tile         []*Tile       `xml:"tile"`
	Wangsets     *Wangsets     `xml:"wangsets"`
}

func (t *Tileset) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Tileset:\n")
	fmt.Fprintf(&b, "\tFirstGID:   (%T) %d\n", t.FirstGID, t.FirstGID)
	fmt.Fprintf(&b, "\tSource:     (%T) %q\n", t.Source, t.Source)
	fmt.Fprintf(&b, "\tName:       (%T) %q\n", t.Name, t.Name)
	fmt.Fprintf(&b, "\tTileWidth:  (%T) %d\n", t.TileWidth, t.TileWidth)
	fmt.Fprintf(&b, "\tTileHeight: (%T) %d\n", t.TileHeight, t.TileHeight)
	fmt.Fprintf(&b, "\tSpacing:    (%T) %d\n", t.Spacing, t.Spacing)
	fmt.Fprintf(&b, "\tMargin:     (%T) %d\n", t.Margin, t.Margin)
	fmt.Fprintf(&b, "\tTileCount:  (%T) %d\n", t.TileCount, t.TileCount)
	fmt.Fprintf(&b, "\tColumns:    (%T) %d\n", t.Columns, t.Columns)

	if t.TileOffset != nil {
		fmt.Fprintf(&b, t.TileOffset.String())
	}

	if t.Grid != nil {
		fmt.Fprintf(&b, t.Grid.String())
	}

	for _, property := range t.Properties {
		fmt.Fprintf(&b, property.String())
	}

	if t.Image != nil {
		fmt.Fprintf(&b, t.Image.String())
	}

	if t.TerrainTypes != nil {
		fmt.Fprintf(&b, t.TerrainTypes.String())
	}

	for _, tile := range t.Tile {
		fmt.Fprintf(&b, tile.String())
	}

	return b.String()
}
