//Matrikelnummern: 3229403, 9964427

package dataHandling

import (
	"Blog/config"
	"io/ioutil"
	"encoding/json"
)

func GetAllUsers() []config.User {
	data, _ := ioutil.ReadFile(config.DataDir + "users.json")
	var users []config.User
	json.Unmarshal(data, &users)
	return users
}

func GetAllBlogEntries() []config.BlogEntry{
	data, _ := ioutil.ReadFile(config.DataDir + "blogEntries.json")
	var blogEntries []config.BlogEntry
	json.Unmarshal(data, &blogEntries)
	return SortBlogEntries(blogEntries)
}

func GetAllComments() []config.Comment {
	data, _ := ioutil.ReadFile(config.DataDir + "comments.json")
	var comments []config.Comment
	json.Unmarshal(data, &comments)
	return SortComments(comments)
}

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
