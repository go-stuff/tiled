package tmx

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Content is of any value using interface{}, to preserver the order of the XML elements.
type Content struct {
	Type  string
	Value interface{}
}

// UnmarshalXML is called by Unmarshal to produce the value from the XML element.
func (c *Content) UnmarshalXML(decoder *xml.Decoder, startElement xml.StartElement) error {

	switch startElement.Name.Local {

	case "tileset":

		tileset := &Tileset{}
		err := decoder.DecodeElement(tileset, &startElement)
		if err != nil {
			return err
		}
		c.Type = startElement.Name.Local
		c.Value = tileset

		// Update any image sources that are embedded the tmx file.
		if tileset.Image != nil {

			// Get image directory and filename and create a safe path.
			imgDir, imgFile := filepath.Split(tileset.Image.Source)
			imgPath := filepath.Join(tmxDir, imgDir, imgFile)

			// Update image source with a safe path.
			tileset.Image.Source = imgPath

		}

		// External Tileset
		// Update any tileset and image sources that come from external tsx files.
		if tileset.Source != "" {

			// Get tsx directory and filename and create a safe path.
			tsxDir, tsxFile := filepath.Split(tileset.Source)
			tsxPath := filepath.Join(tmxDir, tsxDir, tsxFile)

			// Update tileset source with a safe path.
			tileset.Source = tsxPath

			// fmt.Println(tsxPath)

			// Unmarshal a tsx path.
			tsxBytes, err := ioutil.ReadFile(tsxPath)
			if err != nil {
				return fmt.Errorf("error reading tsx file: %w", err)
			}
			err = xml.Unmarshal(tsxBytes, &tileset)
			if err != nil {
				return fmt.Errorf("error unmarshaling tsx bytes: %w", err)
			}

			// Get image directory and filename and create a safe path.
			imgDir, imgFile := filepath.Split(tileset.Image.Source)
			imgPath := filepath.Join(tmxDir, tsxDir, imgDir, imgFile)

			// Update image source with a safe path.
			tileset.Image.Source = imgPath

		}

	case "properties":

		properties := &Properties{}
		err := decoder.DecodeElement(properties, &startElement)
		if err != nil {
			return err
		}
		c.Type = startElement.Name.Local
		c.Value = properties

	case "layer":

		layer := &Layer{}
		err := decoder.DecodeElement(layer, &startElement)
		if err != nil {
			return err
		}
		c.Type = startElement.Name.Local
		c.Value = layer

	case "objectgroup":

		objectGroup := &ObjectGroup{}
		err := decoder.DecodeElement(objectGroup, &startElement)
		if err != nil {
			return err
		}
		c.Type = startElement.Name.Local
		c.Value = objectGroup

	case "imagelayer":

		imageLayer := &ImageLayer{}
		err := decoder.DecodeElement(imageLayer, &startElement)
		if err != nil {
			return err
		}
		c.Type = startElement.Name.Local
		c.Value = imageLayer

	case "group":

		group := &Group{}
		err := decoder.DecodeElement(group, &startElement)
		if err != nil {
			return err
		}
		c.Type = startElement.Name.Local
		c.Value = group

	default:
		return fmt.Errorf("unknown elements: %s", startElement)
	}

	return nil

}

func (c *Content) String() string {
	var b strings.Builder

	switch v := c.Value.(type) {

	case *Tileset:
		fmt.Fprintf(&b, v.String())
	case *Properties:
		fmt.Fprintf(&b, v.String())
	case *Layer:
		fmt.Fprintf(&b, v.String())
	case *ObjectGroup:
		fmt.Fprintf(&b, v.String())
	case *ImageLayer:
		fmt.Fprintf(&b, v.String())
	case *Group:
		fmt.Fprintf(&b, v.String())

	default:
		fmt.Fprintf(&b, "content not handled (%T)", v)
	}

	return b.String()
}
