package climate

const (
	MinTemperature = 15
	MaxTemperature = 30
)

// Constraint - структура, представляющая ограничение снизу или сверху.
// isLowerBound - true, если ограничение снизу (>=), иначе false.
// value - ограничивающее значение.
type Constraint struct {
	isLowerBound bool
	value        int
}
