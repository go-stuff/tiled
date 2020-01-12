package tmx

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

// Polygon structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#polygon
type Polygon struct {
	XMLName   xml.Name `xml:"polygon"`
	RawPoints string   `xml:"points,attr"`

	// Each polygon object is made up of a space-delimited list of x,y coordinates. The origin for these coordinates is
	// the location of the parent object. By default, the first point is created as 0,0 denoting that the point will
	// originate exactly where the object is placed.

	Points []*Point
}

func (p *Polygon) decodePoints() error {

	rawPoints := strings.Split(p.RawPoints, " ")

	//p.Points = make([]*Point, len(rawPoints), len(rawPoints))

	for _, rawPoint := range rawPoints {

		xy := strings.Split(rawPoint, ",")
		if len(xy) != 2 {
			return fmt.Errorf("unexpected number of coordinates in point destructure: %v in %v", len(xy), rawPoint)
		}

		x, err := strconv.ParseInt(xy[0], 10, 32)
		if err != nil {
			return err
		}

		y, err := strconv.ParseInt(xy[1], 10, 32)
		if err != nil {
			return err
		}

		// p.Points[i] = &Point{
		// 	X: int(x),
		// 	Y: int(y),
		// }

		p.Points = append(p.Points, &Point{
			X: int(x),
			Y: int(y),
		})
	}

	return nil
}

func (p *Polygon) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Polygon:\n")
	fmt.Fprintf(&b, "\tRaw Points: (%T) %q\n", p.RawPoints, p.RawPoints)

	for _, point := range p.Points {
		fmt.Fprintf(&b, point.String())
	}

	return b.String()
}
