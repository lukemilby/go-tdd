package arrayslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T){

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1,2,3}

		got := Sum(numbers)
		want := 6

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T){

	t.Run("adds to slices together", func(t *testing.T) {
		number_one := []int{1, 2}
		number_two := []int{0, 9}

		got := SumAll(number_one, number_two)
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}


func TestSumTails(t *testing.T){

	checkSums := func(t *testing.T, got, want []int){
		t.Helper()
		if !reflect.DeepEqual(got, want){
			t.Errorf("got %d want %d", got, want)
		}
	}
	t.Run("adds tails of slices", func(t *testing.T) {
		numberOne := []int{1, 2}
		numberTwo := []int{0, 9}

		got := SumAllTails(numberOne, numberTwo)
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("check empty slices", func(t *testing.T) {
		number_one := []int{}
		number_two := []int{0, 9, 11}

		got := SumAllTails(number_one, number_two)
		want := []int{0, 20}

		checkSums(t, got, want)
	})
}