//Matrikelnummern: 3229403, 9964427

package tests

import (
	"Blog/config"
	"encoding/json"
	"time"
	"os"
	"testing"
	"Blog/dataHandling"
)

//Nutzer mit dem Namen "TestUser" und dem Passwort "123" (nach decrypt) muss in users.json existieren!
//Testblog mit ID=0 muss in blogEntries.json existieren!

var TestUserName = "TestUser"
var TestUserPwSalt = dataHandling.EncryptPassword("123")

func TestSaveUser(t *testing.T) {

	usrName := "abcd123456789"
	usrPw := "12345"

	dataHandling.SaveUser(usrName, usrPw)

	if !dataHandling.UserExists(usrName){
		t.Error("User " + usrName + " wurde nicht gefunden")
	}

	//erstellten user wieder löschen:

	users := dataHandling.GetAllUsers()

	var newUserList []config.User

	for _,u := range users {
		if u.Name == usrName{
			continue
		}
		newUserList = append(newUserList, u)
	}

	file, err := os.Create(config.DataDir + "users.json")
	if err == nil{
		enc := json.NewEncoder(file)
		enc.Encode(newUserList)
	} else{
		t.Error("Test erfolgreich, aber Testnutzer " + usrName + " wurde nicht gelöscht: " + err.Error())
	}
	file.Close()
}

func TestSaveBlogEntry(t *testing.T) {

	author := TestUserName
	blgTitle := "TestTitle"
	blgContent := "TestText"

	dataHandling.SaveBlogEntry(author,blgTitle,blgContent)

	blgEntries := dataHandling.GetAllBlogEntries()

	t.Error("---")


}

func TestSaveComment(t *testing.T) {
	comments := GetAllComments()

	commentData := config.Comment{
		Author: author,
		Date:   time.Now().Format("02.01.2006 um 15:04:05"),
		Text:   text,
		BlogID: blogID,
		ID:     NewCommentID(),
	}

	comments = append(comments, commentData)

	file, err := os.Create(config.DataDir + "comments.json")
	if err == nil {
		enc := json.NewEncoder(file)
		enc.Encode(comments)
	} else {
		panic(err)
	}
	file.Close()
}

func TestChangeUserPassword(t *testing.T) {

	users := GetAllUsers()

	for i, u := range users {
		if u.Name == name {
			users[i].PwSalt = EncryptPassword(password)
			break
		}
	}

	file, err := os.Create(config.DataDir + "users.json")
	if err == nil {
		enc := json.NewEncoder(file)
		enc.Encode(users)
	} else {
		panic(err)
	}
	file.Close()
}

func TestChangeBlogEntry(t *testing.T) {

	blogs := GetAllBlogEntries()

	for i, b := range blogs {
		if b.ID == id {
			blogs[i].Content = content
			break
		}
	}

	file, err := os.Create(config.DataDir + "blogEntries.json")
	if err == nil {
		enc := json.NewEncoder(file)
		enc.Encode(blogs)
	} else {
		panic(err)
	}
	file.Close()
}

func TestDeleteBlogEntry(t *testing.T) {
	blogs := GetAllBlogEntries()

	var newBlogList []config.BlogEntry

	for _, b := range blogs {
		if b.ID == id {
			continue
		}
		newBlogList = append(newBlogList, b)
	}

	file, err := os.Create(config.DataDir + "blogEntries.json")
	if err == nil {
		enc := json.NewEncoder(file)
		enc.Encode(newBlogList)
	} else {
		panic(err)
	}
	file.Close()
}
