package sum_double_test

import (
	"codingbat/warmup-1/sum_double"
	"testing"
)

func TestSumDouble1_1(t *testing.T) {
	result := sum_double.SumDouble(1, 1)
	if result != 4 {
		t.Errorf("Sum double 1 & 1 failed, expected 4, got %v", result)
	} else {
		t.Logf("Sum double 1 & 1 success, expected 4, got %v", result)
	}
}

func TestSumDouble1_2(t *testing.T) {
	result := sum_double.SumDouble(1, 2)
	if result != 3 {
		t.Errorf("Sum double 1 & 2 failed, expected 3, got %v", result)
	} else {
		t.Logf("Sum double 1 & 2 success, expected 3, got %v", result)
	}
}

func TestSumDouble2_2(t *testing.T) {
	result := sum_double.SumDouble(2, 2)
	if result != 8 {
		t.Errorf("Sum double 2 & 2 failed, expected 8, got %v", result)
	} else {
		t.Logf("Sum double 2 & 2 success, expected 8, got %v", result)
	}
}
