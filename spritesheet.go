package spriteplus

import (
	"github.com/dusk125/pixelutils"
	"github.com/faiface/pixel"
)

type SpriteSheet interface {
	//SourcePic returns the *pixel.Picture source of all sprites on this spritesheet
	SourcePic() pixel.Picture
	//GetSprite returns the sprite for the given id
	GetSprite(id interface{}) *pixel.Sprite
	//
}

func NewBasicSheet(countX, countY, sizeX, sizeY int, path string) (SpriteSheet, error) {
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

func NewBasicSheetFromPicture(countX, countY, sizeX, sizeY int, pic pixel.Picture) (SpriteSheet, error) {
	sheet := BasicSpriteSheet{
		sourcePic: pic,
		sizeX:     sizeX,
		sizeY:     sizeY,
		countX:    countX,
		countY:    countY,
	}

	return sheet, nil
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

func (bss BasicSpriteSheet) SourcePic() pixel.Picture {
	return bss.sourcePic
}

func (bss BasicSpriteSheet) GetSprite(id interface{}) *pixel.Sprite {
	return pixel.NewSprite(bss.sourcePic, bss.rectForId(id))
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
