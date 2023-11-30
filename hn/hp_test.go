package hn

import (
	"testing"
	"time"
)

func TestSquareDigits(t *testing.T) {
	type ExampleValue struct {
		value    int64
		expected int64
	}
	tests := []ExampleValue{
		{
			value:    9,
			expected: 81,
		},
		{
			value:    99,
			expected: 162,
		},
		{
			value:    1000,
			expected: 1,
		},
	}

	for _, test := range tests {
		start := time.Now()
		p := NewHappyNumber()
		a := p.SquereDigits(test.value)

		if a == test.expected {
			t.Logf("%d was tasted in %v", test.value, time.Since(start))
		} else {
			t.Errorf("Fail in test %d is expected %d the result was %d", test.value, test.expected, a)
		}
	}

}

func TestIsHappyNumber(t *testing.T) {
	type ExampleValue struct {
		value    int64
		expected bool
	}

	tests := []ExampleValue{
		{
			value:    1,
			expected: true,
		},
		{
			value:    7,
			expected: true,
		},
		{
			value:    10,
			expected: true,
		},
		// {
		// 	value:    13,
		// 	expected: true,
		// },
		// {
		// 	value:    19,
		// 	expected: true,
		// },
		// {
		// 	value:    23,
		// 	expected: true,
		// },
		// {
		// 	value:    28,
		// 	expected: true,
		// },
		// {
		// 	value:    31,
		// 	expected: true,
		// },
		// {
		// 	value:    32,
		// 	expected: true,
		// },
		// {
		// 	value:    44,
		// 	expected: true,
		// },
		// {
		// 	value:    89,
		// 	expected: false,
		// },
	}

	for _, test := range tests {
		p := NewHappyNumber()
		start := time.Now()
		a := p.IsHappyNumber(test.value)

		if a == test.expected {
			t.Logf("%d was tasted in %v", test.value, time.Since(start))
		} else {
			t.Errorf("Fail in test %d is expected %v", test.value, test.expected)
		}
	}
}

func TestCountHappyNumbers(t *testing.T) {

}
