//Matrikelnummern: 3229403, 9964427

package tests

import (
	"testing"
	"Blog/utility"
)

func TestIsIntInArray(t *testing.T){

	array1 := []int{1,2,3,4,5}
	array2 := []int{43,67,149,3,25,114,3217,12}

	success := true

	if !utility.IsIntInArray(4, array1) ||	!utility.IsIntInArray(25, array2){
		success = false
	}

	if !success {
		t.Error("Expected: true, got: false")
	}

}
