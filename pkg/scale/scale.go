package scale

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"

	"golang.org/x/image/draw"
)

type Axis interface {
	GetLength() int
}

type Width int
type Height int

func (width Width) GetLength() int {
	return int(width)
}

func (height Height) GetLength() int {
	return int(height)
}

func Scale(reader io.Reader, axis Axis) ([]byte, error) {
	inputImage, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	var x float64
	switch axis.(type) {
	case Width:
		{
			x = float64(axis.GetLength()) / float64(inputImage.Bounds().Dx())
		}
	case Height:
		{
			x = float64(axis.GetLength()) / float64(inputImage.Bounds().Dy())
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
