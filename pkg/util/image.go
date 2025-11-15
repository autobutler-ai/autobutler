package util

import (
	"fmt"
	"image"
	"io"
	"os"

	"github.com/KononK/resize"
	"github.com/rwcarlsen/goexif/exif"
)

func ImageToThumbnail(filePath string, width, height uint) (image.Image, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", fmt.Errorf("error opening image file %s: %w", filePath, err)
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return nil, "", fmt.Errorf("error decoding image file %s: %w", filePath, err)
	}

	img, _ = CorrectImageOrientation(img, file)

	thumbnail := resize.Resize(width, height, img, resize.Lanczos3)
	return thumbnail, format, nil
}

// CorrectImageOrientation reads EXIF orientation data and rotates/flips the image accordingly.
// This ensures images from cameras display correctly regardless of how the camera was held.
func CorrectImageOrientation(img image.Image, r io.ReadSeeker) (image.Image, error) {
	// Reset reader to beginning
	if _, err := r.Seek(0, 0); err != nil {
		// If seek fails, just return the original image
		return img, nil
	}

	// Try to decode EXIF data
	x, err := exif.Decode(r)
	if err != nil {
		// No EXIF data or error reading it - just return original image
		return img, nil
	}

	// Get orientation tag
	tag, err := x.Get(exif.Orientation)
	if err != nil {
		// No orientation tag - return original image
		return img, nil
	}

	orientation, err := tag.Int(0)
	if err != nil {
		return img, nil
	}

	// Apply the transformation based on EXIF orientation
	// http://sylvana.net/jpegcrop/exif_orientation.html
	switch orientation {
	case 1:
		// Normal - do nothing
		return img, nil
	case 2:
		// Flipped horizontally
		return flipHorizontal(img), nil
	case 3:
		// Rotated 180°
		return rotate180(img), nil
	case 4:
		// Flipped vertically
		return flipVertical(img), nil
	case 5:
		// Flipped horizontally and rotated 90° CCW
		return rotate270(flipHorizontal(img)), nil
	case 6:
		// Rotated 90° CW
		return rotate90(img), nil
	case 7:
		// Flipped horizontally and rotated 90° CW
		return rotate90(flipHorizontal(img)), nil
	case 8:
		// Rotated 90° CCW
		return rotate270(img), nil
	}

	return img, nil
}

func rotate90(img image.Image) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(image.Rect(0, 0, bounds.Dy(), bounds.Dx()))
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			newImg.Set(bounds.Max.Y-y-1, x, img.At(x, y))
		}
	}
	return newImg
}

func rotate180(img image.Image) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			newImg.Set(bounds.Max.X-x-1, bounds.Max.Y-y-1, img.At(x, y))
		}
	}
	return newImg
}

func rotate270(img image.Image) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(image.Rect(0, 0, bounds.Dy(), bounds.Dx()))
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			newImg.Set(y, bounds.Max.X-x-1, img.At(x, y))
		}
	}
	return newImg
}

func flipHorizontal(img image.Image) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			newImg.Set(bounds.Max.X-x-1, y, img.At(x, y))
		}
	}
	return newImg
}

func flipVertical(img image.Image) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			newImg.Set(x, bounds.Max.Y-y-1, img.At(x, y))
		}
	}
	return newImg
}
