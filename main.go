package main

import (
	"fmt"
	"strconv"
	"strings"
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
		fmt.Println("3 - Remove a flower")
		fmt.Println("4 - Create your bouquet")
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
		case 3:
			fmt.Println("Choose a flower to remove:")
			for i, flower := range flowers {
				fmt.Printf("%d - %s\n", i+1, flower.Name)
			}
			var flowerChoice int
			fmt.Print("Enter the number of the flower to remove: ")
			_, err := fmt.Scan(&flowerChoice)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			if flowerChoice > 0 && flowerChoice <= len(flowers) {
				// Create a RemoveFlowerCommand and add it to the invoker.
				removeCommand := &RemoveFlowerCommand{Flowers: &flowers, FlowerIndex: flowerChoice - 1}
				invoker := &GardenerInvoker{}
				invoker.AddCommand(removeCommand)
				invoker.ExecuteCommands()
			} else {
				fmt.Println("Invalid flower choice.")
			}

		case 4:
			fmt.Println("Create your bouquet:")
			if len(flowers) > 0 {
				fmt.Println("Available flowers:")
				for i, flower := range flowers {
					fmt.Printf("%d - %s\n", i+1, flower.Name)
				}
				fmt.Print("Enter the numbers of flowers that will be in your bouquet (comma-separated): ")
				var input string
				_, err := fmt.Scan(&input)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					return
				}
			
				choiceStrings := strings.Split(input, ",")
				var flowerChoices []int

				for _, choiceStr := range choiceStrings {
					choice, err := strconv.Atoi(strings.TrimSpace(choiceStr))
					if err != nil {
						fmt.Printf("Invalid choice: %v\n", choiceStr)
					} else if choice > 0 && choice <= len(flowers) {
						flowerChoices = append(flowerChoices, choice)
					}
				}

				if len(flowerChoices) > 0 {
					fmt.Println("You created your bouquet!")
					fmt.Println("Selected flowers in your bouquet:")
					for i, choice := range flowerChoices {
						fmt.Printf("%d - %s\n", i+1, flowers[choice-1].Name)
					}
				} else {
					fmt.Println("Try to choose other flowers!")
				}
			} else {
				fmt.Println("No available flowers to create a bouquet!")
			}
		case 0:
			fmt.Println("Exiting Flower Care Console.")
			return
		default:
			fmt.Println("Invalid choice. Please choose a valid option.")
		}
	}
}
