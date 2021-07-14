package parrot_trouble_test

import (
	"codingbat/warmup-1/parrot_trouble"
	"testing"
)

func TestParrotTroubleTalkingAtSix(t *testing.T) {
	input := 6
	talking := true
	expected := true
	result := parrot_trouble.ParrotTrouble(talking, input)
	if result != expected {
		t.Errorf("Parrot Talking at Six (%d) failed, expected %v, got %v", input, expected, result)
	} else {
		t.Logf("Parrot Talking at Six (%d) success, expected %v, got %v", input, expected, result)
	}
}

func TestParrotTroubleTalkingAtSeven(t *testing.T) {
	input := 7
	talking := true
	expected := false
	result := parrot_trouble.ParrotTrouble(talking, input)
	if result != expected {
		t.Errorf("Parrot Talking at Seven (%d) failed, expected %v, got %v", input, expected, result)
	} else {
		t.Logf("Parrot Talking at Seven (%d) success, expected %v, got %v", input, expected, result)
	}
}

func TestParrotTroubleTalkingAtTwenty(t *testing.T) {
	input := 20
	talking := true
	expected := false
	result := parrot_trouble.ParrotTrouble(talking, input)
	if result != expected {
		t.Errorf("Parrot Talking at Twenty (%d) failed, expected %v, got %v", input, expected, result)
	} else {
		t.Logf("Parrot Talking at Twenty (%d) success, expected %v, got %v", input, expected, result)
	}
}

func TestParrotTroubleTalkingAtTwentyOne(t *testing.T) {
	input := 21
	talking := true
	expected := true
	result := parrot_trouble.ParrotTrouble(talking, input)
	if result != expected {
		t.Errorf("Parrot Talking at Twenty One (%d) failed, expected %v, got %v", input, expected, result)
	} else {
		t.Logf("Parrot Talking at Twenty One (%d) success, expected %v, got %v", input, expected, result)
	}
}

func TestParrotTroubleNotTalkingAtSix(t *testing.T) {
	input := 6
	talking := false
	expected := false
	result := parrot_trouble.ParrotTrouble(talking, input)
	if result != expected {
		t.Errorf("Parrot Not Talking at Six (%d) failed, expected %v, got %v", input, expected, result)
	} else {
		t.Logf("Parrot Not Talking at Six (%d) success, expected %v, got %v", input, expected, result)
	}
}

func TestParrotTroubleNotTalkingAtSeven(t *testing.T) {
	input := 7
	talking := false
	expected := false
	result := parrot_trouble.ParrotTrouble(talking, input)
	if result != expected {
		t.Errorf("Parrot Not Talking at Seven (%d) failed, expected %v, got %v", input, expected, result)
	} else {
		t.Logf("Parrot Not Talking at Seven (%d) success, expected %v, got %v", input, expected, result)
	}
}

func TestParrotTroubleNotTalkingAtTwenty(t *testing.T) {
	input := 20
	talking := false
	expected := false
	result := parrot_trouble.ParrotTrouble(talking, input)
	if result != expected {
		t.Errorf("Parrot Not Talking at Twenty (%d) failed, expected %v, got %v", input, expected, result)
	} else {
		t.Logf("Parrot Not Talking at Twenty (%d) success, expected %v, got %v", input, expected, result)
	}
}

func TestParrotTroubleNotTalkingAtTwentyOne(t *testing.T) {
	input := 21
	talking := false
	expected := false
	result := parrot_trouble.ParrotTrouble(talking, input)
	if result != expected {
		t.Errorf("Parrot Not Talking at Twenty One (%d) failed, expected %v, got %v", input, expected, result)
	} else {
		t.Logf("Parrot Not Talking at Twenty One (%d) success, expected %v, got %v", input, expected, result)
	}
}
