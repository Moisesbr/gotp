package imageutils

import (
	"bytes"
	"errors"
	"image"
	"io/ioutil"
	"log"

	"github.com/koyachi/go-atkinson"
	"github.com/nfnt/resize"
	"golang.org/x/image/tiff"
)

// ConvertFileToByte open file and convert it
func ConvertFileToByte(fileName string) (imageBytes []byte, w int, h int, err error) {
	var data []byte
	// var err error
	if data, err = ioutil.ReadFile(fileName); err != nil {
		log.Fatal(err)
	}
	imgDecoded, _ := tiff.Decode(bytes.NewBuffer(data))
	imgBytes, w, h, err := ImageToByte(imgDecoded)
	return imgBytes, w, h, err
	// fmt.Println(data, w, h, err)
}

// ImageToByte transform image object to byte array
func ImageToByte(img image.Image) (data []byte, w int, h int, err error) {
	width := 256
	resizedImage := resize.Resize(uint(width), 0, img, resize.Lanczos3)
	height := resizedImage.Bounds().Size().Y
	width = resizedImage.Bounds().Size().X
	if height%8 != 0 {
		return nil, -1, -1, errors.New("image height must be a multiple of 8")
	}
	if width%8 != 0 {
		return nil, -1, -1, errors.New("image width must be a multiple of 8")
	}

	// dither
	ditheredImage, err := atkinson.Dither(resizedImage)
	if err != nil {
		return nil, 0, 0, err
	}

	// convert to byte array
	rowBytes := (width + 7) / 8
	bitmap := make([]byte, rowBytes*height)
	for y := 0; y < height; y++ {
		n := y * rowBytes
		x := 0
		for b := 0; b < rowBytes; b++ {
			sum := 0
			bit := 128
			for bit > 0 {
				if x >= width {
					break
				}
				// r, g, b, a := ditheredImage.At(x, y).RGBA()
				r, _, _, _ := ditheredImage.At(x, y).RGBA()
				if r < 200 {
					// fmt.Println(r, g, b, a)
					sum |= bit
				}
				x += 1
				bit >>= 1
			}
			bitmap[n+b] = byte(sum)
		}
	}
	return bitmap, width, height, nil
}
