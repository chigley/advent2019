package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/chigley/advent2019"
)

type (
	SpaceImage []Layer
	Layer      [][]Pixel
	Pixel      int
)

func main() {
	img, err := ReadImage(os.Stdin, 25, 6)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Part1(img))
}

func Part1(img SpaceImage) int {
	var layerIndex int
	zeroes := math.MaxInt64

	for i, l := range img {
		freq := l.frequency(0)
		if freq < zeroes {
			layerIndex = i
			zeroes = freq
		}
	}

	return img[layerIndex].frequency(1) * img[layerIndex].frequency(2)
}

func (l Layer) frequency(p Pixel) (ret int) {
	for _, row := range l {
		for _, pixel := range row {
			if pixel == p {
				ret++
			}
		}
	}
	return
}

func ReadImage(r io.Reader, width, height int) (SpaceImage, error) {
	lines, err := advent2019.ReadStrings(r)
	if err != nil {
		return nil, err
	}
	if len(lines) != 1 {
		return nil, errors.New("expected exactly one input")
	}
	input := lines[0]

	numLayers := len(input) / (width * height)
	if extra := len(input) % (width * height); extra != 0 {
		return nil, fmt.Errorf("unexpected additional %d pixels at end of input", extra)
	}

	img := make(SpaceImage, numLayers)
	for i := 0; i < numLayers; i++ {
		img[i] = make(Layer, height)
		for j := 0; j < height; j++ {
			img[i][j] = make([]Pixel, width)
			for k := 0; k < width; k++ {
				idx := i*width*height + j*width + k

				n, err := strconv.Atoi(string(input[idx]))
				if err != nil {
					return nil, err
				}

				img[i][j][k] = Pixel(n)
			}
		}
	}

	return img, nil
}
