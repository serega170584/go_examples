package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestInterface interface {
	A()
}

type Test struct{}

func (t *Test) A() {
	fmt.Println("123")
}

func TestGeneral(t *testing.T) {
	tests := []struct {
		name     string
		in       any
		expected any
	}{
		{
			name:     "first test",
			in:       1,
			expected: 1,
		},
		{
			name:     "first test",
			in:       1.00,
			expected: 1.00,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			General(1)
			assert.Equal(t, tt.expected, General(tt.in))
		})
	}
}

func TestFalseInterfaceCompare(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "first test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			General(1)
			assert.Equal(t, false, modifyInterface())
		})
	}
}

func TestInterfacePointerCompare(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "first test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			General(1)
			assert.Equal(t, false, modifyInterface())
		})
	}
}

func TestModifyInterfacePointer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "first test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			General(1)
			assert.Equal(t, true, modifyInterfacePointer())
		})
	}
}

func General(a any) any {
	return a
}

func modifyInterface() bool {
	var a interface{}
	a = 1
	var b interface{}
	b = 1.00
	return a == b
}

func modifyInterfacePointer() bool {
	var a *Test
	//var a TestInterface
	//var a *int
	//var b interface{}
	//b = 1.00
	return nilA(a)
}

func nilA(a TestInterface) bool {
	return a.(*Test) == nil
}
