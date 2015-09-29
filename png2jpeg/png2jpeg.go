package png2jpeg

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

// Convert converts PNG file into JPEG.
func Convert(originalFile, newFile string, quality int) error {
	// open file which should be converted.
	file, err := os.Open(originalFile)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Printf("Processing image: %v \n", originalFile)

	// open new file.
	jpg, err := os.OpenFile(newFile, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer jpg.Close()

	if err = convertPNGToJPEG(jpg, file, quality); err != nil {
		if err2 := os.Remove(newFile); err2 != nil {
			fmt.Fprintln(os.Stderr, err2)
		}

		return err
	}

	fmt.Printf("Image: %v Converted!\n", newFile)
	return nil
}

func convertPNGToJPEG(w io.Writer, r io.Reader, quality int) error {
	img, err := png.Decode(r)
	if err != nil {
		return err
	}
	return jpeg.Encode(w, img, &jpeg.Options{Quality: quality})
}
