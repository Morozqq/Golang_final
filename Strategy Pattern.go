type State interface {
    Water(flower *Flower)
}

type DryState struct{}

func (s *DryState) Water(flower *Flower) {
    fmt.Println("Watering the flower...")
    flower.IsDry = false
}

type HealthyState struct{}

func (s *HealthyState) Water(flower *Flower) {
    fmt.Println("The flower is healthy and doesn't need water.")
}
