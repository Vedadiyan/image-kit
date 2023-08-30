package crop

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Ratios = string

const pattern = `^\d*(\.?)(\d?)*:\d*(\.?)(\d?)*$`

const (
	SQUARE          Ratios = "1:1"
	STANDARD        Ratios = "4:3"
	WIDE            Ratios = "16:9"
	CLASSIC         Ratios = "3:2"
	MEDIUM_FORMAT   Ratios = "5:4"
	ULTRA_WIDE      Ratios = "21:9"
	CINEMATIC_WIDER Ratios = "2.39:1"
	CINEMATIC       Ratios = "1.85:1"
)

var (
	regex *regexp.Regexp
)

func init() {
	regex = regexp.MustCompile(pattern)
}

func Crop(reader io.Reader, ratio string) ([]byte, error) {
	x, y, err := getRatio(ratio)
	if err != nil {
		return nil, err
	}
	inputImage, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	width := inputImage.Bounds().Dx()

	desiredWidth, height := calculateDimensions(width, inputImage.Bounds().Dy(), x, y)

	cropRect := image.Rect(
		(inputImage.Bounds().Dx()-desiredWidth)/2,
		(inputImage.Bounds().Dy()-height)/2,
		(inputImage.Bounds().Dx()+desiredWidth)/2,
		(inputImage.Bounds().Dy()+height)/2,
	)
	croppedImage := inputImage.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(cropRect)
	var buffer bytes.Buffer
	jpeg.Encode(&buffer, croppedImage, nil)
	return buffer.Bytes(), nil
}

func getRatio(ratio string) (float64, float64, error) {
	if !regex.Match([]byte(ratio)) {
		return -1, -1, fmt.Errorf("invalid crop ratio")
	}
	axes := strings.Split(ratio, ":")
	x, err := strconv.ParseFloat(axes[0], 64)
	if err != nil {
		return -1, -1, err
	}
	y, err := strconv.ParseFloat(axes[1], 64)
	if err != nil {
		return -1, -1, err
	}
	return x, y, nil
}

func calculateDimensions(originalWidth, originalHeight int, x, y float64) (int, int) {
	width := float64(originalWidth)
	height := float64(originalHeight)

	originalAspectRatio := width / height

	desiredAspectRatio := x / y

	if desiredAspectRatio > originalAspectRatio {
		height = width * y
		height = height / x
		return int(width), int(height)
	}
	width = height * x
	width = width / y
	return int(width), int(height)
}
