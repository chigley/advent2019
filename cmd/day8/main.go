package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
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

const (
	Black Pixel = iota
	White
	Transparent
)

var errEmptyImage = errors.New("empty image or zero width or height")

func main() {
	img, err := ReadImage(os.Stdin, 25, 6)
	if err != nil {
		log.Fatal(err)
	}

	part1 := Part1(img)

	part2, err := Part2(img)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
	fmt.Printf("Part 2 output saved to %s\n", part2)
}

func Part1(img SpaceImage) int {
	var layerIndex int
	zeroes := math.MaxInt64

	for i, l := range img {
		freq := l.frequency(Black)
		if freq < zeroes {
			layerIndex = i
			zeroes = freq
		}
	}

	return img[layerIndex].frequency(White) * img[layerIndex].frequency(Transparent)
}

func Part2(img SpaceImage) (path string, err error) {
	tmpfile, err := ioutil.TempFile("", "")
	if err != nil {
		return "", err
	}
	defer func() {
		if cerr := tmpfile.Close(); err == nil {
			err = cerr
		}
	}()

	render, err := img.Render()
	if err != nil {
		return "", err
	}

	if err := render.SavePNG(tmpfile); err != nil {
		return "", err
	}

	return tmpfile.Name(), nil
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

func (img SpaceImage) Render() (Layer, error) {
	if len(img) == 0 || len(img[0]) == 0 || len(img[0][0]) == 0 {
		return nil, errEmptyImage
	}

	// Initialise a transparent image
	height, width := len(img[0]), len(img[0][0])
	ret := make(Layer, height)
	for i := 0; i < height; i++ {
		ret[i] = make([]Pixel, width)
		for j := 0; j < width; j++ {
			ret[i][j] = Transparent
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if ret[i][j] != Transparent {
				continue
			}

			for _, layer := range img {
				if layer[i][j] != Transparent {
					ret[i][j] = layer[i][j]
					break
				}
			}
		}
	}

	return ret, nil
}

func (l Layer) SavePNG(w io.Writer) error {
	if len(l) == 0 || len(l[0]) == 0 {
		return errEmptyImage
	}

	height, width := len(l), len(l[0])

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			col, err := l[i][j].Color()
			if err != nil {
				return err
			}
			img.Set(j, i, col)
		}
	}

	return png.Encode(w, img)
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

func (p Pixel) Color() (color.Color, error) {
	switch p {
	case White:
		return color.White, nil
	case Black:
		return color.Black, nil
	case Transparent:
		return color.Transparent, nil
	default:
		return nil, fmt.Errorf("unknown pixel value %d", p)
	}
}
