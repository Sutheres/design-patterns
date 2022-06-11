package main

type Welterweight struct {
	boxer boxer
}

func NewWelterweight(name string) IBoxer {
	return &boxer{
		name:        name,
		weight:      147,
		weightClass: "Welterweight",
	}
}
