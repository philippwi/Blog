//Matrikelnummern: 3229403, 9964427

package tests

import (
	"Blog/config"
	"io/ioutil"
	"encoding/json"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	data, _ := ioutil.ReadFile(config.DataDir + "users.json")
	var users []config.User
	json.Unmarshal(data, &users)
	return users
}

func TestGetAllBlogEntries(t *testing.T){
	data, _ := ioutil.ReadFile(config.DataDir + "blogEntries.json")
	var blogEntries []config.BlogEntry
	json.Unmarshal(data, &blogEntries)
	return SortBlogEntries(blogEntries)
}

func TestGetAllComments(t *testing.T){
	data, _ := ioutil.ReadFile(config.DataDir + "comments.json")
	var comments []config.Comment
	json.Unmarshal(data, &comments)
	return SortComments(comments)
}

