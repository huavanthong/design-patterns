package main

type woodBuilder struct {
	house
	quality string
}

func NewWoodBuilder() *woodBuilder {
	return &woodBuilder{}
}

func (w *woodBuilder) setWindowType() {
	w.house.windowType = "Wood Window"
}

func (w *woodBuilder) setDoorType() {
	w.house.doorType = "Wood Door"
}

func (w *woodBuilder) setNumFloor() {
	w.house.floor = 1
}

func (w *woodBuilder) WithQuality(q string) IBuilder {
	w.quality = q
	return w
}

func (w *woodBuilder) GetQuality() string {
	return w.quality
}

func (w *woodBuilder) getHouse() house {
	return w.house
}
