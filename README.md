# tiled

Loading [TMX Map Format](https://doc.mapeditor.org/de/stable/reference/tmx-map-format/#tmx-map-format) files created by the [Tiled](https://www.mapeditor.org/) map editor into a [Go](https://golang.org/) struct. This package does not do anything fancy, it does not do any decoding, it unmarshalls data from `.tmx` and `.tsx` files and populates a `tmx.Map` struct. It updates tileset and image sources with better path information.

The [TMX Map Format](https://doc.mapeditor.org/de/stable/reference/tmx-map-format/#tmx-map-format) as close as possible. The only field used that is not listed in the spec is [tmx.Data.InnerXML](https://github.com/go-stuff/tiled/blob/master/tmx/data.go), it is the raw XML nested inside the tag `<data>`.


## usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/go-stuff/tiled/tmx"
)

// Test loading and printing a tmx file.
func main() {
	t, err := tmx.LoadTMX("./testdata/map.tmx")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t.String())
}
```