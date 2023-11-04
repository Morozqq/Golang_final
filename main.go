package main

import "fmt"

func main() {
	flowers := []Flower{
		{Name: "Роза", IsDry: true},
		{Name: "Тюльпан", IsDry: false},
		{Name: "Лилия", IsDry: true},
		{Name: "Георгина", IsDry: true},
		{Name: "Подсолнух", IsDry: false},
	}

	waterCommands := make([]Command, 0)
	for _, flower := range flowers {
		waterCommands = append(waterCommands, &WaterFlowerCommand{Flower: &flower})
	}

	fmt.Println("Welcome to the Flower Care Console:")
	for {
		fmt.Println("Options:")
		fmt.Println("1 - Water a flower")
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
				waterCommands[flowerChoice-1].Execute()
			} else {
				fmt.Println("Invalid flower choice.")
			}
		case 0:
			fmt.Println("Exiting Flower Care Console.")
			return
		default:
			fmt.Println("Invalid choice. Please choose a valid option.")
		}
	}
}
