package sum

import "testing"

func sumOfTwoIntegers(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}

func sumOfTwoFloats(t *testing.T) {
	result := Add(2.5, 3.5)
	if result != 6.0 {
		t.Errorf("Expected 6.0, got %v", result)
	}
}

func sumOfNegativeNumbers(t *testing.T) {
	result := Add(-2, -3)
	if result != -5 {
		t.Errorf("Expected -5, got %v", result)
	}
}

func sumOfMixedSignNumbers(t *testing.T) {
	result := Add(-2, 3)
	if result != 1 {
		t.Errorf("Expected 1, got %v", result)
	}
}

func sumOfZeroAndNumber(t *testing.T) {
	result := Add(0, 5)
	if result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}

func sumOfTwoZeros(t *testing.T) {
	result := Add(0, 0)
	if result != 0 {
		t.Errorf("Expected 0, got %v", result)
	}
}
