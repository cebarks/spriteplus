package spriteplus

import (
	"github.com/dusk125/pixelutils/packer"
	"github.com/faiface/pixel"
)

//BuildRuntimeSpriteSheet will build a spritesheet from the supplied []*pixel.PictureData. The ids will be assigned in order
func BuildRuntimeSpriteSheet(pics ...*pixel.PictureData) (SpriteSheet, error) {
	packr := packer.NewPacker(0, 0, packer.AllowGrowth)

	for _, pic := range pics {
		packr.InsertPictureData(packr.GenerateId(), pic)
	}

	packr.Optimize()

	sheet := RuntimeSpriteSheet{
		packr:  packr,
		length: len(pics),
		Cache:  make(map[interface{}]*pixel.Sprite),
	}

	return &sheet, nil
}

type RuntimeSpriteSheet struct {
	Cache  map[interface{}]*pixel.Sprite
	packr  *packer.Packer
	length int
}

func (rss *RuntimeSpriteSheet) SourcePic() pixel.Picture {
	return rss.packr.Picture()
}

//GetSprite will return the sprite in the Cache (or create&add it to the Cache) from the given int id
func (rss *RuntimeSpriteSheet) GetSprite(id interface{}) *pixel.Sprite {
	sprite := rss.Cache[id]

	if sprite == nil {
		sprite = rss.packr.SpriteFrom(id.(int))
		rss.Cache[id] = sprite
	}

	return sprite
}
