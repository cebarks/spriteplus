package spriteplus

import (
	"github.com/dusk125/pixelutils/packer"
	"github.com/faiface/pixel"
)

//NewSpriteSheet creates a new instatiated sprite sheet
func NewSpriteSheet(debugDraw bool) *SpriteSheet {
	var flags uint8 = 0

	flags |= packer.AllowGrowth

	if debugDraw {
		flags |= packer.DebugDraw
	}

	return &SpriteSheet{
		Packr: packer.NewPacker(256, 256, flags),
		Alias: make(map[string]int),
		Cache: make(map[string]*pixel.Sprite),
	}
}

//SpriteSheet
type SpriteSheet struct {
	Cache map[string]*pixel.Sprite
	Alias map[string]int
	Packr *packer.Packer
}

func (ss *SpriteSheet) AddSprite(pic pixel.Picture, id string) error {
	intId := ss.Packr.GenerateId()

	ss.Alias[id] = intId

	err := ss.Packr.InsertPictureDataV(intId, pic.(*pixel.PictureData), 0)
	if err != nil {
		return err
	}

	ss.Cache[id] = ss.Packr.SpriteFrom(ss.Alias[id])

	return nil
}

//Optimize the underlying texture
func (ss *SpriteSheet) Optimize() {
	ss.Packr.Optimize()
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

//SourcePic returns the underlying pixel.Picture of the spritesheet (for sure with Batch rendering)
func (ss SpriteSheet) SourcePic() pixel.Picture {
	return ss.Packr.Picture()
}
