package main

func main() {
	//test of command pattern
	rose := &Flower{Name: "Роза", IsDry: true}
	waterCommand := &WaterFlowerCommand{Flower: rose}

	gardener := &GardenerInvoker{}
	gardener.AddCommand(waterCommand)

	gardener.ExecuteCommands()
}
