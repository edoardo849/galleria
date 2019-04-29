package image

import (
	"image"
	"io"

	"github.com/disintegration/imaging"
)

// Decode reads an image from r.
func Decode(r io.Reader) (image.Image, error) {
	return imaging.Decode(r)
}

// Convert saves the image to file with the specified filename.
// The format is determined from the filename extension: "jpg" (or "jpeg"),
// "png", "gif", "tif" (or "tiff") and "bmp" are supported.
func Convert(w io.Writer, img image.Image, filename string) error {

	f, err := imaging.FormatFromFilename(filename)
	if err != nil {
		return err
	}

	return imaging.Encode(w, img, f)
}
