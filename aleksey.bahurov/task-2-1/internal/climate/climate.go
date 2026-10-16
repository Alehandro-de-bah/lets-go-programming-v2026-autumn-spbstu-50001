package climate

import "errors"

type constraint struct {
	isLowerBound bool
	value        int
}

func NewConstraint(cmp string, temp int) (constraint, error) {
	if cmp != ">=" && cmp != "<=" {
		return constraint{}, errors.New("Error: invalid comparison operator")
	}
	return constraint{
		isLowerBound: cmp == ">=",
		value:        temp,
	}, nil
}

type temperatureRange struct {
	low  int
	high int
}

func NewTemperatureRange(low, high int) (temperatureRange, error) {
	if low > high {
		return temperatureRange{}, errors.New("Error: invalid temperature range")
	}
	return temperatureRange{low: low, high: high}, nil
}

func (r *temperatureRange) ApplyConstraint(constraint constraint) {
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

func (r *temperatureRange) GetOptimalTemperature() (int, error) {
	if r.low > r.high {
		return -1, errors.New("Error: temperature range is empty")
	}

	return r.low, nil
}
