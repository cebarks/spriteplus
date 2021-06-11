package tileset

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

type Tileset struct {
	Image       image.Image
	ImageSource string `json:"image"`

	Rows    int `json:"rows"`
	Columns int `json:"columns"`

	TileHeight int `json:"tile-height"`
	TileWidth  int `json:"tile-width"`

	Index []index `json:"index"`
}

type index struct {
	Id string `json:"id"`
	X  int    `json:"x"`
	Y  int    `json:"y"`
}

func TilesetFromPath(path string) (*Tileset, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	ts, err := TilesetFromReader(filepath.Dir(path), file)
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func TilesetFromReader(basePath string, r io.Reader) (*Tileset, error) {
	var ts *Tileset

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &ts)
	if err != nil {
		fmt.Printf("err (%T): %s\n", err, err.Error())
		return nil, err
	}

	imageFile, err := os.Open(path.Join(basePath, ts.ImageSource))
	if err != nil {
		return nil, err
	}

	ts.Image, _, err = image.Decode(imageFile)
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func (ts Tileset) Ids() []string {
	ids := make([]string, len(ts.Index))

	for _, idx := range ts.Index {
		idIndex := idx.Y*ts.Rows + idx.X
		ids[idIndex] = idx.Id
	}

	return ids
}
