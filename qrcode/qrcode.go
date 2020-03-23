package qrcode

import (
	"fmt"
	"os"

	gim "github.com/ozankasikci/go-image-merge"
	"github.com/skip2/go-qrcode"
	"golang.org/x/image/tiff"
)

// Save qrcode tif file
func Save(content string, filename string) error {
	q, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return err
	}
	q.DisableBorder = true
	err = q.WriteFile(128, "generated_QR.png")
	if err != nil {
		return err
	}

	// accepts *Grid instances, grid unit count x, grid unit count y
	// returns an *image.RGBA object
	grids := []*gim.Grid{
		{ImageFilePath: "padding.png"},
		{ImageFilePath: "generated_QR.png"},
	}
	rgba, _ := gim.New(grids, 2, 1).Merge()

	file, _ := os.Create(filename)
	err = tiff.Encode(file, rgba, nil)
	if err != nil {
		fmt.Println("error converting to tiff")
		return err
	}
	return nil
}

// func main() {
// 	SavePadQRCode("https://example.org", "generated_QR1.tif")
// }
