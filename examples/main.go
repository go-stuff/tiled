package main

import (
	"log"

	"github.com/go-stuff/tiled/tmx"
)

// Test loading and printing a tmx file.
func main() {
	_, err := tmx.LoadTMX("./testdata/map.tmx")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(t.String())

	// fmt.Println(t.TilesetCount(t.Map.Content))
	// fmt.Println(t.ImageCount(t.Map.Content))
	// fmt.Println(t.LayerCount(0, t.Map.Content))

	// file, err := os.Open("./testdata/map.tmx")
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("error opening tmx file: %w", err))
	// }
	// fileInfo, err := file.Stat()
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("error getting info on tmx file: %w", err))
	// }
	// tmxBytes := make([]byte, fileInfo.Size())
	// _, err = file.Read(tmxBytes)
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("error reading tmx file: %w", err))
	// }

	// t, err = tmx.LoadTMXBytes(tmxBytes)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(t.String())

}
