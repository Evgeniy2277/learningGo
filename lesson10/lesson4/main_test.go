package main

import (
	"github.com/usergo/insertsort/insertsort"
	"reflect"
	"testing"
)

func TestSort (t *testing.T) {
	want := []int{1, 2, 3, 12, 44, 51, 55}
	got := insertsort.InsertionSort([]int{51, 12, 3, 2, 1, 55, 44})
	test := reflect.DeepEqual(got, want)
	if !test {
		t.Errorf("got %v wont %v", got, want)
	}
}


