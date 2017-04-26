//
// frImage - Image utils.
//

package frimage

import (
	"image"
	"image/color"
	"image/png"
  	"os"
)

//
// Write - Create an image file.
//
func Write() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
                        // Hardwire to 50% gray for now.
			pixelColor :=  color.Gray{128}
			img.Set(px, py, pixelColor)
		}
	}
	// Change the default png compression level.
	// e := png.Encoder{CompressionLevel: png.BestCompression}
	e := png.Encoder{CompressionLevel: png.NoCompression}
	e.Encode(os.Stdout, img) // NOTE: ignoring errors
}

