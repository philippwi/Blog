package dataHandling

import (
	"Blog/config"
	"encoding/json"
	"time"
	"os"
)

func SaveUser(name, password string){
	users := GetAllUsers()

	userData := config.User{
		Name: name,
		Password: password,
		ID: NewUserID(),
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
	entries := GetAllBlogEntries()

	entryData := config.BlogEntry{
		Author: author,
		Date: time.Now().Format("20060102"),
		Title: title,
		Content: content,
		ID: NewBlogID(),
	}

	entries = append(entries, entryData)

	file, err := os.Create(config.DataDir + "blogEntries.json")
	if err == nil{
		enc := json.NewEncoder(file)
		enc.Encode(entries)
	} else{
		panic(err)
	}
	file.Close()
}

func SaveComment(author, text string, blogID int){
	comments := GetAllComments()

	commentData := config.Comment{
		Author: author,
		Date: time.Now().Format("20060102"),
		Text: text,
		BlogID: blogID,
		ID: NewCommentID(),
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
