package main

import (
	"fmt"

	"github.com/Alehandro-de-bah/task-2-1/internal/climate"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var k int
		fmt.Scan(&k)

		r := climate.NewRange()

		for j := 0; j < k; j++ {
			var (
				comparison  string
				temperature int
			)

			fmt.Scan(&comparison, &temperature)

			r.ApplyConstraint(climate.NewConstraint(comparison, temperature))

			optimalTemperature := r.GetOptimalTemperature()
			fmt.Println(optimalTemperature)
		}
	}
}
