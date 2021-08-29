package util

import (
	"fmt"
	"io"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

//CheckIsImage check the file is image
func CheckIsImage(f io.Reader) error {
	mType, err := mimetype.DetectReader(f)
	if err != nil {
		return err
	}

	if !strings.HasPrefix(mType.String(), "image") {
		return fmt.Errorf("file is not an image")
	}

	return nil
}
