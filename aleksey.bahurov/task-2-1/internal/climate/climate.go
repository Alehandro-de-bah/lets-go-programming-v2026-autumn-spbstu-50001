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

// CreateConstraint - функция, возвращающая объект Constraint
// на основе знака сравнения cmp и температуры temp.
func CreateConstraint(cmp string, temp int) Constraint {
	return Constraint{
		isLowerBound: cmp == ">=",
		value:        temp,
	}
}

// Range - структура для хранения диапазона [low, high].
// low - нижняя граница.
// high - верхняя граница.
type Range struct {
	low  int
	high int
}
