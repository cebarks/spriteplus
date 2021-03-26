# spriteplus

spriteplus is a golang library meant to be used along side the great (Pixel)[] library. It provides a generic SpriteSheet interface, and a few implementations.

## Installation

go get github.com/cebarks/spriteplus

## Usage

```golang
import "github.com/cebarks/spriteplus"

func main() {
    sheet := spriteplus.NewBasicSheet(8, 8, 16, 16, "my-16x16-spritesheet.png")

    sprite := sheet.GetSprite(0)
}

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[GNU GPLv3](https://choosealicense.com/licenses/gpl-3.0/)
