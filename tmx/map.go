package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Map constants
const (
	OrientationOrthogonal string = "orthogonal"
	OrientationIsometric  string = "isometric"
	OrientationHexagonal  string = "hexagonal"

	RenderOrderRightDown string = "right-down"
	RenderOrderRightUp   string = "right-up"
	RenderOrderLeftDown  string = "left-down"
	RenderOrderLeftUp    string = "left-up"

	StaggerIndexOdd  string = "odd"
	StaggerIndexEven string = "even"
)

// Map structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#map
type Map struct {
	XMLName xml.Name `xml:"map"`

	// The TMX format version. Was “1.0” so far, and will be incremented to match minor Tiled releases.
	Version string `xml:"version,attr"`

	// The Tiled version used to save the file (since Tiled 1.0.1). May be a date (for snapshot builds).
	TiledVersion string `xml:"tiledversion,attr"`

	// Map orientation. Tiled supports “orthogonal”, “isometric”, “staggered” and “hexagonal” (since 0.11).
	Orientation string `xml:"orientation,attr"`

	// The order in which tiles on tile layers are rendered. Valid values are right-down (the default), right-up,
	// left-down and left-up. In all cases, the map is drawn row-by-row. (only supported for orthogonal maps at the
	// moment)
	RenderOrder string `xml:"renderorder,attr"`

	// The map width in tiles.
	Width int `xml:"width,attr"`

	// The map height in tiles.
	Height int `xml:"height,attr"`

	// The width of a tile.
	TileWidth int `xml:"tilewidth,attr"`

	// The height of a tile.
	TileHeight int `xml:"tileheight,attr"`

	// Only for hexagonal maps. Determines the width or height (depending on the staggered axis) of the tile’s edge,
	// in pixels.
	HexSideLength int `xml:"hexsidelength,attr"`

	//  For staggered and hexagonal maps, determines which axis (“x” or “y”) is staggered. (since 0.11)
	StaggerAxis string `xml:"staggeraxis,attr"`

	// For staggered and hexagonal maps, determines whether the “even” or “odd” indexes along the staggered axis are
	// shifted. (since 0.11)
	StaggerIndex string `xml:"staggerindex,attr"`

	// The background color of the map. (optional, may include alpha value since 0.15 in the form #AARRGGBB)
	BackgroundColor string `xml:"backgroundcolor,attr"`

	// Stores the next available ID for new layers. This number is stored to prevent reuse of the same ID after layers
	// have been removed. (since 1.2)
	NextLayerID int `xml:"nextlayerid,attr"`

	// Stores the next available ID for new objects. This number is stored to prevent reuse of the same ID after
	// objects have been removed. (since 0.11)
	NextObjectID int `xml:"nextobjectid,attr"`

	// The tilewidth and tileheight properties determine the general grid size of the map. The individual tiles may
	// have different sizes. Larger tiles will extend at the top and right (anchored to the bottom left).

	// A map contains three different kinds of layers. Tile layers were once the only type, and are simply called
	// layer, object layers have the objectgroup tag and image layers use the imagelayer tag. The order in which
	// these layers appear is the order in which the layers are rendered by Tiled.

	// The staggered orientation refers to an isometric map using staggered axes.

	// Can contain: <properties>, <tileset>, <layer>, <objectgroup>, <imagelayer>, <group> (since 1.0)
	Content []Content `xml:",any"`

	// Properties  []*Property    `xml:"properties>property"`
	// Tileset     []*Tileset     `xml:"tileset"`
	// Layer       []*Layer       `xml:"layer"`
	// ObjectGroup []*ObjectGroup `xml:"objectgroup"`
	// ImageLayer  []*ImageLayer  `xml:"imagelayer"`
	// Group       []*Group       `xml:"group"`
}

// CountLayers recurses Map and Group Content to get a count of Layers.
func (m *Map) CountLayers(count int, content []Content) int {
	for _, c := range content {
		switch v := c.Value.(type) {
		case *Layer:
			count++
		case *ImageLayer:
			count++
		case *ObjectGroup:
			count++
		case *Group:
			count = m.CountLayers(count, v.Content)
		}
	}
	return count
}

func (m *Map) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Map:\n")
	fmt.Fprintf(&b, "\tTMX Format Version: (%T) %q\n", m.Version, m.Version)
	fmt.Fprintf(&b, "\tTiled Version:      (%T) %q\n", m.TiledVersion, m.TiledVersion)
	fmt.Fprintf(&b, "\tOrientation:        (%T) %q \n", m.Orientation, m.Orientation)
	fmt.Fprintf(&b, "\tRender Order:       (%T) %q\n", m.RenderOrder, m.RenderOrder)
	fmt.Fprintf(&b, "\tWidth:              (%T) %d\n", m.Width, m.Width)
	fmt.Fprintf(&b, "\tHeight:             (%T) %d\n", m.Height, m.Height)
	fmt.Fprintf(&b, "\tTile Width:         (%T) %d\n", m.TileWidth, m.TileWidth)
	fmt.Fprintf(&b, "\tTile Height:        (%T) %d\n", m.TileHeight, m.TileHeight)
	fmt.Fprintf(&b, "\tHex Side Length:    (%T) %d\n", m.HexSideLength, m.HexSideLength)
	fmt.Fprintf(&b, "\tStagger Axis:       (%T) %q\n", m.StaggerAxis, m.StaggerAxis)
	fmt.Fprintf(&b, "\tStagger Index:      (%T) %q\n", m.StaggerIndex, m.StaggerIndex)
	fmt.Fprintf(&b, "\tBackgroundColor:    (%T) %q\n", m.BackgroundColor, m.BackgroundColor)
	fmt.Fprintf(&b, "\tNext Layer ID:      (%T) %d\n", m.NextLayerID, m.NextLayerID)
	fmt.Fprintf(&b, "\tNext Object ID:     (%T) %d\n", m.NextObjectID, m.NextObjectID)

	for _, content := range m.Content {
		fmt.Fprintf(&b, content.String())
	}

	// for _, property := range m.Properties {
	// 	fmt.Fprintf(&b, property.String())
	// }

	// for _, tileset := range m.Tileset {
	// 	fmt.Fprintf(&b, tileset.String())
	// }

	// for _, layer := range m.Layer {
	// 	fmt.Fprintf(&b, layer.String())
	// }

	// for _, objectGroup := range m.ObjectGroup {
	// 	fmt.Fprintf(&b, objectGroup.String())
	// }

	// for _, imageLayer := range m.ImageLayer {
	// 	fmt.Fprintf(&b, imageLayer.String())
	// }

	// for _, group := range m.Group {
	// 	fmt.Fprintf(&b, group.String())
	// }

	return b.String()
}
