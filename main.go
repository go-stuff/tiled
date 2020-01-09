package main

import (
	"fmt"
	"log"

	"github.com/go-stuff/tiled/tmx"
)

// Test loading and printing a tmx file.
func main() {
	t, err := tmx.LoadTMX("/home/stevo/code/github.com/go-stuff/game/asset/rpg-overworld-tileset v1.2 (wonderdot)/Extras/Scenes.tmx")
	//t, err := tmx.LoadTMX("./testdata/map.tmx")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t.String())

	for _, tileset := range t.Map.Tileset {
		fmt.Println(tileset.Image.String())
		fmt.Println(tileset.FirstGID)
	}

	// for k, v := range t.Map.Tile {
	// 	fmt.Println("key: ", k, ", value: ", v)
	// }

	// gid, err := t.Map.Layer[0].GID(5, 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(gid)

	// tileset, err := t.Map.Layer[0].GIDTileset(gid, t.Map.Tileset)
	// tileset, err := t.Map.Layer[0].GIDTileset(1537, t.Map.Tileset)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(tileset)
}
