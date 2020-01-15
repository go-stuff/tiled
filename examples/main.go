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
