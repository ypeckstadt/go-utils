package utils

func InInt(value int, numbers ...int) bool {
	for _, n := range numbers {
		if n == value {
			return true
		}
	}
	return false
}
