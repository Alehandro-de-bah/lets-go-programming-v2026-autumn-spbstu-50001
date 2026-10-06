package climate

const (
	MinTemperature = 15
	MaxTemperature = 30
)

// Constraint - структура, представляющая ограничение
// снизу или сверху (isLowerBound) заданным значением (value).
type Constraint struct {
	isLowerBound bool
	value        int
}

// CreateConstraint - функция, возвращающая объект Constraint
// на основе знака сравнения cmp и температуры temp.
func NewConstraint(cmp string, temp int) Constraint {
	return Constraint{
		isLowerBound: cmp == ">=",
		value:        temp,
	}
}

// Range - структура для хранения диапазона [low, high].
type Range struct {
	low  int
	high int
}

// CreateRange - функция, возвращающая объект Range
// с границами по умолчанию [MinTemperature, MaxTemperature].
func NewRange() Range {
	return Range{low: MinTemperature, high: MaxTemperature}
}

// ApplyConstraint - метод, применяющий
// ограничение constraint к диапазону r.
func (r *Range) ApplyConstraint(constraint Constraint) {
	if constraint.isLowerBound {
		if constraint.value > r.low {
			r.low = constraint.value
		}
	} else {
		if constraint.value < r.high {
			r.high = constraint.value
		}
	}
}

// GetOptimalTemperature - метод, возвращающий
// оптимальную температуру на основе диапазона r.
func (r *Range) GetOptimalTemperature() int {
	if r.low > r.high {
		return -1
	}

	return r.low
}
