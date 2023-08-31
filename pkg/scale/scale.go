package scale

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"

	"golang.org/x/image/draw"
)

type Mode interface {
	value() int
}

type Width int
type Height int
type Percentage int

func (width Width) value() int {
	return int(width)
}

func (height Height) value() int {
	return int(height)
}

func (percentage Percentage) value() int {
	return int(percentage)
}

func Scale(reader io.Reader, mode Mode) ([]byte, error) {
	inputImage, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	var x float64
	switch mode.(type) {
	case Width:
		{
			x = float64(mode.value()) / float64(inputImage.Bounds().Dx())
		}
	case Height:
		{
			x = float64(mode.value()) / float64(inputImage.Bounds().Dy())
		}
	case Percentage:
		{
			x = float64(mode.value()) / 100
		}
	default:
		{
			return nil, fmt.Errorf("unknown axis type")
		}
	}

	newWidth := float64(inputImage.Bounds().Dx()) * x
	newHeight := float64(inputImage.Bounds().Dy()) * x
	resizedImage := image.NewRGBA(image.Rect(0, 0, int(newWidth), int(newHeight)))
	draw.CatmullRom.Scale(resizedImage, resizedImage.Bounds(), inputImage, inputImage.Bounds(), draw.Over, nil)
	var buffer bytes.Buffer
	jpeg.Encode(&buffer, resizedImage, nil)
	return buffer.Bytes(), nil
}
