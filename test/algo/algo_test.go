package algo

import (
	"reflect"
	"testing"
)

func TestAlgo(t *testing.T) {
	sortedArray := []int{0, 1, 2, 3, 4, 5, 6, 7, 9, 10}

	t.Run("worst case ", func(t *testing.T) {
		numbers := []int{10, 9, 8, 7, 6, 5, 3, 4, 2, 1, 0}
		got := Sort(numbers)
		want := sortedArray
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want  : %v ", got, want)
		}
	})
	t.Run("average case ", func(t *testing.T) {
		numbers := []int{7, 10, 9, 8, 6, 5, 2, 3, 4, 1, 0}
		got := Sort(numbers)
		want := sortedArray
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want  : %v ", got, want)
		}
	})
}
