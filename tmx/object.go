package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Object structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#object
type Object struct {
	XMLName  xml.Name `xml:"object"`
	ID       int      `xml:"id,attr"`       // Unique ID of the object. Each object that is placed on a map gets a unique id. Even if an object was deleted, no object gets the same ID. Can not be changed in Tiled. (since Tiled 0.11)
	Name     string   `xml:"name,attr"`     //  The name of the object. An arbitrary string (defaults to “”).
	Type     string   `xml:"type,attr"`     //  The type of the object. An arbitrary string (defaults to “”).
	X        float64  `xml:"x,attr"`        //  The x coordinate of the object in pixels.
	Y        float64  `xml:"y,attr"`        //  The y coordinate of the object in pixels.
	Width    float64  `xml:"width,attr"`    //  The width of the object in pixels (defaults to 0).
	Height   float64  `xml:"height,attr"`   //  The height of the object in pixels (defaults to 0).
	Rotation float64  `xml:"rotation,attr"` //  The rotation of the object in degrees clockwise around (x, y) (defaults to 0).
	GID      int      `xml:"gid,attr"`      //  A reference to a tile (optional).
	Visible  bool     `xml:"visible,attr"`  //  Whether the object is shown (1) or hidden (0). Defaults to 1.
	Template string   `xml:"template,attr"` //  A reference to a template file (optional).

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
	Properties *Properties `xml:"properties"`
	// TODO ellipse
	// TODO point
	// TODO polygon
	// TODO polyline
	// TODO text
	Image *Image `xml:"image"`
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

	if o.Properties != nil {
		fmt.Fprintf(&b, o.Properties.String())
	}

	if o.Image != nil {
		fmt.Fprintf(&b, o.Image.String())
	}

	return b.String()
}
