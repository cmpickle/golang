package monkey_trouble_test

import (
	"codingbat/warmup-1/monkey_trouble"
	"testing"
)

func TestMonkeyTroubleBothSmiling(t *testing.T) {
	result := monkey_trouble.MonkeyTrouble(true, true)
	if result != true {
		t.Errorf("Monkey Trouble Both Smiling error, got %v expected true", result)
	} else {
		t.Logf("Monkey Trouble Both Smiling success, got %v expected true", result)
	}
}

func TestMonkeyTroubleASmiling(t *testing.T) {
	result := monkey_trouble.MonkeyTrouble(true, false)
	if result != false {
		t.Errorf("Monkey Trouble A Smiling error, got %v expected true", result)
	} else {
		t.Logf("Monkey Trouble A Smiling success, got %v expected true", result)
	}
}

func TestMonkeyTroubleBSmiling(t *testing.T) {
	result := monkey_trouble.MonkeyTrouble(false, true)
	if result != false {
		t.Errorf("Monkey Trouble B Smiling error, got %v expected true", result)
	} else {
		t.Logf("Monkey Trouble B Smiling success, got %v expected true", result)
	}
}

func TestMonkeyTroubleNeitherSmiling(t *testing.T) {
	result := monkey_trouble.MonkeyTrouble(false, false)
	if result != true {
		t.Errorf("Monkey Trouble Neither Smiling error, got %v expected true", result)
	} else {
		t.Logf("Monkey Trouble Neither Smiling success, got %v expected true", result)
	}
}
