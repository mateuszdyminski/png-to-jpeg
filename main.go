package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mateuszdyminski/png2jpeg/png2jpeg"
)

// DefaultQuality default quality use for transformation.
const DefaultQuality = 75

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, fmt.Errorf("Please provide params as png2jpeg <PNG_IMAGE> <NEW_NAME_OF_JPEG> <OPTIONAL-QUALITY>"))
		return
	}

	// set quality if provided in args.
	quality := DefaultQuality
	var err error
	if len(os.Args) > 3 {
		quality, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("Quality level should be an integer"))
			return
		}
	}

	if err = png2jpeg.Convert(os.Args[1], os.Args[2], quality); err != nil {
		fmt.Fprintf(os.Stderr, "Can't convert photo! Err: %v \n", err)
	}
}
