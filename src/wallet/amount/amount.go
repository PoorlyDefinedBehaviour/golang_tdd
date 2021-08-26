package amount

import (
	"fmt"

	"github.com/pkg/errors"
)

type T struct {
	value int
}

func (amount *T) String() string {
	return fmt.Sprintf("Amount(%d)", amount.value)
}

var ErrInvalidAmount = errors.New("invalid amount")

func New(value int) (T, error) {
	var amount T

	if value < 0 {
		return amount, errors.Wrapf(ErrInvalidAmount, "negative value: %d", value)
	}

	amount = T{
		value: value,
	}

	return amount, nil
}

func Add(a, b T) T {
	result, _ := New(a.value + b.value)
	if Int(result) < 0 {
		panic(errors.WithStack(ErrInvalidAmount))
	}
	return result
}

func Int(amount T) int {
	return amount.value
}
