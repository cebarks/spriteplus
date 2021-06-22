package spriteplus

import (
	"image"

	"github.com/cebarks/spriteplus/tileset"
	"github.com/dusk125/pixelutils/packer"
	"github.com/faiface/pixel"
)

func NewSpriteSheet(allowGrowth bool, debugDraw bool) *SpriteSheet {
	var flags uint8 = 0

	if allowGrowth {
		flags |= packer.AllowGrowth
	}

	if debugDraw {
		flags |= packer.DebugDraw
	}

	return &SpriteSheet{
		Packr: packer.NewPacker(256, 256, flags),
		Alias: make(map[string]int),
		Cache: make(map[string]*pixel.Sprite),
	}
}

type SpriteSheet struct {
	Cache map[string]*pixel.Sprite
	Alias map[string]int
	Packr *packer.Packer
}

func (ss *SpriteSheet) AddSprite(pic pixel.Picture, id string) error {
	intId := ss.Packr.GenerateId()

	ss.Alias[id] = intId

	err := ss.Packr.InsertPictureDataV(intId, pic.(*pixel.PictureData), packer.OptimizeOnInsert)

	ss.Cache[id] = ss.Packr.SpriteFrom(ss.Alias[id])

	if err != nil {
		return err
	}

	return nil
}

//AddTilesetRaw parses the tiles from a tileset and adds them to this spritesheet
func (ss *SpriteSheet) AddTilesetRaw(img image.Image, ids []string, tileHeight, tileWidth, rows, columns int) (int, error) {
	var tilesFound, index int

	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			minX := x * tileHeight
			minY := y * tileWidth
			maxX := (x + 1) * tileHeight
			maxY := (y + 1) * tileWidth

			sub, err := Subimage(img, minX, minY, maxX, maxY)
			if err != nil {
				return -1, err
			}

			pic := pixel.PictureDataFromImage(sub)

			ss.AddSprite(pic, ids[index])

			index++
		}
	}

	return tilesFound, nil
}

func (ss *SpriteSheet) AddTilesetFromPath(path string) (int, error) {
	ts, err := tileset.TilesetFromPath(path)
	if err != nil {
		return -1, err
	}
	return ss.AddTilesetRaw(ts.Image, ts.Ids(), ts.TileHeight, ts.TileWidth, ts.Rows, ts.Columns)
}

//GetSprite will return the sprite in the Cache (or create&add it to the Cache) from the given int id
func (ss *SpriteSheet) GetSprite(id string) *pixel.Sprite {
	sprite := ss.Cache[id]

	if sprite == nil {
		sprite = ss.Packr.SpriteFrom(ss.Alias[id])
		ss.Cache[id] = sprite
	}

	return sprite
}

//SourcePic returns the underlying pixel.Picture of the spritesheet
func (ss SpriteSheet) SourcePic() pixel.Picture {
	return ss.Packr.Picture()
}
