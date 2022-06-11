package main

type JuniorMiddleweight struct {
	boxer boxer
}

func NewJuniorMiddleweight(name string) IBoxer {
	return &boxer{
		name:        name,
		weight:      154,
		weightClass: "Junior Middleweight",
	}
}
