package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got : %d , want : %d , %v ", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {

	t.Run("Collection of slices", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{1, 2, 3, 4, 5}
		got := SumAll(slice1, slice2)
		want := []int{3, 15}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v , want : %v ", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v , want : %v ", got, want)
		}
	}

	t.Run("Tail sum of all the slices", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{1, 2, 3, 4, 5}
		got := SumAllTails(slice1, slice2)
		want := []int{2, 14}
		checkSums(t, got, want)

	})

	t.Run("Safely sum empty slices also", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []int{1, 2, 3, 4, 5}
		got := SumAllTails(slice1, slice2)
		want := []int{0, 14}
		checkSums(t, got, want)

	})
}

// here we have also the solved that problem of might be comparing a slice with a stirng by imposing
// a check on got and want to be the type of []int a slice and hence if want ever becomes a string then
// there will be a compile time error.
