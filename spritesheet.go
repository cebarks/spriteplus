package spriteplus

import (
	"github.com/dusk125/pixelutils/packer"
	"github.com/faiface/pixel"
)

func (ss SpriteSheet) SourcePic() pixel.Picture {
	return ss.packr.Picture()
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

	if err != nil {
		return err
	}

	return nil
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

func NewSpriteSheet(allowGrowth bool) *SpriteSheet {
	var flags uint8 = 0

	if allowGrowth {
		flags |= packer.AllowGrowth
	}

	return &SpriteSheet{
		packr: packer.NewPacker(256, 256, flags),
	}
}
