package spriteplus

import (
	"github.com/dusk125/pixelutils"
	"github.com/dusk125/pixelutils/packer"
	"github.com/faiface/pixel"
)

//SpriteSheet is a id-based sprite manager interface
type SpriteSheet interface {
	//SourcePic returns the *pixel.Picture source of all sprites on this spritesheet
	SourcePic() pixel.Picture
	//GetSprite returns the sprite for the given id
	GetSprite(id interface{}) *pixel.Sprite
}

type BasicSpriteSheet struct {
	//size of the sprites on the sprite sheet
	sizeX int
	sizeY int

	//how sprites per x/y
	countX int
	countY int

	//sourcePic is the pixel.Picture backing this spritesheet
	sourcePic pixel.Picture
}

func MakeBasicSheet(countX, countY, sizeX, sizeY int, path string) (SpriteSheet, error) {
	pic, err := pixelutils.LoadPictureData(path)
	if err != nil {
		return nil, err
	}

	sheet := BasicSpriteSheet{
		sourcePic: pic,
		sizeX:     sizeX,
		sizeY:     sizeY,
		countX:    countX,
		countY:    countY,
	}

	return sheet, nil
}

func MakeBasicSheetFromPicture(countX, countY, sizeX, sizeY int, pic pixel.Picture) (SpriteSheet, error) {
	sheet := BasicSpriteSheet{
		sourcePic: pic,
		sizeX:     sizeX,
		sizeY:     sizeY,
		countX:    countX,
		countY:    countY,
	}

	return sheet, nil
}

func (bss BasicSpriteSheet) SourcePic() pixel.Picture {
	return bss.sourcePic
}

//GetSprite will return the sprite from the given int id
func (bss BasicSpriteSheet) GetSprite(id interface{}) *pixel.Sprite {
	return pixel.NewSprite(bss.sourcePic, bss.rectForId(id.(int)))
}

func (bss BasicSpriteSheet) rectForId(idint interface{}) pixel.Rect {
	id := idint.(int)

	x := (id % bss.countX)
	y := int(float64(id) / float64(bss.countY))

	minX := float64(x) * float64(bss.sizeX)
	minY := float64(y) * float64(bss.sizeY)
	maxX := (float64(x) + 1) * float64(bss.sizeX)
	maxY := (float64(y) + 1) * float64(bss.sizeY)

	return pixel.R(minX, minY, maxX, maxY)
}

type CachedSpriteSheet struct {
	BasicSpriteSheet
	Cache map[interface{}]*pixel.Sprite
}

//GetSprite will return the sprite in the Cache (or create&add it to the Cache) from the given int id
func (css *CachedSpriteSheet) GetSprite(id interface{}) *pixel.Sprite {
	sprite := css.Cache[id]

	if sprite == nil {
		sprite = pixel.NewSprite(css.sourcePic, css.rectForId(id))
		css.Cache[id] = sprite
	}

	return sprite
}

func MakeCachedSheet(countX, countY, sizeX, sizeY int, path string) (SpriteSheet, error) {
	pic, err := pixelutils.LoadPictureData(path)
	if err != nil {
		return nil, err
	}

	return MakeCachedSheetFromPicture(countX, countY, sizeX, sizeY, pic)
}

func MakeCachedSheetFromPicture(countX, countY, sizeX, sizeY int, pic pixel.Picture) (SpriteSheet, error) {
	sheet := CachedSpriteSheet{
		BasicSpriteSheet: BasicSpriteSheet{
			sourcePic: pic,
			sizeX:     sizeX,
			sizeY:     sizeY,
			countX:    countX,
			countY:    countY,
		},
		Cache: make(map[interface{}]*pixel.Sprite),
	}

	return &sheet, nil
}

type RuntimeSpriteSheet struct {
	packr *packer.Packer
}

func (rss *RuntimeSpriteSheet) SourcePic() pixel.Picture {
	return rss.packr.Picture()
}

//GetSprite will return the sprite in the Cache (or create&add it to the Cache) from the given int id
func (rss *RuntimeSpriteSheet) GetSprite(id interface{}) *pixel.Sprite {
	rss.packr.SpriteFrom(id.(int))
	return nil
}

//BuildRuntimeSpriteSheet will build a spritesheet from the supplied []*pixel.PictureData. The ids will be assigned in order
func BuildRuntimeSpriteSheet(pics []*pixel.PictureData) (SpriteSheet, error) {
	packr := packer.NewPacker(0, 0, packer.OptimizeOnInsert|packer.AllowGrowth)

	for _, pic := range pics { //TODO make optionally concurrent
		err := packr.InsertPictureData(packr.GenerateId(), pic) //TODO waiting on https://github.com/dusk125/pixelutils/pull/1
		if err != nil {
			return nil, err
		}
	}

	sheet := &RuntimeSpriteSheet{
		packr: packr,
	}

	return sheet, nil
}
