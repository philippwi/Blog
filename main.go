package main

import (
	"Blog/dataHandling"
	"Blog/config"
)

func main(){
	//dataHandling.SaveUser("ich", "pw")

	var users []config.User

	users = append(users, config.User{
		Name: "abc",
		Password: "1234",
	})
	dataHandling.SaveUser("abc","123")

	dataHandling.GetUsers()
}
