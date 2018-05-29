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

// ImageWithType transfer object for Image
type ImageWithType struct {
	Image image.Image
	Mime  string
}

// ReadB64 decodes base64 string to image
func ReadB64(enc string) (*ImageWithType, error) {
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
		image, err := png.Decode(blob)
		return &ImageWithType{
			Image: image,
			Mime:  mime,
		}, err

	case "image/jpeg":
		image, err := jpeg.Decode(blob)
		return &ImageWithType{
			Image: image,
			Mime:  mime,
		}, err

	default:
		return nil, errors.New("unsupported mime type")
	}
}

func ReadImage(image ImageWithType) {
	buff := new(bytes.Buffer)

}
