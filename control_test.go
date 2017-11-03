//Matrikelnummern: 3229403, 9964427

package main

import (
	"testing"
	"Blog/dataHandling"
)

//Nutzer mit dem Namen "TestUser" und dem Passwort "12345" (nach decrypt) muss in users.json existieren!

var TestUserName = "TestUser"
var TestUserPwSalt = dataHandling.EncryptPassword("12345")

func TestUserExists(t *testing.T) {
	if !dataHandling.UserExists(TestUserName) {
		t.Error("Testnutzer angeblich nicht existent")
	}
}

func TestPasswordCorrect(t *testing.T) {
	if !dataHandling.PasswordCorrect(TestUserName, dataHandling.DecryptPassword(TestUserPwSalt)) ||
		dataHandling.PasswordCorrect(TestUserName, "abc") {
		t.Error("Passwortprüfung fehlerhaft")
	}
}

func TestEncryptPassword(t *testing.T) {
	pw := "abcd12345"
	encPw := dataHandling.EncryptPassword(pw)

	if pw == encPw {
		t.Error("Verschlüsselung fehlgeschlagen")
	}

}

func TestDecryptPassword(t *testing.T) {
	pw := "abcd12345"
	encPw := dataHandling.EncryptPassword(pw)

	if dataHandling.DecryptPassword(encPw) != pw {
		t.Error("Entschlüsselung fehlgeschlagen")
	}
}

func TestNewUserID(t *testing.T) {

}

func TestNewBlogID(t *testing.T) {
	newID := dataHandling.NewBlogID()

	blogs := dataHandling.GetAllBlogEntries()

	success := true

	for i, _ := range blogs {
		if blogs[i].ID == newID {
			success = false
			break
		}
	}

	if !success{
		t.Error("BlogID ist bereits vergeben")
	}
}
/*
func TestSortBlogEntries(t *testing.T) {

	sort.Slice(entries, func(i, j int) bool {
		date1, _ := time.Parse("02.01.2006 um 15:04:05", entries[i].Date)
		date2, _ := time.Parse("02.01.2006 um 15:04:05", entries[j].Date)

		date1int, _ := strconv.Atoi(date1.Format("20060102150405"))
		date2int, _ := strconv.Atoi(date2.Format("20060102150405"))
		return date1int > date2int
	})

	return entries
}

func TestSortComments(t *testing.T) {

	sort.Slice(comments, func(i, j int) bool {
		date1, _ := time.Parse("02.01.2006 um 15:04:05", comments[i].Date)
		date2, _ := time.Parse("02.01.2006 um 15:04:05", comments[j].Date)

		date1int, _ := strconv.Atoi(date1.Format("20060102150405"))
		date2int, _ := strconv.Atoi(date2.Format("20060102150405"))
		return date1int > date2int
	})

	return comments
}
*/