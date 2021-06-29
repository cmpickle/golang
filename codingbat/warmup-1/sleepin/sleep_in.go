package sleepin

func SleepIn(weekday bool, vacation bool) bool {
	return !weekday || vacation
}
