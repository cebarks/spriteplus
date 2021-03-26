# spriteplus

spriteplus is a golang library meant to be used along side the great (Pixel)[] library. It provides a generic SpriteSheet interface, and a few implementations.

## Installation

go get github.com/cebarks/spriteplus

## Usage

```golang
import "github.com/cebarks/spriteplus"

sheet := spriteplus.NewBasicSheet(2, 2, 4, 4, "4px-2x2-small.png")

//These can be directly drawn to a pixel.Target
sprite1 := sheet.GetSprite(0) //bottom-left
sprite2 := sheet.GetSprite(1) //bottom-right
sprite3 := sheet.GetSprite(2) //top-left
sprite4 := sheet.GetSprite(3) //top-right

// or you can draw them to a batch using sheet.SourcePic()
pic := sheet.SourcePic()
batch := pixel.NewBatch(&pixel.TrianglesData{}, pic)

sprite1.Draw(batch, pixel.IM)

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[GNU GPLv3](https://choosealicense.com/licenses/gpl-3.0/)
