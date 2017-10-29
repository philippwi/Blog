//Matrikelnummern: 3229403, 9964427

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
		Name:   name,
		PwSalt: EncryptPassword(password),
		ID:     NewUserID(),
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
		Date: time.Now().Format("02.01.2006 um 15:04:05"),
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
		Date: time.Now().Format("02.01.2006 um 15:04:05"),
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

func ChangeUserPassword(name, password string){

	users := GetAllUsers()

	for i,u := range users{
		if u.Name == name{
			users[i].PwSalt = EncryptPassword(password)
			break
		}
	}

	file, err := os.Create(config.DataDir + "users.json")
	if err == nil{
		enc := json.NewEncoder(file)
		enc.Encode(users)
	} else{
		panic(err)
	}
	file.Close()
}

func ChangeBlogEntry(content string, id int){

	blogs := GetAllBlogEntries()

	for i,b := range blogs{
		if b.ID == id{
			blogs[i].Content = content
			break
		}
	}

	file, err := os.Create(config.DataDir + "blogEntries.json")
	if err == nil{
		enc := json.NewEncoder(file)
		enc.Encode(blogs)
	} else{
		panic(err)
	}
	file.Close()
}

func DeleteBlogEntry(id int){
	blogs := GetAllBlogEntries()

	var newBlogList []config.BlogEntry

	for _,b := range blogs{
		if b.ID == id{
			continue
		}
		newBlogList = append(newBlogList, b)
	}

	file, err := os.Create(config.DataDir + "blogEntries.json")
	if err == nil{
		enc := json.NewEncoder(file)
		enc.Encode(newBlogList)
	} else{
		panic(err)
	}
	file.Close()
}