//Matrikelnummern: 3229403, 9964427

package utility

func IsIntInArray(number int, list []int) bool{
	for _, x := range list {
		if x == number {
			return true
		}
	}
	return false
}
