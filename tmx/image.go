package tmx

import "encoding/xml"

// Image structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#image
type Image struct {
	XMLName xml.Name `xml:"image"`
	Format  string   `xml:"format,attr"` // Used for embedded images, in combination with a data child element. Valid values are file extensions like png, gif, jpg, bmp, etc.
	ID      int      `xml:"id,attr"`     // Used by some versions of Tiled Java. Deprecated and unsupported by Tiled Qt.
	Source  string   `xml:"source,attr"` // The reference to the tileset image file (Tiled supports most common image formats).
	Trans   string   `xml:"trans,attr"`  // Defines a specific color that is treated as transparent (example value: “#FF00FF” for magenta). Up until Tiled 0.12, this value is written out without a # but this is planned to change.
	Width   int      `xml:"width,attr"`  // The image width in pixels (optional, used for tile index correction when the image changes)
	Height  int      `xml:"height,attr"` // The image height in pixels (optional)

	// Note that it is not currently possible to use Tiled to create maps with embedded image data, even though the
	// TMX format supports this. It is possible to create such maps using libtiled (Qt/C++) or tmxlib (Python).

	// Can contain: <data>
	Data *Data `xml:"data"`
}
