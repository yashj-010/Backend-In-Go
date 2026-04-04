package main

import (
	"fmt"
)

type Truck interface {
	loadcargo() error
	unloadcargo() error
}

type DieselTruck struct {
	id       string
	capacity int
}

func (d *DieselTruck) loadcargo() error {
	fmt.Println("Loading cargo in diesel truck", d.id)
	return nil
}

func (d *DieselTruck) unloadcargo() error {
	fmt.Println("Unloading cargo from diesel truck", d.id)
	return nil
}

type ElectricTruck struct {
	id           string
	capacity     int
	batteryLevel int
}

func (e *ElectricTruck) loadcargo() error {
	fmt.Println("Loading cargo in electric truck", e.id)
	return nil
}

func (e *ElectricTruck) unloadcargo() error {
	fmt.Println("Unloading cargo from electric truck", e.id)
	return nil
}

func processTruck(t Truck) {
	t.loadcargo()
	t.unloadcargo()
}

func main() {
	dieselTruck := &DieselTruck{
		id:       "D1",
		capacity: 10,
	}

	electricTruck := &ElectricTruck{
		id:           "E1",
		capacity:     10,
		batteryLevel: 100,
	}

	processTruck(dieselTruck)
	processTruck(electricTruck)
}
