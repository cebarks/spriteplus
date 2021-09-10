package spriteplus

import (
	"fmt"

	"github.com/faiface/pixel"
)

//Animation is a series of sprites that change after a certain number of frames in sequential order
type Animation struct {
	sprites            []*pixel.Sprite
	currentSpriteIndex int

	//frameLength is the number of times Update() is called before the animation moves to the next sprite
	frameLength int
	//drawCount is how many times Draw() has been called since the last sprite change
	drawCount int
}

//Next force updates the animation to the next frame.
func (ba *Animation) Next() {
	if ba.currentSpriteIndex++; ba.currentSpriteIndex >= len(ba.sprites) {
		ba.currentSpriteIndex = 0
	}
}

//Draw draws the current sprite to the given target and updates
func (ba *Animation) Draw(target pixel.Target, mat pixel.Matrix) {
	ba.sprites[ba.currentSpriteIndex].Draw(target, mat)

	if ba.drawCount++; ba.drawCount >= ba.frameLength {
		ba.drawCount = 0
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
		drawCount:          0,
		frameLength:        frameLength,
		currentSpriteIndex: 0,
	}, nil
}

//MakeAnimationFromSheet builds an animation using sprites pulled from the given spritesheet using the given ids and frame lengths.
func MakeAnimationFromSheet(ss SpriteSheet, ids []string, frameLength int) (*Animation, error) {
	var sprites []*pixel.Sprite
	for _, id := range ids {
		sprites = append(sprites, ss.GetSprite(id))
	}

	return MakeAnimation(sprites, frameLength)
}
