package main

import "fmt"

type Car struct {
	Make       string
	Model      string
	Year       int
	Color      string
	EngineSize float64
}

type CarBuilder struct {
	Car
}

func (cb *CarBuilder) SetMake(make string) *CarBuilder {
	cb.Make = make
	return cb
}

func (cb *CarBuilder) SetModel(model string) *CarBuilder {
	cb.Model = model
	return cb
}

func (cb *CarBuilder) SetYear(year int) *CarBuilder {
	cb.Year = year
	return cb
}

func (cb *CarBuilder) SetColor(color string) *CarBuilder {
	cb.Color = color
	return cb
}

func (cb *CarBuilder) SetEngineSize(engineSize float64) *CarBuilder {
	cb.EngineSize = engineSize
	return cb
}

func (cb *CarBuilder) Build() *Car {
	return &cb.Car
}

func main() {

	carBuilder := &CarBuilder{}

	car := carBuilder.
		SetMake("Toyota").
		SetModel("Corolla").
		SetYear(2021).
		SetColor("Red").
		SetEngineSize(1.8).
		Build()

	fmt.Printf("Make: %s\n", car.Make)                // Output: Make: Toyota
	fmt.Printf("Model: %s\n", car.Model)              // Output: Model: Corolla
	fmt.Printf("Year: %d\n", car.Year)                // Output: Year: 2021
	fmt.Printf("Color: %s\n", car.Color)              // Output: Color: Red
	fmt.Printf("Engine Size: %.1f\n", car.EngineSize) // Output: Engine Size: 1.8

}
