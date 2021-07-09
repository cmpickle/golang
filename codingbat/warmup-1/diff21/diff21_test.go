package diff21_test

import (
	"codingbat/warmup-1/diff21"
	"testing"
)

func TestDiff21input1(t *testing.T) {
	input := 1
	expected := 20
	result := diff21.Diff21(input)
	if result != expected {
		t.Errorf("Diff21 (%d) failed, expected %d, got %v", input, expected, result)
	} else {
		t.Logf("Diff21 (%d) success, expected %d, got %v", input, expected, result)
	}
}

func TestDiff21input0(t *testing.T) {
	input := 0
	expected := 21
	result := diff21.Diff21(input)
	if result != expected {
		t.Errorf("Diff21 (%d) failed, expected %d, got %v", input, expected, result)
	} else {
		t.Logf("Diff21 (%d) success, expected %d, got %v", input, expected, result)
	}
}

func TestDiff21inputMinus1(t *testing.T) {
	input := -1
	expected := 22
	result := diff21.Diff21(input)
	if result != expected {
		t.Errorf("Diff21 (%d) failed, expected %d, got %v", input, expected, result)
	} else {
		t.Logf("Diff21 (%d) success, expected %d, got %v", input, expected, result)
	}
}

func TestDiff21input22(t *testing.T) {
	input := 22
	expected := 2
	result := diff21.Diff21(input)
	if result != expected {
		t.Errorf("Diff21 (%d) failed, expected %d, got %v", input, expected, result)
	} else {
		t.Logf("Diff21 (%d) success, expected %d, got %v", input, expected, result)
	}
}

func TestDiff21inputMinus22(t *testing.T) {
	input := -22
	expected := 43
	result := diff21.Diff21(input)
	if result != expected {
		t.Errorf("Diff21 (%d) failed, expected %d, got %v", input, expected, result)
	} else {
		t.Logf("Diff21 (%d) success, expected %d, got %v", input, expected, result)
	}
}
