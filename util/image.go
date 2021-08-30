package util

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

//CheckIsImage check the file is an image. returns of recycled reader
func CheckIsImage(input io.Reader) (io.Reader, error) {
	// header will store the bytes mimetype uses for detection.
	header := bytes.NewBuffer(nil)

	// After DetectReader, the data read from input is copied into header.
	mType, err := mimetype.DetectReader(io.TeeReader(input, header))
	if err != nil {
		return nil, err
	}

	if !strings.HasPrefix(mType.String(), "image") {
		return nil, fmt.Errorf("file is not an image")
	}

	// Concatenate back the header to the rest of the file.
	// recycled now contains the complete, original data.
	recycled := io.MultiReader(header, input)
	return recycled, nil
}
