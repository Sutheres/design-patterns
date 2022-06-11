package main

import "fmt"

func main() {
	errolSpence, _ := NewBoxer("Welterweight", "Errol Spence")
	jermellCharlo, _ := NewBoxer("Junior Middleweight", "Jermell Charlo")

	printDetails(errolSpence)
	printDetails(jermellCharlo)
}

func printDetails(b IBoxer) {
	fmt.Printf("Name: %s", b.Name())
	fmt.Println()
	fmt.Printf("Weight: %d lbs.", b.Weight())
	fmt.Println()
	fmt.Printf("Weight class: %s", b.WeightClass())
	fmt.Println()
}
