package main

type IBoxer interface {
	SetName(name string)
	SetWeight(weight int)
	SetWeightClass(weightClass string)
	Name() string
	Weight() int
	WeightClass() string
}
