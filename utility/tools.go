//Matrikelnummern: 3229403, 9964427

package utility

import "os"

func IsIntInArray(number int, list []int) bool{
	for _, x := range list {
		if x == number {
			return true
		}
	}
	return false
}

func FixPath(path string) string{
	wd,_ := os.Getwd()

	if wd[len(wd)-4:] != "Blog"{
		return "../"+path
	}else{
		return path
	}
}