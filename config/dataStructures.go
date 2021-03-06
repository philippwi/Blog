//Matrikelnummern: 3229403, 9964427

package config

//general

type User struct {
	Name   string
	PwSalt string
}

type BlogEntry struct {
	Author  string
	Date    string
	Title   string
	Content string
	ID      int
}

type Comment struct {
	Author string
	Date   string
	Text   string
	BlogID int
}

//page specific

type HomeData struct {
	CurrentUser string
	BlogEntries []BlogEntry
}

type ViewblogData struct {
	CurrentUser  string
	NickName	 string
	Blog         BlogEntry
	BlogComments []Comment
}

type ChangeblogData struct {
	CurrentUser string
	BlogContent string
}

type Message struct {
	MsgText  string
	Redirect string
}
