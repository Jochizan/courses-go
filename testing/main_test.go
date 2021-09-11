package main

import "testing"

func TestSum(t *testing.T) {
	tables := []struct {
		x int
		y int
		r int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{3, 2, 5},
		{25, 26, 51},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.r {
			t.Errorf("Sum(%d, %d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.r)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		x int
		y int
		r int
	}{
		{1, 2, 2},
		{2, 2, 2},
		{3, 2, 3},
		{25, 26, 26},
	}

	for _, table := range tables {
		max := GetMax(table.x, table.y)
		if max != table.r {
			t.Errorf("Max(%d, %d) was incorrect, got: %d, want: %d.", table.x, table.y, max, table.r)
		}
	}
}
