# spriteplus

spriteplus is a golang library meant to be used along side the great [Pixel](https://github.com/faiface/pixel) library. It provides a generic SpriteSheet interface, and a few implementations.

---

## Installation

`go get github.com/cebarks/spriteplus`

---

## Usage

```golang
import "github.com/cebarks/spriteplus"
```

```golang
//
sheet := spriteplus.NewBasicSheet(2, 2, 4, 4, "4px-2x2-small.png")
sheet := spriteplus.NewCachedSheet(2, 2, 4, 4, "4px-2x2-small.png")

//These can be directly drawn to a Window
sprite1 := sheet.GetSprite(0) //bottom-left
sprite2 := sheet.GetSprite(1) //bottom-right
sprite3 := sheet.GetSprite(2) //top-left
sprite4 := sheet.GetSprite(3) //top-right

// or you can draw them to a batch (or any pixel.Target) using sheet.SourcePic()
pic := sheet.SourcePic()
batch := pixel.NewBatch(&pixel.TrianglesData{}, pic)

sprite1.Draw(batch, pixel.IM)
sprite2.Draw(batch, pixel.IM)
sprite3.Draw(batch, pixel.IM)
sprite4.Draw(batch, pixel.IM)


batch.Draw(...
```

### Sheet Types
#### `BasicSpriteSheet`
This is the most basic implementation of a sprite sheet. It has not bells or whistles.
The sprites on the sheet must all be square, of the same size, and aligned neatly to a grid.

#### `CachedSpriteSheet`
This is the same as a `BasicSpriteSheet` except the `GetSprite()` method uses a map to cache the `*pixel.Sprite` that are created instead of a new one being created every time the method is called.

### Todo Sheet Types
#### `RuntimeSpriteSheet`
Uses [pixelutils](https://github.com/dusk125/pixelutils)

---

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

---

## License
[GNU GPLv3](https://choosealicense.com/licenses/gpl-3.0/)
