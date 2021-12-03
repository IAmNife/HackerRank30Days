package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		mealCost, tipPercent, taxpercent float64
	)

	//Reading the meal cost, tip percent & tax percent
	fmt.Scan(&mealCost)
	fmt.Scan(&tipPercent)
	fmt.Scan(&taxpercent)

	//Calculating the tip, tax & total cost
	tip := mealCost * (tipPercent / 100)
	tax := mealCost * (taxpercent / 100)
	totalCost := mealCost + tax + tip
	var t = int(round(totalCost))
	fmt.Println(t)

}

//Function to round up the amount
func round(num float64) float64 {
	if num < 0 {
		return math.Ceil(num - 0.5)
	}
	return math.Floor(num + 0.5)
}
