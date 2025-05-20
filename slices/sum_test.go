package slices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of  5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		expected := 15

		got := Sum(numbers)

		if got != expected {
			t.Errorf("got %d want %d given, %v", got, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sums collection of collection", func(t *testing.T) {
		collection1 := []int{1, 2, 3}
		collection2 := []int{9, 8, 7}
		expected := []int{6, 24}

		got := SumAll(collection1, collection2)

		if !slices.Equal(got, expected) {
			t.Errorf("got %v want %v", got, expected)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !slices.Equal(got, expected) {
			t.Errorf("got %v want %v", got, expected)
		}
	}
	t.Run("sums tails of all collections", func(t *testing.T) {
		collection1 := []int{1, 2, 3}
		collection2 := []int{9, 8, 7}
		expected := []int{5, 15}

		got := SumAllTails(collection1, collection2)
		checkSums(t, got, expected)

	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
