//Matrikelnummern: 3229403, 9964427

package dataHandling

import (
	"Blog/utility"
	"Blog/config"
	"sort"
	"time"
	"strconv"
)

func UserExists(name string) bool {
	existingUsers := GetAllUsers()

	for _, user := range existingUsers {
		if user.Name == name {
			return true
		}
	}
	return false
}

func PasswordCorrect(name, password string) bool {
	existingUsers := GetAllUsers()

	for _, user := range existingUsers {
		if name == user.Name && password == DecryptPassword(user.PwSalt) {
			return true
		}
	}
	return false
}

func EncryptPassword(pw string) string {
	return pw
}

func DecryptPassword(code string) string {
	return code
}

func GetBlog(blogID int) (blog config.BlogEntry) {

	allBlogEntries := GetAllBlogEntries()

	for _, b := range allBlogEntries {
		if b.ID == blogID {
			blog = b
			break
		}
	}
	return blog
}

func GetBlogWithComments(blogID int) (blog config.BlogEntry, blogComments []config.Comment) {

	allBlogEntries := GetAllBlogEntries()
	allComments := GetAllComments()

	for _, b := range allBlogEntries {
		if b.ID == blogID {
			blog = b
			break
		}
	}

	for _, c := range allComments {
		if c.BlogID == blogID {
			blogComments = append(blogComments, c)
		}
	}

	return blog, blogComments
}

/*func NewUserID() int {
	users := GetAllUsers()
	x := 1

	var assignedIDs []int

	for _, u := range users {
		assignedIDs = append(assignedIDs, u.ID)
	}

	for utility.IsIntInArray(x, assignedIDs) {
		x++
	}

	return x
}*/

func NewBlogID() int {
	blogEntries := GetAllBlogEntries()
	x := 1

	var assignedIDs []int

	for _, b := range blogEntries {
		assignedIDs = append(assignedIDs, b.ID)
	}

	for utility.IsIntInArray(x, assignedIDs) {
		x++
	}

	return x
}

func NewCommentID() int {
	commentList := GetAllComments()

	x := 1

	var assignedIDs []int

	for _, c := range commentList {
		assignedIDs = append(assignedIDs, c.ID)
	}

	for utility.IsIntInArray(x, assignedIDs) {
		x++
	}

	return x
}

func SortBlogEntries(entries []config.BlogEntry) []config.BlogEntry {

	sort.Slice(entries, func(i, j int) bool {
		date1, _ := time.Parse("02.01.2006 um 15:04:05", entries[i].Date)
		date2, _ := time.Parse("02.01.2006 um 15:04:05", entries[j].Date)

		date1int, _ := strconv.Atoi(date1.Format("20060102150405"))
		date2int, _ := strconv.Atoi(date2.Format("20060102150405"))
		return date1int > date2int
	})

	return entries
}

func SortComments(comments []config.Comment) []config.Comment {

	sort.Slice(comments, func(i, j int) bool {
		date1, _ := time.Parse("02.01.2006 um 15:04:05", comments[i].Date)
		date2, _ := time.Parse("02.01.2006 um 15:04:05", comments[j].Date)

		date1int, _ := strconv.Atoi(date1.Format("20060102150405"))
		date2int, _ := strconv.Atoi(date2.Format("20060102150405"))
		return date1int > date2int
	})

	return comments
}