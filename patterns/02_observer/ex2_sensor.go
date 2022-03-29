package main

import (
	"fmt"
	"time"
)

// Step 1: Consider alarm listener is Observer
type AlarmListener interface {
	alarm()
}

type alarmListener struct{}

func (alarmListener) alarm() {
	fmt.Println("Alarm ring")
}

// Step 2: Consider sensor system is Subject
type SensorSystem interface {
	register(sensor AlarmListener)
	remove(sensor AlarmListener)
	soundTheAlarm()
}

// This class refer at:
// https://docs.microsoft.com/en-us/previous-versions/windows/desktop/ipmiprv/sensor
type Sensor struct {
	Caption      string
	Description  string
	ElementName  string
	InstallDate  time.Time
	Name         string
	SensorType   uint16
	CurrentState string
}

type sensorSystem struct {
	alarmListeners []AlarmListener
}

func (s *sensorSystem) register(sensor AlarmListener) {
	s.alarmListeners = append(s.alarmListeners, sensor)
}

func (s *sensorSystem) remove(sensor AlarmListener) {

}

func (s sensorSystem) soundTheAlarm() {

	for i := 0; i < len(s.alarmListeners); i++ {
		sensor := s.alarmListeners[i]
		sensor.alarm()
	}
}

// func RemoveIndex(s []alarmListener, index int) []string {
// 	return append(s[:index], s[index+1:]...)
// }

type Lighting struct{}

// To inheritant methods from AlarmListener
func (Lighting) alarm() {
	fmt.Println("lights up")
}

func NewLighting() AlarmListener {
	return &Lighting{}
}

type Gates struct{}

func (Gates) alarm() {
	fmt.Println("gates close")
}

func NewGates() AlarmListener {
	return &Gates{}
}

type checkList struct{}
type CheckList interface {
	byTheNumbers()
	localize()
	isolate()
	identify()
}

func (c checkList) byTheNumbers() {
	c.localize()
	c.isolate()
	c.identify()
}

func (checkList) localize() {
	fmt.Println("   establish a perimeter")
}

func (checkList) isolate() {
	fmt.Println("   isolate the grid")

}

func (checkList) identify() {
	fmt.Println("   identify the source")
}

type Surveillance struct {
	check checkList
}

func (s Surveillance) alarm() {
	fmt.Println("Surveillance - by the numbers:")
	s.check.byTheNumbers()
}

func NewSurveillance() AlarmListener {
	return &Surveillance{}
}

func main() {
	sensorSystem := &sensorSystem{}
	lighting := NewLighting()
	gates := NewGates()
	surveillance := NewSurveillance()

	sensorSystem.register(lighting)
	sensorSystem.register(gates)
	sensorSystem.register(surveillance)
	sensorSystem.soundTheAlarm()
}
