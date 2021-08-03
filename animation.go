package spriteplus

import (
	"github.com/faiface/pixel"
)

type Animation struct {
	sprites            []*pixel.Sprite
	currentSpriteIndex int

	//frameLength is the number of times Update() is called before the animation moves to the next sprite
	frameLength int
	//updateCounter is how many times Update() has been called since the last sprite change
	updateCounter int
}

//Next moves the animation to the next frame.
func (ba *Animation) Next() {
	if ba.currentSpriteIndex++; ba.currentSpriteIndex > len(ba.sprites) {
		ba.currentSpriteIndex = 0
	}
}

func (ba *Animation) Draw(target pixel.Target, mat pixel.Matrix) {
	ba.sprites[ba.currentSpriteIndex].Draw(target, mat)

	if ba.updateCounter++; ba.updateCounter >= ba.frameLength {
		ba.updateCounter = 0
		ba.Next()
	}
}

func MakeAnimation(sprites []*pixel.Sprite, frameLength int) *Animation {
	targetBounds := sprites[0].Picture().Bounds()
	for _, sp := range sprites {
		if sp.Picture().Bounds() != targetBounds {
			panic("Tried to make Animation from non-consistent sprite sizes.") //TODO return errror instead of panic
		}
	}
	return &Animation{
		sprites:            sprites,
		updateCounter:      0,
		frameLength:        frameLength,
		currentSpriteIndex: 0,
	}
}

func MakeAnimationFromSheet(ss SpriteSheet, ids []string, frameLength int) *Animation {
	var sprites []*pixel.Sprite
	for _, id := range ids {
		sprites = append(sprites, ss.GetSprite(id))
	}

	return MakeAnimation(sprites, frameLength)
}
