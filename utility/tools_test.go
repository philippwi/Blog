//Matrikelnummern: 3229403, 9964427

package utility

import (
	"testing"
	"os"
)

func TestIsIntInArray(t *testing.T){

	array1 := []int{1,2,3,4,5}
	array2 := []int{43,67,149,3,25,114,3217,12}

	success := true

	if !IsIntInArray(4, array1) ||	!IsIntInArray(25, array2){
		success = false
	}

	if !success {
		t.Error("Expected: true, got: false")
	}
}

func TestFixPath(t *testing.T) {
	testPath := "test"
	fixedPath := FixPath(testPath)
	wd, err := os.Getwd()

	if err != nil{
		panic(err)
	}

	switch wd[len(wd)-4:]{
	case "Blog":
		if fixedPath != testPath{
			t.Error("Pfad fehlerhaft - wd: " +wd+ " - fixedPath: " + fixedPath)
		}
	default:
		if fixedPath != "../"+testPath{
			t.Error("Pfad fehlerhaft - wd: " +wd+ " - fixedPath: " + fixedPath)
		}

	}
}
