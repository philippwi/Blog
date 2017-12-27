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
	data, err := ioutil.ReadFile(utility.FixPath(config.DataDir) + "users.json")

	if err != nil{
		utility.HandleError("read.GetAllUsers -> users auslesen", err)
	}

	var users []config.User
	json.Unmarshal(data, &users)
	return users
}

//liefert alle Blogeinträge aus entsprechender .json Datei
func GetAllBlogEntries() []config.BlogEntry{
	data, err := ioutil.ReadFile(utility.FixPath(config.DataDir) + "blogEntries.json")

	if err != nil{
		utility.HandleError("read.GetAllBlogEntries -> Blogs auslesen", err)
	}

	var blogEntries []config.BlogEntry
	json.Unmarshal(data, &blogEntries)
	return SortBlogEntries(blogEntries)
}

//liefert alle Kommentare aus entsprechender .json Datei
func GetAllComments() []config.Comment {
	data, err := ioutil.ReadFile(utility.FixPath(config.DataDir) + "comments.json")

	if err != nil{
		utility.HandleError("read.GetAllComments -> Kommentare auslesen", err)
	}

	var comments []config.Comment
	json.Unmarshal(data, &comments)
	return SortComments(comments)
}

//liefert Blog mit bestimmter ID
func GetBlog(blogID int) (blog config.BlogEntry) {

	allBlogEntries := GetAllBlogEntries()

	for i, _ := range allBlogEntries {
		if allBlogEntries[i].ID == blogID {
			blog = allBlogEntries[i]
			break
		}
	}
	return blog
}

//liefert Blog mit bestimmter ID und zugehörige Kommentare
func GetBlogWithComments(blogID int) (blog config.BlogEntry, blogComments []config.Comment) {

	allBlogEntries := GetAllBlogEntries()
	allComments := GetAllComments()

	for i, _ := range allBlogEntries {
		if allBlogEntries[i].ID == blogID {
			blog = allBlogEntries[i]
			break
		}
	}

	for i, _ := range allComments {
		if allComments[i].BlogID == blogID {
			blogComments = append(blogComments, allComments[i])
		}
	}

	return blog, blogComments
}
