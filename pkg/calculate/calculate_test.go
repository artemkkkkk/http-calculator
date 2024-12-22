package calculate

import (
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr      string
		expected  float64
		shouldErr bool
	}{
		{"2 + 2", 4.0, false},
		{"3 * 3", 9.0, false},
		{"10 / 2", 5.0, false},
		{"5 - 3", 2.0, false},
		{"(1 + 2) * 3", 9.0, false},
		{"10 / (2 + 3)", 2.0, false},
		{"invalid expression", 0.0, true}, // Пример неверного выражения
		{"5 / 0", 0.0, true},              // Деление на ноль
	}

	for _, test := range tests {
		t.Run(test.expr, func(t *testing.T) {
			result, err := Eval(test.expr)
			if (err != nil) != test.shouldErr {
				t.Errorf("Eval(%q) returned error: %v, expected error: %v", test.expr, err != nil, test.shouldErr)
				return
			}
			if result != test.expected {
				t.Errorf("Eval(%q) = %v; want %v", test.expr, result, test.expected)
			}
		})
	}
}
