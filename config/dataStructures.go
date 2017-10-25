package config

//general

/*type UserList struct {
	Users []User
}
type BlogEntryList struct {
	BlogEntries []BlogEntry
}

type CommentList struct {
	Comments []Comment
}*/

type User struct {
	Name     string
	Password string
	ID       int
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
	ID     int
}

//page specific

type HomeData struct {
	CurrentUser string
	BlogEntries []BlogEntry
}

type ViewblogData struct {
	CurrentUser  string
	Blog         BlogEntry
	BlogComments []Comment
}
