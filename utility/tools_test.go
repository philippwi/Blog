//Matrikelnummern: 3229403, 9964427

package utility

import (
	"testing"
	"os"
	"bytes"
	"log"
	"errors"
)

func TestIsIntInArray(t *testing.T) {

	array1 := []int{1, 2, 3, 4, 5}
	array2 := []int{43, 67, 149, 3, 25, 114, 3217, 12}
	array3 := []int{48, 1, 27, 100, 2}

	success := true

	if !IsIntInArray(4, array1) ||
		!IsIntInArray(25, array2) ||
		IsIntInArray(55, array3) {
		success = false
	}

	if !success {
		t.Error("Fehler: Zahl wurde nicht oder fälschlicherweise im Array gefunden")
	}
}

func TestFixPath(t *testing.T) {
	testPath := "test"
	fixedPath := FixPath(testPath)
	wd, err := os.Getwd()

	if err != nil {
		t.Error("Fehler beim Lesen des Working-Directory")
	}

	switch wd[len(wd)-4:] {
	case "Blog":
		if fixedPath != testPath {
			t.Error("Pfad fehlerhaft - wd: " + wd + " - fixedPath: " + fixedPath)
		}
	default:
		if fixedPath != "../"+testPath {
			t.Error("Pfad fehlerhaft - wd: " + wd + " - fixedPath: " + fixedPath)
		}

	}
}

func TestEncryptCookie(t *testing.T){
	testString := "Test12345;:"
	testStringEnc := "Uftu23456<;"

	if EncryptCookie(testString) != testStringEnc{
		t.Error("Verschlüsselung fehlerhaft")
	}

	//Testfall leerer String
	if EncryptCookie("") != ""{
		t.Error("Verschlüsselung fehlerhaft")
	}
}

func TestDecryptCookie(t *testing.T){
	testString := "Test12345;:"
	testStringEnc := "Uftu23456<;"

	if DecryptCookie(testStringEnc) != testString{
		t.Error("Entschlüsselung fehlerhaft")
	}

	//Testfall leerer String
	if DecryptCookie("") != ""{
		t.Error("Entschlüsselung fehlerhaft")
	}
}

func TestHandleError(t *testing.T) {
	errorMsg1 := "Ein Test-Fehler 1"
	errorMsg2 := "some random error: 1234567890"
	context1 := "Irgenein Kontext"
	context2 := "zufall.dies.das 1234567890"

	var buf bytes.Buffer

	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	HandleError(context1, errors.New(errorMsg1))
	out1:= buf.String()

	if  out1[len(out1)-len(errorMsg1 + "\nKontext: " + context1)-1:] !=
		errorMsg1 + "\nKontext: " + context1+ "\n"{
		t.Error("Falsche Fehlerausgabe (1)")
	}

	HandleError(context2, errors.New(errorMsg2))
	out2:= buf.String()

	if  out2[len(out2)-len(errorMsg2 + "\nKontext: " + context2)-1:] !=
		errorMsg2 + "\nKontext: " + context2+ "\n"{
		t.Error("Falsche Fehlerausgabe (2)")
	}

}

