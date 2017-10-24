package config

//general

type UserList struct {
	Users []User
}

type User struct {
	Name     string
	Password string
	ID       int
}

type BlogEntryList struct {
	BlogEntries []BlogEntry
}

type BlogEntry struct {
	Author  string
	Date    string
	Title   string
	Content string
	ID      int
}

type CommentList struct {
	Comments []Comment
}

type Comment struct {
	Author string
	Date   string
	Text   string
	BlogID int
	ID     int
}

//page specific

type ViewblogData struct {
	Blog         BlogEntry
	BlogComments []Comment
}
