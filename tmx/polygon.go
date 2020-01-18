package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Polygon structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#polygon
type Polygon struct {
	XMLName xml.Name `xml:"polygon"`

	// A list of x,y coordinates in pixels.
	Points string `xml:"points,attr"`

	// Each polygon object is made up of a space-delimited list of x,y coordinates. The origin for these coordinates is
	// the location of the parent object. By default, the first point is created as 0,0 denoting that the point will
	// originate exactly where the object is placed.
}

func (p *Polygon) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Polygon:\n")
	fmt.Fprintf(&b, "\tPoints: (%T) %q\n", p.Points, p.Points)

	return b.String()
}
