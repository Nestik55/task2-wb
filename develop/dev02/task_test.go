package main

import (
	"testing"
)

func isEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestUnpackString(t *testing.T) {
	testTable := []struct {
		input    []rune
		expected []rune
	}{
		{
			input:    []rune("a4bc2d5e"),
			expected: []rune("aaaabccddddde"),
		},
		{
			input:    []rune("abcd"),
			expected: []rune("abcd"),
		},
		{
			input:    []rune(""),
			expected: []rune(""),
		},
		{
			input:    []rune(`qwe\4\5`),
			expected: []rune("qwe45"),
		},
		{
			input:    []rune(`qwe\45`),
			expected: []rune(`qwe44444`),
		},
		{
			input:    []rune(`qwe\\5`),
			expected: []rune(`qwe\\\\\`),
		},
	}

	for _, testCase := range testTable {
		result, err := UnpackString(testCase.input)
		if !isEqual(result, testCase.expected) || err != nil {
			t.Errorf("Incorrect result:\nwant - %v - <nil>\nhave - %v - %v", string(testCase.expected), string(result), err)
		}
	}
}

func TestErrorUnpack(t *testing.T) {
	testTable := [][]rune{[]rune("45"), []rune("5\\")}
	for _, testCase := range testTable {
		_, err := UnpackString(testCase)
		if err == nil {
			t.Errorf("Test:%v\nIncorrect error result:\nwant - Incorrect string\nhave - %v", string(testCase), err)
		}
	}

}
