package main

// A couple of type definitions to make the units clear.
type earthMass float64
type au float64

// A Planet defines the properties of a solar system object.
type Emojo struct {
	rang    int
	name    string
	unicode string
	e
	mass     earthMass
	distance au
}

var planets = []Planet{
	{"Mercury", 0.055, 0.4},
	{"Venus", 0.815, 0.7},
	{"Earth", 1.0, 1.0},
	{"Mars", 0.107, 1.5},
}

func main() {}
