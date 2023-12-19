package main

import "testing"

func TestSum(t *testing.T) {
	got := Sum(1, 2)
	want := 3
	if got != want {
		t.Errorf("Sum(1,2) == %d, want %d", got, want)
	}
}

func TestSum1(t *testing.T) {
	data := []struct {
		a, b, res int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{1, -1, 0},
		{1000, 234, 1234},
	}

	for _, d := range data {
		if got := Sum(d.a, d.b); got != d.res {
			t.Errorf("Sum(%d,%d) == %d, want %d", d.a, d.b, got, d.res)
		}
	}
}
