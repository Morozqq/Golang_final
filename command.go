package main

import "fmt"

type Command interface {
	Execute()
}

// Watering
type WaterFlowerCommand struct {
	Flower *Flower
}

func (w *WaterFlowerCommand) Execute() {
	w.Flower.Water()
}

type Flower struct {
	Name  string
	IsDry bool
	State FlowerState // Change the type of State to FlowerState
}

func (f *Flower) Water() {
	fmt.Printf("Watering the flower %s\n", f.Name)
	f.IsDry = false
	f.State.Water(f) // You need to call the Water method on the State
}

type GardenerInvoker struct {
	Commands []Command
}

func (g *GardenerInvoker) AddCommand(cmd Command) {
	g.Commands = append(g.Commands, cmd)
}

func (g *GardenerInvoker) ExecuteCommands() {
	for _, cmd := range g.Commands {
		cmd.Execute()
	}
}

// Adding a flower
type AddFlowerCommand struct {
	Flowers   *[]Flower
	NewFlower Flower
}

func (a *AddFlowerCommand) Execute() {
	*a.Flowers = append(*a.Flowers, a.NewFlower)
	fmt.Printf("Added a new flower: %s\n", a.NewFlower.Name)
}

// Removing a flower
type RemoveFlowerCommand struct {
	Flowers     *[]Flower
	FlowerIndex int
}

func (r *RemoveFlowerCommand) Execute() {
	if r.FlowerIndex >= 0 && r.FlowerIndex < len(*r.Flowers) {
		removedFlower := (*r.Flowers)[r.FlowerIndex]
		*r.Flowers = append((*r.Flowers)[:r.FlowerIndex], (*r.Flowers)[r.FlowerIndex+1:]...)
		fmt.Printf("Removed the flower: %s\n", removedFlower.Name)
	} else {
		fmt.Println("Invalid flower index. Flower not removed.")
	}
}
