package spriteplus

import (
	"fmt"
	"image"

	"github.com/dusk125/pixelutils/packer"
	"github.com/faiface/pixel"
)

func NewSpriteSheet(allowGrowth bool) *SpriteSheet {
	var flags uint8 = 0

	if allowGrowth {
		flags |= packer.AllowGrowth
	}

	return &SpriteSheet{
		packr: packer.NewPacker(256, 256, flags),
	}
}

type SpriteSheet struct {
	Cache map[interface{}]*pixel.Sprite
	Alias map[string]int
	packr *packer.Packer
}

func (ss *SpriteSheet) AddSprite(pic *pixel.PictureData, id string) error {
	intId := ss.packr.GenerateId()

	ss.Alias[id] = intId

	err := ss.packr.InsertPictureDataV(intId, pic, packer.OptimizeOnInsert)

	ss.Cache[id] = ss.packr.SpriteFrom(ss.Alias[id])

	if err != nil {
		return err
	}

	return nil
}

//AddTileset parses the tiles from a tileset and adds them to this spritesheet
func (ss *SpriteSheet) AddTileset(img image.Image, ids []string, tileHeight, tileWidth, rows, columns int) (int, error) {
	var tilesFound, index int

	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			minX := (x) * tileHeight
			minY := (y) * tileWidth
			maxX := (x + 1) * tileHeight
			maxY := (y + 1) * tileWidth

			sub, err := subimage(img, minX, minY, maxX, maxY)
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

//subimage returns a subimage of the given image.Image
func subimage(img image.Image, minX, minY, maxX, maxY int) (image.Image, error) {
	return subimageRect(img, image.Rect(minX, minY, maxX, maxY))
}

//subimageRect returns a subimage of the given image.Image
func subimageRect(img image.Image, rect image.Rectangle) (image.Image, error) {
	var sub image.Image

	switch img.(type) {
	case *image.RGBA:
		i := img.(*image.RGBA)
		sub = i.SubImage(rect)
	case *image.RGBA64:
		i := img.(*image.RGBA64)
		sub = i.SubImage(rect)
	case *image.NRGBA:
		i := img.(*image.NRGBA)
		sub = i.SubImage(rect)
	case *image.NRGBA64:
		i := img.(*image.NRGBA64)
		sub = i.SubImage(rect)
	case *image.Alpha:
		i := img.(*image.Alpha)
		sub = i.SubImage(rect)
	case *image.Alpha16:
		i := img.(*image.Alpha16)
		sub = i.SubImage(rect)
	case *image.Gray:
		i := img.(*image.Gray)
		sub = i.SubImage(rect)
	case *image.Gray16:
		i := img.(*image.Gray16)
		sub = i.SubImage(rect)
	case *image.CMYK:
		i := img.(*image.CMYK)
		sub = i.SubImage(rect)
	case *image.Paletted:
		i := img.(*image.Paletted)
		sub = i.SubImage(rect)
	default:
		return nil, fmt.Errorf("Couldn't cast type of img. Unknown type: %T", img)
	}

	return sub, nil
}

//GetSprite will return the sprite in the Cache (or create&add it to the Cache) from the given int id
func (ss *SpriteSheet) GetSprite(id string) *pixel.Sprite {
	sprite := ss.Cache[id]

	if sprite == nil {
		sprite = ss.packr.SpriteFrom(ss.Alias[id])
		ss.Cache[id] = sprite
	}

	return sprite
}

//GetSourcePic returns the underlying pixel.Picture of the spritesheet
func (ss SpriteSheet) GetSourcePic() pixel.Picture {
	return ss.packr.Picture()
}
