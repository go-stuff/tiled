package tmx

// LayerTiles structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#tmx-tilelayer-tile
type LayerTiles struct {
	GID []int

	// Not to be confused with the tile element inside a tileset, this element defines the value of a single tile on a
	// tile layer. This is however the most inefficient way of storing the tile layer data, and should generally be
	// avoided.
}
