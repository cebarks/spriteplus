package spriteplus

import (
	"github.com/faiface/pixel"
)

//Animation represents a multi-frame sprite
type Animation interface {
	Update()
	Draw(target pixel.Target, mat pixel.Matrix)
}

type BasicAnimation struct {
	sprites            []*pixel.Sprite
	currentSpriteIndex int

	//frameLength is the number of times Update() is called before the animation moves to the next sprite
	frameLength int
	//updateCounter is how many times Update() has been called since the last sprite change
	updateCounter int
}

func (ba *BasicAnimation) Update() {
	ba.updateCounter++
	if ba.updateCounter >= ba.frameLength {
		ba.updateCounter = 0

		if ba.currentSpriteIndex++; ba.currentSpriteIndex > len(ba.sprites) {
			ba.currentSpriteIndex = 0
		}
	}
}

func (ba *BasicAnimation) Draw(target pixel.Target, mat pixel.Matrix) {
	ba.sprites[ba.currentSpriteIndex].Draw(target, mat)
}

func MakeBasicAnimation(sprites []*pixel.Sprite, frameLength int) Animation {
	targetBounds := sprites[0].Picture().Bounds()
	for _, sp := range sprites {
		if sp.Picture().Bounds() != targetBounds {
			panic("Tried to make BasicAnimation from non-consistent sprite sizes.") //TODO return errror instead of panic
		}
	}
	return &BasicAnimation{
		sprites:            sprites,
		updateCounter:      0,
		frameLength:        frameLength,
		currentSpriteIndex: 0,
	}
}

func MakeBasicAnimationFromSheet(ss SpriteSheet, ids []string, frameLength int) Animation {
	var sprites []*pixel.Sprite
	for _, id := range ids {
		sprites = append(sprites, ss.GetSprite(id))
	}

	return MakeBasicAnimation(sprites, frameLength)
}
