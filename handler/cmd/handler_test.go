package cmd

import (
	"testing"

	"fbnoi.com/framework/handler"
)

type TestBag struct {
	val int
}

func (t *TestBag) Add(i int) {
	t.val = t.val + i
}

func TestHandler(t *testing.T) {
	var i = 0
	handler.New[*int]().Then(func(i *int, fn func(*int)) {
		*i = *i + 1
		fn(i)
	}).Then(func(i *int, fn func(*int)) {
		*i = *i + 2
		fn(i)
	}).Then(func(i *int, fn func(*int)) {
		*i = *i + 3
		fn(i)
	}).Handle(&i)

	if i != 6 {
		t.Errorf("expected %d, got %d", 6, i)
	}
}
