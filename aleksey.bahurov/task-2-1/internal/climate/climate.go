package climate

import "errors"

type constraint struct {
	isLowerBound bool
	value        int
}

func NewConstraint(cmp string, temp int) constraint {
	return constraint{
		isLowerBound: cmp == ">=",
		value:        temp,
	}
}

type Range struct {
	low  int
	high int
}

func NewRange(low, high int) Range {
	return Range{low: low, high: high}
}

func (r *Range) ApplyConstraint(constraint constraint) {
	if constraint.isLowerBound {
		if constraint.value > r.low {
			r.low = constraint.value
		}
		return
	}
	if constraint.value < r.high {
		r.high = constraint.value
	}
}

func (r *Range) GetOptimalTemperature() (int, error) {
	if r.low > r.high {
		return -1, errors.New("Error: temperature range is empty")
	}

	return r.low, nil
}
