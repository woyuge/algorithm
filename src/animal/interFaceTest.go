package animal

type Integer int

func (a Integer) Equal(i Integer) bool {
	return a == i
}

func (a Integer) LessThan(i Integer) bool {
	return a < i
}

func (a Integer) MoreThan(i Integer) bool {
	return a > i
}

func (a *Integer) Increase(i Integer) {
	*a = *a + i
}

func (a *Integer) Decrease(i Integer) {
	*a = *a - i
}
type IntNumber interface {
	Equal(i Integer) bool
	LessThan(i Integer) bool
	MoreThan(i Integer) bool
	Increase(i Integer)
	Decrease(i Integer)
}