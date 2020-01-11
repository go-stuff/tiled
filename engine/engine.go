package engine

import (
	"fmt"
	"image"
	_ "image/jpeg" // register JPEG decoder
	_ "image/png"  // register PNG decoder
	"strings"

	"github.com/go-stuff/tiled/tmx"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// AnimationTile has fields to keep track of animations.
type AnimationTile struct {
	// FrameIndex keeps track of the current animated frame.
	FrameIndex int `xml:"-"`
	// FrameDuration keeps track of how long the current frame has been drawn.
	FrameDuration int64 `xml:"-"`
}

// Engine can be used when creating games in golang.
type Engine struct {
	Map *tmx.Map

	// Image is a custom field, it is a map of tileset images, accessed by image source path.
	Image map[string]*ebiten.Image `xml:"-"`

	// TilesetTile is a custom field, it is a map of tile, accessed by GID.
	TilesetTile map[*tmx.Tileset]map[int]*tmx.Tile `xml:"-"`

	// AnimationTile is map to keep track of animation frames.
	AnimationTile map[*tmx.Tile]*AnimationTile `xml:"-"`
}

// LoadEngine loads a tmx file as an engine to be used when creating games in golang.
func LoadEngine(source string) (*Engine, error) {
	e := new(Engine)

	tmxFile, err := tmx.LoadTMX(source)
	if err != nil {
		return nil, err
	}

	e.Map = tmxFile.Map

	// map[Tileset.Source]*image.Image
	e.Image = make(map[string]*ebiten.Image)

	// map[GID]*Tile
	e.TilesetTile = make(map[*tmx.Tileset]map[int]*tmx.Tile)

	// map[*Tile]*AnimationTile
	e.AnimationTile = make(map[*tmx.Tile]*AnimationTile)

	for _, tileset := range e.Map.Tileset {

		// imgBytes, err := ioutil.ReadFile(tileset.Image.Source)
		// if err != nil {
		// 	return nil, fmt.Errorf("error reading image file: %w", err)
		// }
		// pngImage, _, err := image.Decode(bytes.NewReader(imgBytes))
		// if err != nil {
		// 	return nil, fmt.Errorf("error decoding image file: %w", err)
		// }

		// // Add Image to custom Image Map
		// e.Image[tileset.Image.Source] = &pngImage

		// Load images as *ebiten.Image.
		img, _, err := ebitenutil.NewImageFromFile(tileset.Image.Source, ebiten.FilterDefault)
		if err != nil {
			return nil, err
		}
		e.Image[tileset.Image.Source] = img

		// Add Tiles to custom Tile Map
		for _, tile := range tileset.Tile {
			if e.TilesetTile[tileset] == nil {
				e.TilesetTile[tileset] = make(map[int]*tmx.Tile)
			}

			e.TilesetTile[tileset][tile.ID] = tile

			if tile.Animation != nil {
				e.AnimationTile[tile] = new(AnimationTile)
			}
		}
	}

	return e, nil
}

// UpdateAnimationTile updates the animation frame on each animated tile.
func (e *Engine) UpdateAnimationTile(milliseconds int64) {
	for tile, animaionTile := range e.AnimationTile {
		// Update the Milliseconds of this Particular Tile Frame
		animaionTile.FrameDuration += milliseconds

		// If the elapsed time is greater than the current frame duration, go to the next frame index and
		// reset time.
		if animaionTile.FrameDuration >
			int64(tile.Animation.Frame[animaionTile.FrameIndex].Duration) {
			animaionTile.FrameIndex++
			animaionTile.FrameDuration = 0
		}

		// If the current frame equals the last frame start over.
		if animaionTile.FrameIndex == len(tile.Animation.Frame) {
			animaionTile.FrameIndex = 0
			animaionTile.FrameDuration = 0
		}
	}
}

// GIDTileset returns the tileset a GID resides on.
func (e *Engine) GIDTileset(gid int, tileset []*tmx.Tileset) (*tmx.Tileset, error) {
	for _, tileset := range tileset {
		if gid >= tileset.FirstGID && gid < tileset.FirstGID+tileset.TileCount {
			return tileset, nil
		}
	}
	return nil, fmt.Errorf("tileset not found")
}

// GIDRectangle returns an image.Rectangle and Tileset of a GID.
func (e *Engine) GIDRectangle(gid int, tilesets []*tmx.Tileset) (image.Rectangle, *tmx.Tileset, error) {

	// Get the Tileset of the current GID.
	tileset, err := e.GIDTileset(gid, tilesets)
	if err != nil {
		return image.Rectangle{}, nil, err
	}

	// Get the real GID by subtracting the tileset firtst GID.
	gid -= tileset.FirstGID

	if e.TilesetTile[tileset][gid] != nil {
		tile := e.TilesetTile[tileset][gid]
		if tile.Animation != nil {
			gid = tile.Animation.Frame[e.AnimationTile[tile].FrameIndex].TileID
		}
	}

	// From the GID and Tileset calculate the row and column.
	row := int(gid / tileset.Columns)
	column := (gid % tileset.Columns)

	rectangle := image.Rectangle{
		Min: image.Point{column * tileset.TileWidth, row * tileset.TileHeight},
		Max: image.Point{(column * tileset.TileWidth) + tileset.TileWidth, (row * tileset.TileHeight) + tileset.TileHeight},
	}

	return rectangle, tileset, nil
}

func (e *Engine) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Custom:\n")

	if e.Image != nil {
		fmt.Fprintf(&b, "\tImage: (%T) %v\n", e.Image, e.Image)
	}

	return b.String()
}
