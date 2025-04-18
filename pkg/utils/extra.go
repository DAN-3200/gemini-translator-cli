package utils

func Ternary[Anyone any](conditional bool, value1 Anyone, value2 Anyone) Anyone {
	if conditional {
		return value1
	}
	return value2
}