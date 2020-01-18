package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Image structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#image
type Image struct {
	XMLName xml.Name `xml:"image"`

	// Used for embedded images, in combination with a data child element. Valid values are file extensions like png,
	// gif, jpg, bmp, etc.
	// Format  string   `xml:"format,attr"`

	// Used by some versions of Tiled Java. Deprecated and unsupported by Tiled Qt.
	// ID      int      `xml:"id,attr"`

	// The reference to the tileset image file (Tiled supports most common image formats).
	Source string `xml:"source,attr"`

	// Defines a specific color that is treated as transparent (example value: “#FF00FF” for magenta). Up until Tiled
	// 0.12, this value is written out without a # but this is planned to change.
	Trans string `xml:"trans,attr"`

	// The image width in pixels (optional, used for tile index correction when the image changes)
	Width int `xml:"width,attr"`

	// The image height in pixels (optional)
	Height int `xml:"height,attr"`

	// Note that it is not currently possible to use Tiled to create maps with embedded image data, even though the
	// TMX format supports this. It is possible to create such maps using libtiled (Qt/C++) or tmxlib (Python).

	// Can contain: <data>
	// Data *Data `xml:"data"`
}

func (i *Image) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Image (%s):\n", i.Source)
	fmt.Fprintf(&b, "\tSource: (%T) %q\n", i.Source, i.Source)
	fmt.Fprintf(&b, "\tTrans:  (%T) %q\n", i.Trans, i.Trans)
	fmt.Fprintf(&b, "\tWidth:  (%T) %d\n", i.Width, i.Width)
	fmt.Fprintf(&b, "\tHeight: (%T) %d\n", i.Height, i.Height)

	return b.String()
}
