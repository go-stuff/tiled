package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Point structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#point
type Point struct {
	XMLName xml.Name `xml:"point"`

	// Used to mark an object as a point. The existing x and y attributes are used to determine the position of the
	// point.
}

func (p *Point) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Point\n")

	return b.String()
}
