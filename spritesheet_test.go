package spriteplus

import (
	"testing"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/stretchr/testify/assert"
)

func Test_NewBasicSpriteSheet(t *testing.T) {
	ss, err := NewBasicSheet(2, 2, 4, 4, "testdata/4px-2x2-small.png")
	assert.Nil(t, err)

	assert.Equal(t, ss.SourcePic().Bounds(), pixel.R(0, 0, 8, 8))
}

func Test_rectForId(t *testing.T) {
	ss, _ := NewBasicSheet(2, 2, 4, 4, "testdata/4px-2x2-small.png")
	bss := ss.(BasicSpriteSheet)
	assert.Equal(t, pixel.R(0, 0, 4, 4), bss.rectForId(0))
	assert.Equal(t, pixel.R(4, 0, 8, 4), bss.rectForId(1))
	assert.Equal(t, pixel.R(0, 4, 4, 8), bss.rectForId(2))
	assert.Equal(t, pixel.R(4, 4, 8, 8), bss.rectForId(3))
}

//TODO figure out how to test GetSprite
func Test_GetSprite(t *testing.T) {
	t.Skip()
	// ss, _ := NewBasicSheet(2, 2, 4, 4, "testdata/4px-2x2-small.png")

	// pd := pixel.MakePictureData(pixel.R(0, 0, 4, 4))

	// target := pixel.NewBatch(&pixel.TrianglesData{}, pd)
	// target := pixelgl.NewGLPicture()

	// s := ss.GetSprite(0)
	// assert.Equal(t, []uint8{}, pd.Pix)
	// s = ss.GetSprite(1)
	// s.Draw(target, pixel.IM)
	// assert.Equal(t, []uint8{}, pd.Pix)
	// s = ss.GetSprite(2)
	// s.Draw(target, pixel.IM)
	// assert.Equal(t, []uint8{}, pd.Pix)
	// s = ss.GetSprite(3)
	// s.Draw(target, pixel.IM)
	// assert.Equal(t, []uint8{}, pd.Pix)
}
