package main

import (
	"fmt"

	"github.com/Alehandro-de-bah/task-2-1/internal/climate"
)

func main() {
	var n int
	fmt.Scan(&n)

	for range n {
		var k int
		if _, err := fmt.Scan(&k); err != nil {
			return
		}

		currentRange := climate.NewRange()

		for range k {
			var (
				comparison  string
				temperature int
			)

			if _, err := fmt.Scan(&comparison, &temperature); err != nil {
				return
			}

			currentRange.ApplyConstraint(climate.NewConstraint(comparison, temperature))

			optimalTemperature := currentRange.GetOptimalTemperature()
			fmt.Println(optimalTemperature)
		}
	}
}
