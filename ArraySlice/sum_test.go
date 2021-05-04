package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got '%d' want '%d' given, '%v'", got, want, numbers)
		}
	})

	// go test -cover を試す
	// このテストは冗長であることが分かる
	// t.Run("collection of any size", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3}
	// 	got := Sum(numbers)
	// 	want := 6
	// 	if got != want {
	// 		t.Errorf("got '%d' want '%d' given, '%v'", got, want, numbers)
	// 	}
	// })
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	// reflect.DeepEqual は型安全ではない
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}