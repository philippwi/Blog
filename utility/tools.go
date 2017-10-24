package utility

func IsIntInArray(number int, list []int) bool{
	for _, x := range list {
		if x == number {
			return true
		}
	}
	return false
}
