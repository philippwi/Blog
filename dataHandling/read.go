//Matrikelnummern: 3229403, 9964427

package dataHandling

import (
	"Blog/config"
	"io/ioutil"
	"encoding/json"
	"Blog/utility"
)

//liefert alle Nutzer aus entsprechender .json Datei
func GetAllUsers() []config.User {
	data, _ := ioutil.ReadFile(utility.FixPath(config.DataDir) + "users.json")
	var users []config.User
	json.Unmarshal(data, &users)
	return users
}

//liefert alle Blogeinträge aus entsprechender .json Datei
func GetAllBlogEntries() []config.BlogEntry{
	data, _ := ioutil.ReadFile(utility.FixPath(config.DataDir) + "blogEntries.json")
	var blogEntries []config.BlogEntry
	json.Unmarshal(data, &blogEntries)
	return SortBlogEntries(blogEntries)
}

//liefert alle Kommentare aus entsprechender .json Datei
func GetAllComments() []config.Comment {
	data, _ := ioutil.ReadFile(utility.FixPath(config.DataDir) + "comments.json")
	var comments []config.Comment
	json.Unmarshal(data, &comments)
	return SortComments(comments)
}

//liefert Blog mit bestimmter ID
func GetBlog(blogID int) (blog config.BlogEntry) {

	allBlogEntries := GetAllBlogEntries()

	for _, b := range allBlogEntries {
		if b.ID == blogID {
			blog = b
			break
		}
	}
	return blog
}

//liefert Blog mit bestimmter ID und zugehörige Kommentare
func GetBlogWithComments(blogID int) (blog config.BlogEntry, blogComments []config.Comment) {

	allBlogEntries := GetAllBlogEntries()
	allComments := GetAllComments()

	for _, b := range allBlogEntries {
		if b.ID == blogID {
			blog = b
			break
		}
	}

	for _, c := range allComments {
		if c.BlogID == blogID {
			blogComments = append(blogComments, c)
		}
	}

	return blog, blogComments
}
