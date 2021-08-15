package spriteplus

import (
	"fmt"

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

//MakeAnimation builds an animation from the given sprite sprites and frame lengths.
func MakeAnimation(sprites []*pixel.Sprite, frameLength int) (*Animation, error) {
	targetBounds := sprites[0].Picture().Bounds()
	for index, sp := range sprites {
		if sp.Picture().Bounds() != targetBounds {
			return nil, fmt.Errorf("tried to make Animation from non-consistent sprite size at index: %v", index)
		}
	}
	return &Animation{
		sprites:            sprites,
		updateCounter:      0,
		frameLength:        frameLength,
		currentSpriteIndex: 0,
	}, nil
}

//MakeAnimation builds an animation using sprites pulled from the given spritesheet using the given ids and frame lengths.
func MakeAnimationFromSheet(ss SpriteSheet, ids []string, frameLength int) (*Animation, error) {
	var sprites []*pixel.Sprite
	for _, id := range ids {
		sprites = append(sprites, ss.GetSprite(id))
	}

	return MakeAnimation(sprites, frameLength)
}
