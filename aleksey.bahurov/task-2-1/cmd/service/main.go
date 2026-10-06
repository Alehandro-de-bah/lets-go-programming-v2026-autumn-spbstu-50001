package main

import (
	"fmt"

	"github.com/Alehandro-de-bah/task-2-1/internal/climate"
)

func main() {
	var departmentsNumber int
	if _, err := fmt.Scan(&departmentsNumber); err != nil {
		return
	}

	for range departmentsNumber {
		var employeesNumber int
		if _, err := fmt.Scan(&employeesNumber); err != nil {
			return
		}

		currentRange := climate.NewRange()

		for range employeesNumber {
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
