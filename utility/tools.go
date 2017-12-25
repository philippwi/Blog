//Matrikelnummern: 3229403, 9964427

package utility

import (
	"os"
	"log"
)

//testet ob eine Zahl in einer Zahlenreihe vorhanden ist
func IsIntInArray(number int, list []int) bool{
	for _, x := range list {
		if x == number {
			return true
		}
	}
	return false
}

//korrigiert Pfade
//benötigt für korrekten Testablauf
func FixPath(path string) string{
	wd, err := os.Getwd()

	if err != nil{
		HandleError(err)
	}

	if wd[len(wd)-4:] != "Blog"{
		return "../"+path
	}else{
		return path
	}
}

func EncryptCookie(txt string) string{

	btSlc := []byte(txt)

	for i,_ := range btSlc{
		btSlc[i] += 1
	}

	return string(btSlc)
}

func DecryptCookie(txt string) string{

	btSlc := []byte(txt)

	for i,_ := range btSlc{
		btSlc[i] -= 1
	}
	return string(btSlc)
}

func HandleError(e error){
	log.Println("FEHLER: " + e.Error())
}
