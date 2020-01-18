package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Polyline structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#polyline
type Polyline struct {
	XMLName xml.Name `xml:"polyline"`

	// A list of x,y coordinates in pixels.
	Points string `xml:"points,attr"`

	// A polyline follows the same placement definition as a polygon object.
}

func (p *Polyline) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Polyline:\n")
	fmt.Fprintf(&b, "\tPoints: (%T) %q\n", p.Points, p.Points)

	return b.String()
}
