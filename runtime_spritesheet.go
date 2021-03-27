package spriteplus

import (
	"github.com/dusk125/pixelutils/packer"
	"github.com/faiface/pixel"
)

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
	packr := packer.NewPacker(0, 0, packer.AllowGrowth)

	for _, pic := range pics { //TODO make optionally concurrent
		err := packr.InsertPictureData(packr.GenerateId(), pic) //TODO waiting on https://github.com/dusk125/pixelutils/pull/1
		if err != nil {
			return nil, err
		}
	}

	packr.Optimize()

	sheet := RuntimeSpriteSheet{
		packr: packr,
	}

	return &sheet, nil
}
