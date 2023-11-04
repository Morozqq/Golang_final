package main

import "fmt"

type Command interface {
	Execute()
}

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
