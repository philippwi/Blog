package dataHandling

import (
	"Blog/config"
	"io/ioutil"
	"encoding/json"
)

func GetUserList() config.UserList {
	data, _ := ioutil.ReadFile(config.DataDir + "users.json")
	var users []config.User
	json.Unmarshal(data, &users)
	list := config.UserList{Users: users}
	return list
}

func GetBlogEntryList() config.BlogEntryList{
	data, _ := ioutil.ReadFile(config.DataDir + "blogEntries.json")
	var blogEntries []config.BlogEntry
	json.Unmarshal(data, &blogEntries)
	list := config.BlogEntryList{BlogEntries: blogEntries}
	return list
}

func GetCommentList() config.CommentList {
	data, _ := ioutil.ReadFile(config.DataDir + "comments.json")
	var comments []config.Comment
	json.Unmarshal(data, &comments)
	list := config.CommentList{Comments: comments}
	return list
}

//alt
/*func GetUserList() []config.User {
	data, _ := ioutil.ReadFile(config.DataDir + "users.json")
	var users []config.User
	json.Unmarshal(data, &users)
	return users
}

func GetBlogEntryList() []config.BlogEntry {
	data, _ := ioutil.ReadFile(config.DataDir + "blogEntries.json")
	var blogEntries []config.BlogEntry
	json.Unmarshal(data, &blogEntries)
	return blogEntries
}

func GetCommentList() []config.Comment {
	data, _ := ioutil.ReadFile(config.DataDir + "comments.json")
	var comments []config.Comment
	json.Unmarshal(data, &comments)
	return comments
}*/
