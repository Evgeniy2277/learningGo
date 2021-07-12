package main

import (
	"github.com/usergo/fibonacci/fib"
	"testing"
)

func TestFib (t *testing.T) {
	table :=[]struct {
		arg int
		want int
	}{
		{0, 0},
		{1, 1},
		{10, 55},
		{20, 6765},

	}
	for _, entry := range table {
		got := fib.FibCache(entry.arg)

		if got != entry.want {
			t.Errorf("got %d want %d", entry.arg, entry.want)
		}
	}
}
