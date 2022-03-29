package main

import "fmt"

// Step 1: define a get method for all decoders
type DecodeMethod interface {
	getDecodeImage(image string) DecodeImage
}

// Step 2: Define decode type to get a suitable method
type DecodeType int

// Step 2.1: Define table to summary all decoders
const (
	GIF DecodeType = iota
	JPEG
)

// Step 3: Create a concrete class
type decodeImage struct {
	image string
}

// Step 3.1: Implement Interfaces to build from concrete class
type DecodeImage interface {
	init(image string)
	toString() string
}

// Step 3.2: Implement methods from interface
func (d *decodeImage) init(image string) {
	d.image = image
}

func (d decodeImage) toString() string {
	return d.image + ": is decoded"
}

// Step 4: Define struct decoder method
type GifReader struct {
	// Fix issue 1: using next to structure
	next decodeImage
}
type JpegReader struct {
	// Fix issue 1: using next to structure
	next decodeImage
}

// Step 4: Implement
func (g *GifReader) getDecodeImage(image string) DecodeImage {
	fmt.Println("get gif decoder success with image " + image)
	return &g.next
}

func (j *JpegReader) getDecodeImage(image string) DecodeImage {
	fmt.Println("get jpeg decoder success with image" + image)

	return &j.next
}

// Step 5: Select decoder from table
// Issue 1: How to fix, missing a return type from struct
func GetDecodeMethod(t DecodeType) DecodeMethod {
	switch t {
	case GIF:
		return new(GifReader)
	case JPEG:
		return new(JpegReader)
	default:
		return nil
	}
}

func main() {

	image := "test.gif"
	format := "gif"

	if format == "gif" {
		reader := GetDecodeMethod(GIF)
		fmt.Print(reader.getDecodeImage(image).toString())
	}

	if format == "jpeg" {
		reader := GetDecodeMethod(JPEG)
		fmt.Print(reader.getDecodeImage(image).toString())
	}
}
