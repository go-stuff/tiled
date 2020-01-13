package tmx

import (
	"fmt"
	"strings"
)

// Tiles can be flipped horizontally, vertically or both.
const (
	FlippedHorizontally = 0x80000000
	FlippedVertically   = 0x40000000
	FlippedDiagonally   = 0x20000000
	Flipped             = FlippedHorizontally | FlippedVertically | FlippedDiagonally
)

// Flipping structure: https://doc.mapeditor.org/en/stable/reference/tmx-map-format/#tile-flipping
type Flipping struct {
	// The highest three bits of the gid store the flipped states. Bit 32 is used for storing whether the tile is
	// horizontally flipped, bit 31 is used for the vertically flipped tiles and bit 30 indicates whether the tile is
	// flipped (anti) diagonally, enabling tile rotation. These bits have to be read and cleared before you can find
	// out which tileset a tile belongs to.

	// When rendering a tile, the order of operation matters. The diagonal flip (x/y axis swap) is done first, followed
	// by the horizontal and vertical flips.

	// The following C++ pseudo-code should make it all clear:

	// // Bits on the far end of the 32-bit global tile ID are used for tile flags
	// const unsigned FLIPPED_HORIZONTALLY_FLAG = 0x80000000;
	// const unsigned FLIPPED_VERTICALLY_FLAG   = 0x40000000;
	// const unsigned FLIPPED_DIAGONALLY_FLAG   = 0x20000000;

	// ...

	// // Extract the contents of the <data> element
	// string tile_data = ...

	// unsigned char *data = decompress(base64_decode(tile_data));
	// unsigned tile_index = 0;

	// // Here you should check that the data has the right size
	// // (map_width * map_height * 4)

	// for (int y = 0; y < map_height; ++y) {
	//   for (int x = 0; x < map_width; ++x) {
	//     unsigned global_tile_id = data[tile_index] |
	//                               data[tile_index + 1] << 8 |
	//                               data[tile_index + 2] << 16 |
	//                               data[tile_index + 3] << 24;
	//     tile_index += 4;

	//     // Read out the flags
	//     bool flipped_horizontally = (global_tile_id & FLIPPED_HORIZONTALLY_FLAG);
	//     bool flipped_vertically = (global_tile_id & FLIPPED_VERTICALLY_FLAG);
	//     bool flipped_diagonally = (global_tile_id & FLIPPED_DIAGONALLY_FLAG);

	//     // Clear the flags
	//     global_tile_id &= ~(FLIPPED_HORIZONTALLY_FLAG |
	//                         FLIPPED_VERTICALLY_FLAG |
	//                         FLIPPED_DIAGONALLY_FLAG);

	//     // Resolve the tile
	//     for (int i = tileset_count - 1; i >= 0; --i) {
	//       Tileset *tileset = tilesets[i];

	//       if (tileset->first_gid() <= global_tile_id) {
	//         tiles[y][x] = tileset->tileAt(global_tile_id - tileset->first_gid());
	//         break;
	//       }
	//     }
	//   }
	// }

	// (Since the above code was put together on this wiki page and can’t be directly tested, please make sure to
	// report any errors you encounter when basing your parsing code on it, thanks.)

	Horizontal bool
	Vertical   bool
}

// StripFlipping just returns the tile gid, it does not care about horizontal or vertical flipping.
func StripFlipping(id int) int {
	return id &^ Flipped
}

func (f *Flipping) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Flipping:\n")
	fmt.Fprintf(&b, "\tHorizontal: (%T) %t\n", f.Horizontal, f.Horizontal)
	fmt.Fprintf(&b, "\tVertical:   (%T) %t\n", f.Vertical, f.Vertical)

	return b.String()
}
