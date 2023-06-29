package imageprocessing

import (
	"image"
	"image/color"
)

// goImageToRGBA32 converts an image to a RGBA32 byte slice.
func goImageToRGBA32(img image.Image) []byte {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	ret := make([]byte, width*height*4)

	index := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			ret[index] = uint8(r >> 8)
			ret[index+1] = uint8(g >> 8)
			ret[index+2] = uint8(b >> 8)
			ret[index+3] = uint8(a >> 8)
			index += 4
		}
	}

	return ret
}

// rgb8PaletteToGoImage converts RGB8 data with a color palette to an image.Image.
func rgb8PaletteToGoImage(width, height int, rgb8data []byte, palette color.Palette) image.Image {
	rect := image.Rect(0, 0, width, height)
	ret := image.NewPaletted(rect, palette)

	index := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			ret.SetColorIndex(x, y, rgb8data[index])
			index++
		}
	}

	return ret
}
