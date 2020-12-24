package utils

func InInt(value int, numbers ...int) bool {
	for _, n := range numbers {
		if n == value {
			return true
		}
	}
	return false
}

func InString(value string, strings ...string) bool {
	for _, s := range strings {
		if s == value {
			return true
		}
	}
	return false
}
