package dataHandling

import (
	"Blog/config"
	"Blog/utility"
)

func UserExists(name string) bool {
	existingUsers := GetUserList()

	for _, user := range existingUsers.Users {
		if user.Name == name {
			return true
		}
	}
	return false
}

func PasswordCorrect(name, password string) bool {
	existingUsers := GetUserList()

	for _, user := range existingUsers.Users {
		if user.Name == name && user.Password == password {
			return true
		}
	}
	return false
}

func NewUserID(userList config.UserList) int{
	x := 1

	var assignedIDs []int

	for _,u:= range userList.Users{
		assignedIDs = append(assignedIDs, u.ID)
	}

	for utility.IsIntInArray(x, assignedIDs){
		x++
	}

	return x
}

func NewBlogID(blogEntryList config.BlogEntryList) int{
	x := 1

	var assignedIDs []int

	for _,b:= range blogEntryList.BlogEntries{
		assignedIDs = append(assignedIDs, b.ID)
	}

	for utility.IsIntInArray(x, assignedIDs){
		x++
	}

	return x
}

func NewCommentID(commentList config.CommentList) int{
	x := 1

	var assignedIDs []int

	for _,c:= range commentList.Comments{
		assignedIDs = append(assignedIDs, c.ID)
	}

	for utility.IsIntInArray(x, assignedIDs){
		x++
	}

	return x
}


