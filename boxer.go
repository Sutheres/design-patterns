package main

type boxer struct {
	name        string
	weight      int
	weightClass string
}

func (b *boxer) SetName(name string) {
	b.name = name
}

func (b *boxer) Name() string {
	return b.name
}

func (b *boxer) SetWeight(weight int) {
	b.weight = weight
}

func (b *boxer) Weight() int {
	return b.weight
}

func (b *boxer) SetWeightClass(weightClass string) {
	b.weightClass = weightClass
}

func (b *boxer) WeightClass() string {
	return b.weightClass
}
