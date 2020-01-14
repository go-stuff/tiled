package tmx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Text structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#text
type Text struct {
	XMLName    xml.Name `xml:"text"`
	FontFamily string   `xml:"fontfamily,attr"` // The font family used (default: “sans-serif”)
	PixelSize  int      `xml:"pixelsize,attr"`  // The size of the font in pixels (not using points, because other sizes in the TMX format are also using pixels) (default: 16)
	Wrap       bool     `xml:"wrap,attr"`       // Whether word wrapping is enabled (1) or disabled (0). Defaults to 0.
	Color      string   `xml:"color,attr"`      // Color of the text in #AARRGGBB or #RRGGBB format (default: #000000)
	Bold       bool     `xml:"bold,attr"`       // Whether the font is bold (1) or not (0). Defaults to 0.
	Italic     bool     `xml:"italic,attr"`     // Whether the font is italic (1) or not (0). Defaults to 0.
	Underline  bool     `xml:"underline,attr"`  // Whether a line should be drawn below the text (1) or not (0). Defaults to 0.
	Strikeout  bool     `xml:"strikeout,attr"`  // Whether a line should be drawn through the text (1) or not (0). Defaults to 0.
	Kerning    bool     `xml:"kerning,attr"`    // Whether kerning should be used while rendering the text (1) or not (0). Default to 1.
	HAlign     string   `xml:"halign,attr"`     // Horizontal alignment of the text within the object (left (default), center, right or justify (since Tiled 1.2.1))
	VAlign     string   `xml:"valign,attr"`     // Vertical alignment of the text within the object (top (default), center or bottom)

	// Used to mark an object as a text object. Contains the actual text as character data.

	// For alignment purposes, the bottom of the text is the descender height of the font, and the top of the text is
	// the ascender height of the font. For example, bottom alignment of the word “cat” will leave some space below the
	// text, even though it is unused for this word with most fonts. Similarly, top alignment of the word “cat” will
	// leave some space above the “t” with most fonts, because this space is used for diacritics.

	// If the text is larger than the object’s bounds, it is clipped to the bounds of the object.
}

func (t *Text) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Text:\n")
	fmt.Fprintf(&b, "\tFontFamily: (%T) %q\n", t.FontFamily, t.FontFamily)
	fmt.Fprintf(&b, "\tPixelSize:  (%T) %d\n", t.PixelSize, t.PixelSize)
	fmt.Fprintf(&b, "\tWrap:       (%T) %t\n", t.Wrap, t.Wrap)
	fmt.Fprintf(&b, "\tColor:      (%T) %q\n", t.Color, t.Color)
	fmt.Fprintf(&b, "\tBold:       (%T) %t\n", t.Bold, t.Bold)
	fmt.Fprintf(&b, "\tItalic:     (%T) %t\n", t.Italic, t.Italic)
	fmt.Fprintf(&b, "\tUnderline:  (%T) %t\n", t.Underline, t.Underline)
	fmt.Fprintf(&b, "\tStrikeout:  (%T) %t\n", t.Strikeout, t.Strikeout)
	fmt.Fprintf(&b, "\tKerning:    (%T) %t\n", t.Kerning, t.Kerning)
	fmt.Fprintf(&b, "\tHAlign:     (%T) %q\n", t.HAlign, t.HAlign)
	fmt.Fprintf(&b, "\tVAlign:     (%T) %q\n", t.VAlign, t.VAlign)

	return b.String()
}
