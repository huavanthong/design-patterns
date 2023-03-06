package main

type woodBuilder struct {
	house
	quality string
}

func newWoodBuilder() *woodBuilder {
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

func (w *woodBuilder) getHouse() woodBuilder {
	return woodBuilder{house: house{
		doorType:   w.house.doorType,
		windowType: w.house.windowType,
		floor:      w.house.floor,
	},
		quality: "best",
	}
}
