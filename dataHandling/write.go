//Matrikelnummern: 3229403, 9964427

package dataHandling

import (
	"Blog/config"
	"encoding/json"
	"time"
	"os"
	"Blog/utility"
)

//speichert Nutzer in .json Datei
func SaveUser(name, password string){
	users := GetAllUsers()

	userData := config.User{
		Name:   name,
		PwSalt: EncryptPassword(password),
		//ID:     NewUserID(),
	}
	users = append(users, userData)

	file, err := os.Create(utility.FixPath(config.DataDir) + "users.json")
	if err == nil{
		enc := json.NewEncoder(file)
		enc.Encode(users)
	} else{
		panic(err)
	}
	file.Close()
}

//speichert Blogeintrag in .json Datei
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

	file, err := os.Create(utility.FixPath(config.DataDir) + "blogEntries.json")
	if err == nil{
		enc := json.NewEncoder(file)
		enc.Encode(entries)
	} else{
		panic(err)
	}
	file.Close()
}

//speichert Kommentar in .json Datei
func SaveComment(author, text string, blogID int){
	comments := GetAllComments()

	commentData := config.Comment{
		Author: author,
		Date: time.Now().Format("02.01.2006 um 15:04:05"),
		Text: text,
		BlogID: blogID,
//		ID: NewCommentID(),
	}

	comments = append(comments, commentData)

	file, err := os.Create(utility.FixPath(config.DataDir) + "comments.json")
	if err == nil{
		enc := json.NewEncoder(file)
		enc.Encode(comments)
	} else{
		panic(err)
	}
	file.Close()
}

//ändert Passwort eines bestimmten Nutzer
func ChangeUserPassword(name, password string){

	users := GetAllUsers()

	for i,u := range users{
		if u.Name == name{
			users[i].PwSalt = EncryptPassword(password)
			break
		}
	}

	file, err := os.Create(utility.FixPath(config.DataDir) + "users.json")
	if err == nil{
		enc := json.NewEncoder(file)
		enc.Encode(users)
	} else{
		panic(err)
	}
	file.Close()
}

//ändert Inhalt eines bestimmten Blogeintrages
func ChangeBlogEntry(content string, id int){

	blogs := GetAllBlogEntries()

	for i,b := range blogs{
		if b.ID == id{
			blogs[i].Content = content
			break
		}
	}

	file, err := os.Create(utility.FixPath(config.DataDir) + "blogEntries.json")
	if err == nil{
		enc := json.NewEncoder(file)
		enc.Encode(blogs)
	} else{
		panic(err)
	}
	file.Close()
}

//löscht bestimmten Blogeintrag aus .json Datei
func DeleteBlogEntry(id int){
	blogs := GetAllBlogEntries()

	var newBlogList []config.BlogEntry

	for i, _ := range blogs{
		if blogs[i].ID == id{
			continue
		}
		newBlogList = append(newBlogList, blogs[i])
	}

	file, err := os.Create(utility.FixPath(config.DataDir) + "blogEntries.json")
	if err == nil{
		enc := json.NewEncoder(file)
		enc.Encode(newBlogList)
	} else{
		panic(err)
	}
	file.Close()
}