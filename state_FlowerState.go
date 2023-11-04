package main

import "fmt"

// FlowerState представляет интерфейс для состояния цветка.
type FlowerState interface {
	Water(flower *Flower)
}

// DryState - состояние "сухой цветок".
type DryState struct{}

func (d *DryState) Water(flower *Flower) {
	fmt.Printf("Watering the dry flower %s\n", flower.Name)
	flower.State = &HealthyState{}
}

// HealthyState - состояние "здоровый цветок".
type HealthyState struct{}

func (h *HealthyState) Water(flower *Flower) {
	fmt.Printf("The flower %s is already healthy and doesn't need watering.\n", flower.Name)
}
