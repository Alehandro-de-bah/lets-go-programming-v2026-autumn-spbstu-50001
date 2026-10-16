package climate

const (
	MinTemperature = 15
	MaxTemperature = 30
)

type Constraint struct {
	isLowerBound bool
	value        int
}

func NewConstraint(cmp string, temp int) Constraint {
	return Constraint{
		isLowerBound: cmp == ">=",
		value:        temp,
	}
}

type Range struct {
	low  int
	high int
}

func NewRange() Range {
	return Range{low: MinTemperature, high: MaxTemperature}
}

func (r *Range) ApplyConstraint(constraint Constraint) {
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

func (r *Range) GetOptimalTemperature() int {
	if r.low > r.high {
		return -1
	}

	return r.low
}
