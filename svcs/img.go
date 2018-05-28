package svcs

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"strings"
)

// ReadB64 decodes base64 string to image
func ReadB64(enc string) (image.Image, error) {
	cur := strings.Index(enc, ";base64,")
	if cur < 0 {
		return nil, errors.New("unsupported mime type")
	}

	mime := enc[5:cur]
	unbased, err := base64.StdEncoding.DecodeString(string(enc[8+cur:]))

	if err != nil {
		return nil, err
	}

	blob := bytes.NewReader(unbased)

	switch mime {

	case "image/png":
		return png.Decode(blob)

	case "image/jpeg":
		return jpeg.Decode(blob)

	default:
		return nil, errors.New("unsupported mime type")
	}
}
