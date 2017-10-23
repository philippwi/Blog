package main

import (
	"Blog/dataHandling"
)

func main(){
	dataHandling.SaveUser("TestUser", "123")
	dataHandling.SaveBlogEntry("TestUser", "TestBlog", "Test test test")
	dataHandling.SaveComment("TestUser", "Testi test")
}
