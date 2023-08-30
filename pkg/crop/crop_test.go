package crop

import (
	"os"
	"testing"
)

func TestCrop(t *testing.T) {
	image, err := os.Open("C:\\Users\\Pouya\\Downloads\\sample-city-park-400x300.jpg")
	if err != nil {
		t.FailNow()
	}
	result, err := Crop(image, ULTRA_WIDE)
	if err != nil {
		t.FailNow()
	}

	os.WriteFile("result.jpeg", result, os.ModePerm)
}
