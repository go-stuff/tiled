package engine

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-stuff/tiled/tmx"
)

// DecodePoints decodes a tmx object points string.
func (e *Engine) DecodePoints(pointsData string) ([]*tmx.Point, error) {

	splitPoints := strings.Split(pointsData, " ")

	points := make([]*tmx.Point, len(splitPoints), len(splitPoints))

	for i, value := range splitPoints {

		xy := strings.Split(value, ",")
		if len(xy) != 2 {
			return nil, fmt.Errorf("unexpected number of coordinates in point structure: %v in %v", len(xy), value)
		}

		x, err := strconv.ParseInt(xy[0], 10, 32)
		if err != nil {
			return nil, err
		}

		y, err := strconv.ParseInt(xy[1], 10, 32)
		if err != nil {
			return nil, err
		}

		points[i] = &tmx.Point{
			X: int(x),
			Y: int(y),
		}
	}

	return points, nil
}
