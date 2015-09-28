package main

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

// DefaultQuality default quality use for transformation.
const DefaultQuality = 75

func main() {
	// open file which should be converted.
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// get path of current directory.
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		log.Fatalf("Please provide png file")
	}

	// set quality if provided in args.
	quality := DefaultQuality
	if len(os.Args) > 2 {
		quality, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Quality level should be an integer")
		}
	}

	fmt.Printf("Processing image: %v \n", path.Join(currentDir, "new_"+stat.Name()))

	// open new file.
	newFileName := "new_" + strings.TrimSuffix(stat.Name(), filepath.Ext(stat.Name())) + ".jpg"
	jpg, err := os.OpenFile(path.Join(currentDir, newFileName), os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	convertPNGToJPEG(jpg, file, quality)

	fmt.Printf("Image: %v Converted!\n", newFileName)

	// close files
	file.Close()
	jpg.Close()
}

func convertPNGToJPEG(w io.Writer, r io.Reader, quality int) error {
	img, err := png.Decode(r)
	if err != nil {
		return err
	}
	return jpeg.Encode(w, img, &jpeg.Options{Quality: quality})
}

