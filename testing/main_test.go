package main

import "testing"
import "fmt"

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

func TestFib(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tables {
		fib := Fibonacci(item.a)
		if fib != item.n {
			t.Errorf("Fibonacci was incorrect, got: %d, want: %d", fib, item.n)
		}
	}
}
