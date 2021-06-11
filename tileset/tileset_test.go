package tileset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tileset(t *testing.T) {
	ts, err := TilesetFromPath("../testdata/test.json")

	assert.Nil(t, err, "Error while loading tileset from path.")

	assert.Equal(t, "4px-2x2-small.png", ts.ImageSource)
	assert.Equal(t, 4, ts.TileHeight)
	assert.Equal(t, 4, ts.TileWidth)
	assert.Equal(t, 2, ts.Rows)
	assert.Equal(t, 2, ts.Columns)

	expectedIndexes := []index{
		{Id: "red", X: 0, Y: 0},
		{Id: "green", X: 1, Y: 0},
		{Id: "blue", X: 0, Y: 1},
		{Id: "yellow", X: 1, Y: 1},
	}

	assert.Equal(t, expectedIndexes, ts.Index)

	expectedIds := []string{"red", "green", "blue", "yellow"}

	assert.Equal(t, expectedIds, ts.Ids())
}
