package main

import (
	"fmt"
)

func main() {
	flowers := []Flower{
		{Name: "Rose", IsDry: true, State: &DryState{}},
		{Name: "Tulip", IsDry: false, State: &HealthyState{}},
		{Name: "Lily", IsDry: true, State: &DryState{}},
		{Name: "Dahlia", IsDry: true, State: &DryState{}},
		{Name: "Sunflower", IsDry: false, State: &HealthyState{}},
	}

	fmt.Println("Welcome to the Flower Care Console:")
	for {
		fmt.Println("Options:")
		fmt.Println("1 - Water a flower")
		fmt.Println("2 - Add a new flower")
		fmt.Println("0 - Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		switch choice {
		case 1:
			fmt.Println("Choose a flower to water:")
			for i, flower := range flowers {
				fmt.Printf("%d - %s\n", i+1, flower.Name)
			}
			var flowerChoice int
			fmt.Print("Enter the number of the flower to water: ")
			_, err := fmt.Scan(&flowerChoice)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			if flowerChoice > 0 && flowerChoice <= len(flowers) {
				// Create a WaterFlowerCommand and add it to the invoker.
				waterCommand := &WaterFlowerCommand{Flower: &flowers[flowerChoice-1]}
				invoker := &GardenerInvoker{}
				invoker.AddCommand(waterCommand)
				invoker.ExecuteCommands()
			} else {
				fmt.Println("Invalid flower choice.")
			}
		case 2:
			// Get the details of the new flower and create an AddFlowerCommand.
			var newFlowerName string
			fmt.Print("Enter the name of the new flower: ")
			_, err := fmt.Scan(&newFlowerName)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			newFlower := Flower{Name: newFlowerName, IsDry: true}
			addCommand := &AddFlowerCommand{Flowers: &flowers, NewFlower: newFlower}

			// Add the AddFlowerCommand to the invoker and execute it.
			invoker := &GardenerInvoker{}
			invoker.AddCommand(addCommand)
			invoker.ExecuteCommands()
		case 0:
			fmt.Println("Exiting Flower Care Console.")
			return
		default:
			fmt.Println("Invalid choice. Please choose a valid option.")
		}
	}
}
