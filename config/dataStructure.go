package config

type UserList struct {
	Users []User
}

type User struct {
	Name     string
	Password string
}

type BlogEntryList struct {
	BlogEntries []BlogEntry
}

type BlogEntry struct {
	Author  string
	Date    string
	Title   string
	Content string
	ID		int
}

type CommentList struct {
	Comments []Comment
}

type Comment struct {
	Author string
	//Date		int
	Text string
}
