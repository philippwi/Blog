//Matrikelnummern: 3229403, 9964427

package dataHandling

import (
	"Blog/config"
	"io/ioutil"
	"encoding/json"
	"testing"
	"Blog/utility"
)

//Testblog mit ID=0 und Autor=TestUser muss in blogEntries.json existieren!
//Testkommentar mit BlogID=0 und Autor=TestUser muss in comments.json existieren!

func TestGetAllUsers(t *testing.T) {
	data, err := ioutil.ReadFile(utility.FixPath(config.DataDir) + "users.json")

	if err != nil{
		utility.HandleError(err)
	}

	var actualUsers []config.User
	json.Unmarshal(data, &actualUsers)

	readUsers := GetAllUsers()

	if len(actualUsers) != len(readUsers) {
		t.Error("Nutzerdaten wurden nicht korrekt gelesen")
	}
}

func TestGetAllBlogEntries(t *testing.T) {
	data, err := ioutil.ReadFile(utility.FixPath(config.DataDir) + "blogEntries.json")

	if err != nil{
		utility.HandleError(err)
	}

	var actualBlogs []config.BlogEntry
	json.Unmarshal(data, &actualBlogs)

	actualBlogs = SortBlogEntries(actualBlogs)

	readBlogs := GetAllBlogEntries()

	if len(actualBlogs) != len(readBlogs) {
		t.Error("Blogdaten wurden nicht korrekt gelesen")
	}
}

func TestGetAllComments(t *testing.T) {
	data, err := ioutil.ReadFile(utility.FixPath(config.DataDir) + "comments.json")

	if err != nil{
		utility.HandleError(err)
	}

	var actualComments []config.Comment
	json.Unmarshal(data, &actualComments)

	actualComments = SortComments(actualComments)

	readComments := GetAllComments()

	if len(actualComments) != len(readComments) {
		t.Error("Blogdaten wurden nicht korrekt gelesen")
	}
}

func TestGetBlog(t *testing.T) {

	blog := GetBlog(0)

	if blog.Author != TestUserName {
		t.Error("Testblog von TestUser wurde nicht gefunden")
	}
}

func TestGetBlogWithComments(t *testing.T) {
	//Testet nur den GetComment-Teil, GetBlog hat eigenen Test
	_, blogComments := GetBlogWithComments(0)

	success := false

	for i,_ := range blogComments{
		if blogComments[i].Author==TestUserName{
			success = true
			break
		}
	}

	if !success{
		t.Error("Testkommentar nicht gefunden")
	}
}
