package parrot_trouble

func ParrotTrouble(talking bool, hour int) bool {
	return talking && (hour < 7 || hour > 20)
}
