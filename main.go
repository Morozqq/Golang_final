package main

import (
	"fmt"
)

func main() {
	flowers := []Flower{
		{Name: "Роза", IsDry: true, State: &DryState{}},
		{Name: "Тюльпан", IsDry: false, State: &HealthyState{}},
		{Name: "Лилия", IsDry: true, State: &DryState{}},
		{Name: "Георгина", IsDry: true, State: &DryState{}},
		{Name: "Подсолнух", IsDry: false, State: &HealthyState{}},
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
				flowers[flowerChoice-1].Water()
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
