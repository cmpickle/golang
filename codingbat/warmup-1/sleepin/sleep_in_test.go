package sleepin_test

import (
	"codingbat/warmup-1/sleepin"
	"testing"
)

const (
	weekday     = true
	weekend     = false
	vacation    = true
	notVacation = false
	willSleepIn = true
	wontSleepIn = false
)

func TestSleepInWeekdayVacation(t *testing.T) {
	result := sleepin.SleepIn(weekday, vacation)
	if result != willSleepIn {
		t.Errorf("SleepIn on Weekday Vacation failed, expected true, got %v", result)
	} else {
		t.Logf("SleepIn on Weekday Vacation success, expected true, got %v", result)
	}
}

func TestSleepInWeekdayNotVacation(t *testing.T) {
	result := sleepin.SleepIn(weekday, notVacation)
	if result == willSleepIn {
		t.Errorf("SleepIn on Weekday not Vacation failed, expected false, got %v", result)
	} else {
		t.Logf("SleepIn on Weekday not Vacation success, expected false, got %v", result)
	}
}

func TestSleepInWeekendVacation(t *testing.T) {
	result := sleepin.SleepIn(weekend, vacation)
	if result != willSleepIn {
		t.Errorf("SleepIn on Weekdend Vacation failed, expected true, got %v", result)
	} else {
		t.Logf("SleepIn on Weekdend Vacation success, expected true, got %v", result)
	}
}

func TestSleepInWeekendNotVacation(t *testing.T) {
	result := sleepin.SleepIn(weekend, notVacation)
	if result != willSleepIn {
		t.Errorf("SleepIn on Weekend not Vacation failed, expected true, got %v", result)
	} else {
		t.Logf("SleepIn on Weekend not Vacation success, expected true, got %v", result)
	}
}
