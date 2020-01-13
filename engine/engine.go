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

// Tiles can be flipped horizontally, vertically or both.
const (
	FlippedHorizontally = 0x80000000
	FlippedVertically   = 0x40000000
	FlippedDiagonally   = 0x20000000
	Flipped             = FlippedHorizontally | FlippedVertically | FlippedDiagonally
)

// AnimationTile has fields to keep track of animations.
type AnimationTile struct {
	// FrameIndex keeps track of the current animated frame.
	FrameIndex int
	// FrameDuration keeps track of how long the current frame has been drawn.
	FrameDuration int64
}

type LayerRectangle struct {
	X         float64
	Y         float64
	Image     *ebiten.Image
	Rectangle image.Rectangle
}

// Engine can be used when creating games in golang.
type Engine struct {
	Map *tmx.Map

	ScaleX float64
	ScaleY float64

	// Image is a custom field, it is a map of tileset images, accessed by image source path.
	Image map[string]*ebiten.Image

	// TilesetTile is a custom field, it is a map of tile, accessed by GID.
	TilesetTile map[*tmx.Tileset]map[int]*tmx.Tile

	// AnimationTile is map to keep track of animation frames.
	AnimationTile map[*tmx.Tile]*AnimationTile

	// LayerRectangle is a map to store rectangles for less compute.
	LayerRectangle map[*tmx.Layer]map[int]*LayerRectangle
}

// LoadEngine loads a tmx file as an engine to be used when creating games in golang.
func LoadEngine(source string) (*Engine, error) {
	e := new(Engine)

	tmxFile, err := tmx.LoadTMX(source)
	if err != nil {
		return nil, err
	}

	e.Map = tmxFile.Map

	e.ScaleX = 1.0
	e.ScaleY = 1.0

	// map[Tileset.Source]*image.Image
	e.Image = make(map[string]*ebiten.Image)

	// map[GID]*Tile
	e.TilesetTile = make(map[*tmx.Tileset]map[int]*tmx.Tile)

	// map[*Tile]*AnimationTile
	e.AnimationTile = make(map[*tmx.Tile]*AnimationTile)

	// map[*Layer][Data.GID]*image.Rectangle
	e.LayerRectangle = make(map[*tmx.Layer]map[int]*LayerRectangle)

	for _, tileset := range e.Map.Tileset {

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

	for _, layer := range e.Map.Layer {

		e.LayerRectangle[layer] = make(map[int]*LayerRectangle)

		x, y := 0, 0

		for _, gid := range layer.Data.Tile.GID {
			// Increment y co-ordinate.
			if x >= layer.Width {
				x = 0
				y++
			}

			// If tile GID on a Layer is 0, it is empty, skip it, still increment x co-ordinate.
			if gid == 0 {
				x++
				continue
			}

			// Get the image.Rectangle of the GID on its Tileset.
			rectangle, tileset, err := e.GIDRectangle(gid, e.Map.Tileset)
			if err != nil {
				return nil, err
			}

			e.LayerRectangle[layer][gid] = new(LayerRectangle)

			e.LayerRectangle[layer][gid].X = float64(x*tileset.TileWidth) * e.ScaleX
			e.LayerRectangle[layer][gid].Y = float64(y*tileset.TileHeight) * e.ScaleY

			e.LayerRectangle[layer][gid].Rectangle = rectangle

			e.LayerRectangle[layer][gid].Image = e.Image[tileset.Image.Source]

			// Increment x co-ordinate.
			x++

			// // Skip tiles not visible by the camera.
			// w, h := screen.Size()
			// if float64(x*tileset.TileWidth)*game.ScaleX > float64(w) || float64(y*tileset.TileHeight)*game.ScaleY > float64(h) {
			// 	x++
			// 	continue
			// }

			// offsetX, offsetY := 100.0, 100.0

			// Translage tile on the screen based on the x and y co-ordinates.
			// opts := &ebiten.DrawImageOptions{}
			// opts.GeoM.Scale(game.ScaleX, game.ScaleY)
			// opts.GeoM.Translate(float64(x*tileset.TileWidth)*game.ScaleX, float64(y*tileset.TileHeight)*game.ScaleY)

			// Draw the tile to the screen.
			// screen.DrawImage(
			// 	l.Engine.Image[tileset.Image.Source].SubImage(rectangle).(*ebiten.Image),
			// 	//game.Image[tileset.Image.Source].SubImage(rectangle).(*ebiten.Image),
			// 	opts,
			// )
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

// GIDStripFlipping just returns the tile gid, it does not care about horizontal or vertical flipping.
func (e *Engine) GIDStripFlipping(gid int) int {
	return gid &^ Flipped
}

// GIDFlipped returns status of horizontal, vertical and diagonal flipping also the stripped gid.
func (e *Engine) GIDFlipped(gid int) (bool, bool, bool, int) {
	var h, v, d bool
	if gid&FlippedHorizontally != 0 {
		h = true
	}
	if gid&FlippedVertically != 0 {
		v = true
	}
	if gid&FlippedDiagonally != 0 {
		d = true
	}
	return h, v, d, gid &^ Flipped
}

func (e *Engine) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Custom:\n")

	if e.Image != nil {
		fmt.Fprintf(&b, "\tImage: (%T) %v\n", e.Image, e.Image)
	}

	return b.String()
}
