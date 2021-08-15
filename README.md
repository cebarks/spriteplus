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

### Sprite Sheets

```golang
//Create your sheet
sheet := spriteplus.NewSpriteSheet(false)

//Add your sprites to the sheet
err := sheet.AddSprite(gopherSprite)
if err != nil {
  ...
}

//Optimize the texture
sheet.Optimize()


//These can be directly drawn to a Window (or any pixel.Target)
sprite1 := sheet.GetSprite(0) //bottom-left
sprite1.Draw(win, pixel.IM)

// or you can efficiently draw them using a batch with sheet.SourcePic()
pic := sheet.SourcePic()
batch := pixel.NewBatch(&pixel.TrianglesData{}, pic)

sprite1.Draw(batch, pixel.IM)
sprite2.Draw(batch, pixel.IM)
sprite3.Draw(batch, pixel.IM)
sprite4.Draw(batch, pixel.IM)


batch.Draw(win, pixel.IM)
```
---

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

---

## License
[GNU GPLv3](https://choosealicense.com/licenses/gpl-3.0/)
