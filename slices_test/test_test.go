package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArraySlice(t *testing.T) {
	tests := []struct {
		name     string
		in       [2]int
		expected [2]int
	}{
		{
			name:     "first test",
			in:       [2]int{1, 2},
			expected: [2]int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, testSlice(tt.in))
		})
	}
}

func TestArrayChangeSlice(t *testing.T) {
	tests := []struct {
		name     string
		in       [2]int
		expected [2]int
	}{
		{
			name:     "change array test",
			in:       [2]int{1, 2},
			expected: [2]int{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, changeArray(tt.in))
		})
	}
}

func TestUseOtherArray(t *testing.T) {
	tests := []struct {
		name     string
		in       [2]int
		expected [2]int
	}{
		{
			name:     "use other array test",
			in:       [2]int{1, 2},
			expected: [2]int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, useOtherArray(tt.in))
		})
	}
}

func TestChangeSliceThroughFunc(t *testing.T) {
	tests := []struct {
		name     string
		in       [2]int
		expected [2]int
	}{
		{
			name:     "change slice through func",
			in:       [2]int{1, 2},
			expected: [2]int{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, changeSliceThroughFunc(tt.in))
		})
	}
}

func TestChangeTwoSlices(t *testing.T) {
	tests := []struct {
		name     string
		in       [2]int
		expected [2]int
	}{
		{
			name:     "change slice through func",
			in:       [2]int{1, 2},
			expected: [2]int{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, changeTwoSlices(tt.in))
		})
	}
}

func TestChangeTwoSlicesWithDifferentArrays(t *testing.T) {
	tests := []struct {
		name     string
		in       [2]int
		expected [2]int
	}{
		{
			name:     "change slice through func",
			in:       [2]int{1, 2},
			expected: [2]int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, changeTwoSlicesWithDifferentArrays(tt.in))
		})
	}
}

func TestChangeSlicesWithTwoArrays(t *testing.T) {
	tests := []struct {
		name     string
		in       [2]int
		expected [2]int
	}{
		{
			name:     "change slice through func",
			in:       [2]int{1, 2},
			expected: [2]int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, changeSliceWithDTwoArrays(tt.in))
		})
	}
}

func testSlice(a [2]int) [2]int {
	s := a[:]
	s = append(s, 3)
	return a
}

func changeArray(a [2]int) [2]int {
	s := a[:1]
	s = append(s, 3)
	return a
}

func useOtherArray(a [2]int) [2]int {
	b := a
	s := b[:1]
	s = append(s, 3)
	return a
}

func changeSliceThroughFunc(a [2]int) [2]int {
	s := a[:1]
	func(s []int) {
		s = append(s, 3)
		s = append(s, 4)
	}(s)
	return a
}

func changeTwoSlices(a [2]int) [2]int {
	s := a[:1]
	func(s []int) {
		b := append(s, 3)
		fmt.Println(b)
	}(s)
	return a
}

func changeTwoSlicesWithDifferentArrays(a [2]int) [2]int {
	s := a[:]
	_ = append(s, 3)
	return a
}

func changeSliceWithDTwoArrays(a [2]int) [2]int {
	b := a
	s := b[:1]
	s = append(s, 3)
	return a
}
