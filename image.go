package spriteplus

import (
	"fmt"
	"image"
)

//Subimage returns a subimage of the given image.Image
func Subimage(img image.Image, minX, minY, maxX, maxY int) (image.Image, error) {
	return SubimageRect(img, image.Rect(minX, minY, maxX, maxY))
}

//SubimageRect returns a subimage of the given image.Image
func SubimageRect(img image.Image, rect image.Rectangle) (image.Image, error) {
	var sub image.Image

	switch img := img.(type) {
	case *image.RGBA:
		sub = img.SubImage(rect)
	case *image.RGBA64:
		sub = img.SubImage(rect)
	case *image.NRGBA:
		sub = img.SubImage(rect)
	case *image.NRGBA64:
		sub = img.SubImage(rect)
	case *image.Alpha:
		sub = img.SubImage(rect)
	case *image.Alpha16:
		sub = img.SubImage(rect)
	case *image.Gray:
		sub = img.SubImage(rect)
	case *image.Gray16:
		sub = img.SubImage(rect)
	case *image.CMYK:
		sub = img.SubImage(rect)
	case *image.Paletted:
		sub = img.SubImage(rect)
	default:
		return nil, fmt.Errorf("couldn't cast type of img. Unknown type: %T", img)
	}

	return sub, nil
}
