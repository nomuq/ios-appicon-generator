package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/nfnt/resize"
	"github.com/sirupsen/logrus"
)

func main() {

	if len(os.Args) == 1 {
		logrus.Fatal("No arguments provided")
	}
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		logrus.Fatal(fmt.Sprintf("Error opening file: %s", err))
	}
	defer file.Close()

	// decode png into image.Image
	img, err := png.Decode(file)
	if err != nil {
		logrus.Fatal(err)
	}
	file.Close()

	// create ios icons
	sizes := []float32{
		1024,
		108,
		128,
		16,
		20,
		24,
		256,
		27.5,
		29,
		32,
		40,
		44,
		50,
		512,
		57,
		60,
		72,
		76,
		83.5,
		86,
		98,
	}

	// make sure the output directory exists
	logrus.Info("Creating output directory")
	os.MkdirAll("output", os.ModePerm)

	logrus.Info("Creating icons")
	for _, size := range sizes {
		m := resize.Resize(uint(size), uint(size), img, resize.Lanczos3)
		out, err := os.Create(fmt.Sprintf("output/%dx%d.png", uint(size), uint(size)))
		if err != nil {
			logrus.Fatal(err)
		}
		defer out.Close()
		// write new image to file
		png.Encode(out, m)
		logrus.Info(fmt.Sprintf("Created %dx%d.png", uint(size), uint(size)))
	}
	logrus.Info("icons created successfully")
}
