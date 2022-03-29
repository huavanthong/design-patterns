package main

type Image interface {
	display()
}

type HighResolutionImage struct {
	imageFilePath string
}

type ImageProxy struct {
	imageFilePath string
	realImage     Image
}

func (this *HighResolutionImage) loadImage(path string) {
	// load image from disk into memory
	// this is a heavy and costly operation
	fmt.Printf("load image %s from disk\n", path)
}

func (this *HighResolutionImage) display() {
	this.loadImage(this.imageFilePath)
	fmt.Printf("display high resolution image\n")
}

func (this *ImageProxy) display() {
	this.realImage = &HighResolutionImage{imageFilePath: this.imageFilePath}
	this.realImage.display()
}

func NewImageProxy(path string) Image {
	return &ImageProxy{imageFilePath: path}
}

// usage

highResolutionImage := NewImageProxy("sample/img1.png")
// the realImage won't be loaded until user calls display
// later
highResolutionImage.display()