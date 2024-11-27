package task

import (
	"testing"
)

func isEqual(a, b map[string][]string) bool {
	for k, v := range a {
		v2, ok := b[k]
		if !ok {
			return false
		}
		if len(v2) != len(v) {
			return false
		}
		for i := 0; i < len(v); i++ {
			if v[i] != v2[i] {
				return false
			}
		}
	}
	return true
}
func TestTask(t *testing.T) {
	testTable := []struct {
		input    []string
		expected map[string][]string
	}{
		{
			input:    []string{"пятка", "тяпка", "Катяп", "п", "пяТАК", "пятка", "листок", "слиток", "столик"},
			expected: map[string][]string{"пятка": []string{"катяп", "пятак", "пятка", "тяпка"}, "листок": []string{"листок", "слиток", "столик"}},
		},
		{
			input:    []string{"ф", "А", "п", "цЦАУА", "ыуп", "", "ыуа", "уац", "уац"},
			expected: map[string][]string{"цЦАУА": []string{"ццауа"}, "ыуа": []string{"ыуа"}, "уац": []string{"уац"}, "ыуп": []string{"ыуп"}},
		},
	}

	for _, testCase := range testTable {
		result := task(&testCase.input)
		if !isEqual(result, testCase.expected) {
			t.Errorf("Wrong Answer:\nwant - %v\nhave - %v", testCase.expected, result)
		}
	}

}
