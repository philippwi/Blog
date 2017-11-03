//Matrikelnummern: 3229403, 9964427

package main

import (
	"Blog/config"
	"encoding/json"
	"os"
	"testing"
	"Blog/dataHandling"
)

//Nutzer mit dem Namen "TestUser" und dem Passwort "12345" (nach decrypt) muss in users.json existieren!
//Testblog mit ID=0 muss in blogEntries.json existieren!

func TestSaveUser(t *testing.T) {

	usrName := "abcd123456789"
	usrPw := "1234567"

	dataHandling.SaveUser(usrName, usrPw)

	if !dataHandling.UserExists(usrName) {
		t.Error("User " + usrName + " wurde nicht gefunden")
	}

	//erstellten user wieder löschen:

	users := dataHandling.GetAllUsers()

	var newUserList []config.User

	for i, _ := range users {
		if users[i].Name == usrName {
			continue
		}
		newUserList = append(newUserList, users[i])
	}

	file, err := os.Create(config.DataDir + "users.json")
	if err == nil {
		enc := json.NewEncoder(file)
		enc.Encode(newUserList)
	} else {
		t.Error("Test erfolgreich, aber Testnutzer " + usrName + " wurde nicht gelöscht: " + err.Error())
	}
	file.Close()
}

func TestSaveBlogEntry(t *testing.T) {

	author := TestUserName
	blgTitle := "TestTitle"
	blgContent := "TestText"

	dataHandling.SaveBlogEntry(author, blgTitle, blgContent)

	blgEntries := dataHandling.GetAllBlogEntries()

	id := 0

	success := false

	for i, _ := range blgEntries {
		if blgEntries[i].Title == blgTitle && blgEntries[i].Content == blgContent {
			id = blgEntries[i].ID
			success = true
			break
		}
	}
	if !success {
		t.Error("Blog '" + blgTitle + "' sollte gespeichert werden, wurde aber nicht gefunden")
	}

	if id != 0 {
		dataHandling.DeleteBlogEntry(id)
	}

}

func TestSaveComment(t *testing.T) {

	author := TestUserName
	txt := "TestText"
	blgID := 0

	dataHandling.SaveComment(author, txt, blgID)

	comments := dataHandling.GetAllComments()

	success := false

	for i, _ := range comments {
		if comments[i].BlogID == blgID && comments[i].Text == txt {
			success = true
			break
		}
	}

	if success {
		//Kommentar löschen:
		comments = dataHandling.GetAllComments()

		var newComments []config.Comment

		for _, c := range comments {
			if c.BlogID == blgID && c.Text == txt {
				continue
			}
			newComments = append(newComments, c)
		}

		file, err := os.Create(config.DataDir + "comments.json")
		if err == nil {
			enc := json.NewEncoder(file)
			enc.Encode(newComments)
		} else {
			t.Error("Kommentar wurde erfolgreich gespeichert, aber Fehler beim Löschen: " + err.Error())
		}
		file.Close()
	} else {
		t.Error("Kommentar sollte gespeichert werden, wurde aber nicht gefunden")
	}
}

func TestChangeUserPassword(t *testing.T) {

	newPw := "abcde"

	var curPw string

	dataHandling.ChangeUserPassword(TestUserName, newPw)

	users := dataHandling.GetAllUsers()

	success := false

	for i, _ := range users {
		if users[i].Name == TestUserName {
			if dataHandling.DecryptPassword(users[i].PwSalt) == newPw {
				success = true
				users[i].PwSalt = TestUserPwSalt //Passwort wieder zurücksetzen
			} else {
				curPw = dataHandling.DecryptPassword(users[i].PwSalt)
			}
			break
		}
	}

	if !success {
		t.Error("Passwort soll " + newPw + " sein, aber ist: " + curPw)
	}
}

func TestChangeBlogEntry(t *testing.T) {
	testBlogID := 0
	newTxt := "abcde"
	var curTxt string

	dataHandling.ChangeBlogEntry(newTxt, testBlogID)

	blogs := dataHandling.GetAllBlogEntries()

	success := false

	for i, _ := range blogs {
		if blogs[i].ID == testBlogID {
			curTxt = blogs[i].Content
			dataHandling.ChangeBlogEntry(newTxt, testBlogID)
			if blogs[i].Content == newTxt {
				success = true
				blogs[i].Content = curTxt
				break
			}
		}
	}

	if !success {
		t.Error("Bloginhalt wurde nicht korrekt geändert")
	}
}

func TestDeleteBlogEntry(t *testing.T) {
	author := TestUserName
	blgTitle := "TestTitle"
	blgContent := "TestText"

	var blgID int

	success := true

	dataHandling.SaveBlogEntry(author, blgTitle, blgContent)

	blgEntries := dataHandling.GetAllBlogEntries()

	//ID finden
	for i, _ := range blgEntries {
		if blgEntries[i].Title == blgTitle && blgEntries[i].Content == blgContent {
			blgID = blgEntries[i].ID
			break
		}
	}
	dataHandling.DeleteBlogEntry(blgID)

	blgEntries = dataHandling.GetAllBlogEntries()

	for i, _ := range blgEntries {
		if blgEntries[i].ID == blgID {
			success = false
			break
		}
	}

	if !success{
		t.Error("Blog wurde nicht korrekt gelöscht")
	}
}
