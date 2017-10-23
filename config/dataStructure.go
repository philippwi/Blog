package config

type User struct {
	Name		string
	Password	string
}

type BlogEntry struct {
	Author		string
	Date		string
	Title		string
	Content		string
}

type Comment struct {
	Author		string
	//Date		int
	Text		string
}

