package greet

import "fmt"

var Shark = "Sammy"

type Octopus struct {
	Name  string
	Color string
}

func (o Octopus) String() string {
	return fmt.Sprintf("The octopus's name is %q and is the color %s.", o.Name, o.Color)
}

func (o Octopus) Reset() {
	o.Name = "23"
	fmt.Println("Print! ", o.Name)
	o.Color = "54"
	fmt.Println("Print! ", o.Color)
}

func Hello() {
	fmt.Println("Hello, World!")
}