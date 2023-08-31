package scale

import (
	"os"
	"testing"
)

func TestScale(t *testing.T) {
	image, err := os.Open("C:\\Users\\Pouya\\Downloads\\sample-city-park-400x300.jpg")
	if err != nil {
		t.FailNow()
	}
	result, err := Scale(image, Percentage(100))
	if err != nil {
		t.FailNow()
	}

	os.WriteFile("result.jpeg", result, os.ModePerm)
}
