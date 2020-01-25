package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Object structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#object
type Object struct {
	XMLName xml.Name `xml:"object"`

	// Unique ID of the object. Each object that is placed on a map gets a unique id. Even if an object was deleted, no
	// object gets the same ID. Can not be changed in Tiled. (since Tiled 0.11)
	ID int `xml:"id,attr"`

	// The name of the object. An arbitrary string (defaults to “”).
	Name string `xml:"name,attr"`

	// The type of the object. An arbitrary string (defaults to “”).
	Type string `xml:"type,attr"`

	// The x coordinate of the object in pixels.
	X float64 `xml:"x,attr"`

	// The y coordinate of the object in pixels.
	Y float64 `xml:"y,attr"`

	// The width of the object in pixels (defaults to 0).
	Width float64 `xml:"width,attr"`

	// The height of the object in pixels (defaults to 0).
	Height float64 `xml:"height,attr"`

	// The rotation of the object in degrees clockwise around (x, y) (defaults to 0).
	Rotation float32 `xml:"rotation,attr"`

	// A reference to a tile (optional).
	GID int `xml:"gid,attr"`

	// Whether the object is shown (1) or hidden (0). Defaults to 1.
	Visible bool `xml:"visible,attr"`

	// A reference to a template file (optional).
	Template string `xml:"template,attr"`

	// While tile layers are very suitable for anything repetitive aligned to the tile grid, sometimes you want to
	// annotate your map with other information, not necessarily aligned to the grid. Hence the objects have their
	// coordinates and size in pixels, but you can still easily align that to the grid when you want to.

	// You generally use objects to add custom information to your tile map, such as spawn points, warps, exits, etc.

	// When the object has a gid set, then it is represented by the image of the tile with that global ID. The image
	// alignment currently depends on the map orientation. In orthogonal orientation it’s aligned to the bottom-left
	// while in isometric it’s aligned to the bottom-center. The image will rotate around the bottom-left or
	// bottom-center, respectively.

	// When the object has a template set, it will borrow all the properties from the specified template, properties
	// saved with the object will have higher priority, i.e. they will override the template properties.

	// Can contain: <properties>, <ellipse> (since 0.9), <point>, <polygon>, <polyline>, <text> (since 1.0), image
	Properties []*Property `xml:"properties>property"`
	Ellipse    []*Ellipse  `xml:"ellipse"`
	Point      []*Point    `xml:"point"`
	Polygon    []*Polygon  `xml:"polygon"`
	Polyline   []*Polyline `xml:"polyline"`
	Text       []*Text     `xml:"text"`
	Image      *Image      `xml:"image"`
}

func (o *Object) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Object:\n")
	fmt.Fprintf(&b, "\tID:       (%T) %d\n", o.ID, o.ID)
	fmt.Fprintf(&b, "\tName:     (%T) %q\n", o.Name, o.Name)
	fmt.Fprintf(&b, "\tType:     (%T) %q\n", o.Type, o.Type)
	fmt.Fprintf(&b, "\tX:        (%T) %f\n", o.X, o.X)
	fmt.Fprintf(&b, "\tY:        (%T) %f\n", o.Y, o.Y)
	fmt.Fprintf(&b, "\tWidth:    (%T) %f\n", o.Width, o.Width)
	fmt.Fprintf(&b, "\tHeight:   (%T) %f\n", o.Height, o.Height)
	fmt.Fprintf(&b, "\tRotation: (%T) %f\n", o.Rotation, o.Rotation)
	fmt.Fprintf(&b, "\tGID:      (%T) %d\n", o.GID, o.GID)
	fmt.Fprintf(&b, "\tVisible:  (%T) %t\n", o.Visible, o.Visible)
	fmt.Fprintf(&b, "\tTemplate: (%T) %q\n", o.Template, o.Template)

	for _, property := range o.Properties {
		fmt.Fprintf(&b, property.String())
	}

	for _, ellipse := range o.Ellipse {
		fmt.Fprintf(&b, ellipse.String())
	}

	for _, point := range o.Point {
		fmt.Fprintf(&b, point.String())
	}

	for _, polygon := range o.Polygon {
		fmt.Fprintf(&b, polygon.String())
	}

	for _, polyline := range o.Polyline {
		fmt.Fprintf(&b, polyline.String())
	}

	for _, text := range o.Text {
		fmt.Fprintf(&b, text.String())
	}

	if o.Image != nil {
		fmt.Fprintf(&b, o.Image.String())
	}

	// Set as rectangle if all other options are nil.
	if o.Ellipse == nil &&
		o.Point == nil &&
		o.Polygon == nil &&
		o.Polyline == nil &&
		o.Text == nil {
		fmt.Fprintf(&b, "Rectangle\n")
	}

	return b.String()
}
