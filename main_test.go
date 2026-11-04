package main

import (
	"fmt"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-10, -1, 0, 1, 7, 8, 17, 1000} {
		t.Run(fmt.Sprintf("size_%d", size), func(t *testing.T) {
			data := generateRandomElements(size)
			if size <= 0 {
				if data != nil {
					t.Fatalf("generateRandomElements(%d) = %v, want nil", size, data)
				}
				return
			}
			if len(data) != size {
				t.Fatalf("length = %d, want %d", len(data), size)
			}
			for i, value := range data {
				if value <= 0 {
					t.Errorf("data[%d] = %d, want a positive integer", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
	cases := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{42}, 42},
		{"single_negative", []int{-42}, -42},
		{"zeros", []int{0, 0, 0}, 0},
		{"equal", []int{5, 5, 5}, 5},
		{"ascending", []int{1, 2, 3, 4}, 4},
		{"descending", []int{4, 3, 2, 1}, 4},
		{"middle", []int{1, 9, 3}, 9},
		{"negative", []int{-9, -2, -7}, -2},
		{"mixed", []int{-9, 0, 7, -2}, 7},
		{"min_int", []int{minInt, minInt}, minInt},
		{"integer_limits", []int{minInt, 0, maxInt}, maxInt},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := maximum(tc.data); got != tc.want {
				t.Errorf("maximum = %d, want %d", got, tc.want)
			}
		})
	}
}
