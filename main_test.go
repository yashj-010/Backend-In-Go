package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("processTruck",func(t *testing.T) {
		t.Run("diesel truck",func(t *testing.T) {
			dieselTruck := &DieselTruck{
				id: "D1",
				capacity: 10,
			}
			processTruck(dieselTruck)
		})
		t.Run("electric truck",func(t *testing.T) {
			electricTruck := &ElectricTruck{
				id: "E1",
				capacity: 10,
				batteryLevel: 100,
			}
			processTruck(electricTruck)
		})
	})
}