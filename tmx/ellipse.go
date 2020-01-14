package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Ellipse structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#ellipse
type Ellipse struct {
	XMLName xml.Name `xml:"ellipse"`

	// Used to mark an object as an ellipse. The existing x, y, width and height attributes are used to determine the
	// size of the ellipse.
}

func (p *Ellipse) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Ellipse\n")

	return b.String()
}
