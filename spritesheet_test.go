package spriteplus

import (
	"fmt"
	"testing"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/stretchr/testify/assert"
)

func Test_MakeBasicSpriteSheet(t *testing.T) {
	ss, err := MakeBasicSheet(2, 2, 4, 4, "testdata/4px-2x2-small.png")
	assert.Nil(t, err)

	assert.Equal(t, ss.SourcePic().Bounds(), pixel.R(0, 0, 8, 8))
}

func Test_rectForId(t *testing.T) {
	ss, _ := MakeBasicSheet(2, 2, 4, 4, "testdata/4px-2x2-small.png")
	bss := ss.(BasicSpriteSheet)
	assert.Equal(t, pixel.R(0, 0, 4, 4), bss.rectForId(0))
	assert.Equal(t, pixel.R(4, 0, 8, 4), bss.rectForId(1))
	assert.Equal(t, pixel.R(0, 4, 4, 8), bss.rectForId(2))
	assert.Equal(t, pixel.R(4, 4, 8, 8), bss.rectForId(3))
}

//TODO figure out how to test GetSprite properly
func Test_BasicGetSprite(t *testing.T) {
	ss, _ := MakeBasicSheet(2, 2, 4, 4, "testdata/4px-2x2-small.png")

	// pd := pixel.MakePictureData(pixel.R(0, 0, 4, 4))

	// target := pixel.NewBatch(&pixel.TrianglesData{}, pd)
	// target := pixelgl.NewGLPicture()

	s := ss.GetSprite(0)
	assert.NotNil(t, s)
	s = ss.GetSprite(1)
	assert.NotNil(t, s)
	s = ss.GetSprite(2)
	assert.NotNil(t, s)
	s = ss.GetSprite(3)
	assert.NotNil(t, s)
	// s = ss.GetSprite(0)
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

//TODO figure out how to test GetSprite properly
func Test_CachedGetSprite(t *testing.T) {
	ss, _ := MakeCachedSheet(2, 2, 4, 4, "testdata/4px-2x2-small.png")

	s1 := ss.GetSprite(0)
	s2 := ss.GetSprite(0)

	assert.Equal(t, *s1, *s2)
	assert.Equal(t, fmt.Sprintf("%p", s1), fmt.Sprintf("%p", s2))
}
