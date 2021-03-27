package spriteplus

import (
	"fmt"
	"testing"

	"github.com/faiface/pixel"
	"github.com/stretchr/testify/assert"
)

func Test_MakeBasicAnimation(t *testing.T) {
	ss, _ := MakeBasicSheet(2, 2, 4, 4, "testdata/4px-2x2-small.png")

	var sprites []*pixel.Sprite
	for i := 0; i < 4; i++ {
		sprites = append(sprites, ss.GetSprite(i))
	}

	a := MakeBasicAnimation(sprites, 4)
	ani := a.(*BasicAnimation)

	assert.Equal(t, 4, ani.frameLength)

	for i := 0; i < 4; i++ {
		assert.Equal(t, i, ani.currentSpriteIndex)
		assert.Equal(t, fmt.Sprintf("%p", sprites[i]), fmt.Sprintf("%p", ani.sprites[ani.currentSpriteIndex]))
		for j := 0; j < 4; j++ {
			assert.Equal(t, j, ani.updateCounter)
			ani.Update()
		}
	}
}

func Test_MakeBasicAnimationFromSheet(t *testing.T) {
	ss, _ := MakeCachedSheet(2, 2, 4, 4, "testdata/4px-2x2-small.png")

	a := MakeBasicAnimationFromSheet(ss, []interface{}{0, 1, 2, 3}, 4)
	ani := a.(*BasicAnimation)

	assert.Equal(t, 4, ani.frameLength)

	for i := 0; i < 4; i++ {
		assert.Equal(t, i, ani.currentSpriteIndex)
		assert.Equal(t, fmt.Sprintf("%p", ss.GetSprite(i)), fmt.Sprintf("%p", ani.sprites[ani.currentSpriteIndex]))
		for j := 0; j < 4; j++ {
			assert.Equal(t, j, ani.updateCounter)
			ani.Update()
		}
	}
}
