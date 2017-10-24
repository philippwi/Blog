package dataHandling

import (
	"Blog/config"
	"encoding/json"
	"time"
	"os"
)

func SaveUser(name, password string){
	userList := GetUserList()

	userData := config.User{
		Name: name,
		Password: password,
		ID: NewUserID(userList),
	}
	userList.Users = append(userList.Users, userData)

	file, err := os.Create(config.DataDir + "users.json")
	if err == nil{
		enc := json.NewEncoder(file)
		enc.Encode(userList.Users)
	} else{
		panic(err)
	}
	file.Close()
}

func SaveBlogEntry(author, title, content string){
	entryList := GetBlogEntryList()

	entryData := config.BlogEntry{
		Author: author,
		Date: time.Now().Format("20060102"),
		Title: title,
		Content: content,
		ID: NewBlogID(entryList),
	}

	entryList.BlogEntries = append(entryList.BlogEntries, entryData)

	file, err := os.Create(config.DataDir + "blogEntries.json")
	if err == nil{
		enc := json.NewEncoder(file)
		enc.Encode(entryList.BlogEntries)
	} else{
		panic(err)
	}
	file.Close()
}

func SaveComment(author, text string, blogID int){
	commentList := GetCommentList()

	commentData := config.Comment{
		Author: author,
		Date: time.Now().Format("20060102"),
		Text: text,
		BlogID: blogID,
		ID: NewCommentID(commentList),
	}

	commentList.Comments = append(commentList.Comments, commentData)

	file, err := os.Create(config.DataDir + "comments.json")
	if err == nil{
		enc := json.NewEncoder(file)
		enc.Encode(commentList.Comments)
	} else{
		panic(err)
	}
	file.Close()
}
