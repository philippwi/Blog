//Matrikelnummern: 3229403, 9964427

package main

import (
	"Blog/config"
	"io/ioutil"
	"encoding/json"
	"testing"
	"Blog/dataHandling"
)

//Testblog mit ID=0 und Autor=TestUser muss in blogEntries.json existieren!
//Testkommentar mit BlogID=0 und Autor=TestUser muss in comments.json existieren!

func TestGetAllUsers(t *testing.T) {
	data, _ := ioutil.ReadFile(config.DataDir + "users.json")
	var actualUsers []config.User
	json.Unmarshal(data, &actualUsers)

	readUsers := dataHandling.GetAllUsers()

	if len(actualUsers) != len(readUsers) {
		t.Error("Nutzerdaten wurden nicht korrekt gelesen")
	}
}

func TestGetAllBlogEntries(t *testing.T) {
	data, _ := ioutil.ReadFile(config.DataDir + "blogEntries.json")
	var actualBlogs []config.BlogEntry
	json.Unmarshal(data, &actualBlogs)

	actualBlogs = dataHandling.SortBlogEntries(actualBlogs)

	readBlogs := dataHandling.GetAllBlogEntries()

	if len(actualBlogs) != len(readBlogs) {
		t.Error("Blogdaten wurden nicht korrekt gelesen")
	}
}

func TestGetAllComments(t *testing.T) {
	data, _ := ioutil.ReadFile(config.DataDir + "comments.json")
	var actualComments []config.Comment
	json.Unmarshal(data, &actualComments)

	actualComments = dataHandling.SortComments(actualComments)

	readComments := dataHandling.GetAllComments()

	if len(actualComments) != len(readComments) {
		t.Error("Blogdaten wurden nicht korrekt gelesen")
	}
}

func TestGetBlog(t *testing.T) {

	blog := dataHandling.GetBlog(0)

	if blog.Author != TestUserName {
		t.Error("Testblog von TestUser wurde nicht gefunden")
	}
}

func TestGetBlogWithComments(t *testing.T) {
	//Testet nur den GetComment-Teil, GetBlog hat eigenen Test
	_, blogComments := dataHandling.GetBlogWithComments(0)

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
