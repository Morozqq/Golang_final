package main

import "fmt"

type bouquet interface {
	AddFlower(flower Flower)
	Describe() string
}

type CompositeBouquet struct {
	Flowers []Flower
}

func (cb *CompositeBouquet) AddFlower(flower Flower) {
	cb.Flowers = append(cb.Flowers, flower)
}

func (cb *CompositeBouquet) Describe() string {
	description := "Bouquet with:\n"
	for i, flower := range cb.Flowers {
		description += fmt.Sprintf("%d - %s\n", i+1, flower.Name)
	}
	return description
}

type flower interface {
	Describe() string
}

type ConcreteFlower struct {
	Name string
}

func (f *ConcreteFlower) Describe() string {
	return f.Name
}
