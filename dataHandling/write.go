package dataHandling

import (
	"Blog/config"
	"encoding/json"
	"time"
	"os"
)

func SaveUser(name, password string){
	users := GetUserList().Users
	userData := config.User{
		Name: name,
		Password: password,
	}
	users = append(users, userData)

	file, err := os.Create(config.DataDir + "users.json")
	if err == nil{
		enc := json.NewEncoder(file)
		enc.Encode(users)
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

func SaveComment(author, text string){
	comments := GetCommentList().Comments
	commentData := config.Comment{
		Author: author,
		Text: text,
	}
	comments = append(comments, commentData)

	file, err := os.Create(config.DataDir + "comments.json")
	if err == nil{
		enc := json.NewEncoder(file)
		enc.Encode(comments)
	} else{
		panic(err)
	}
	file.Close()
}
