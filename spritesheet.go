package spriteplus

import (
	"image"

	"github.com/dusk125/pixelutils"
	"github.com/dusk125/pixelutils/packer"
	"github.com/faiface/pixel"
)

//NewSpriteSheet creates a new instatiated sprite sheet
func NewSpriteSheet(debugDraw bool) *SpriteSheet {

	return &SpriteSheet{
		Packr: packer.New(),
		Alias: make(map[string]int),
		idGen: &pixelutils.IDGen{},
		Cache: make(map[string]*pixel.Sprite),
	}
}

//SpriteSheet
type SpriteSheet struct {
	Cache map[string]*pixel.Sprite
	Alias map[string]int
	idGen *pixelutils.IDGen
	Packr *packer.Packer
}

func (ss *SpriteSheet) AddSprite(pic *image.RGBA, id string) error {
	ss.idGen.Lock()
	intId := ss.idGen.Gen()
	ss.idGen.Unlock()

	ss.Alias[id] = intId

	ss.Packr.Insert(intId, pic)

	if err := ss.Packr.Pack(); err != nil {
		return err
	}

	ss.Packr.Get(intId)

	ss.Cache[id] = pixel.NewSprite(ss.Packr.Picture(), )

	return nil
}

//GetSprite will return the sprite in the Cache (or create&add it to the Cache) from the given int id
func (ss *SpriteSheet) GetSprite(id string) *pixel.Sprite {
	sprite := ss.Cache[id]

	if sprite == nil {
		sprite = ss.Packr.Get()
		sprite = ss.Packr.SpriteFrom(ss.Alias[id])
		ss.Cache[id] = sprite
	}

	return sprite
}

//SourcePic returns the underlying pixel.Picture of the spritesheet (for sure with Batch rendering)
func (ss SpriteSheet) SourcePic() pixel.Picture {
	return ss.Packr.Picture()
}
