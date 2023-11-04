package main

import "fmt"

type Command interface {
	Execute()
}

// Sprinkling
type WaterFlowerCommand struct {
	Flower *Flower
}

// Execute method for the sprinkling
func (w *WaterFlowerCommand) Execute() {
	w.Flower.Water()
}

type Flower struct {
	Name  string
	IsDry bool
}

// Water Simulation, sprinkling method
func (f *Flower) Water() {
	fmt.Printf("Watering the flower %s\n", f.Name)
	f.IsDry = false
}

// Execution of command
type GardenerInvoker struct {
	Commands []Command
}

func (g *GardenerInvoker) AddCommand(cmd Command) {
	g.Commands = append(g.Commands, cmd)
}

// ExecuteCommands executes all the commands stored in the invoker
func (g *GardenerInvoker) ExecuteCommands() {
	for _, cmd := range g.Commands {
		cmd.Execute()
	}
}
