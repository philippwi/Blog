//Matrikelnummern: 3229403, 9964427

package tests

import (
	"Blog/utility"
	"sort"
	"time"
	"strconv"
	"testing"
)

func TestUserExists(t *testing.T){
	existingUsers := GetAllUsers()

	for _, user := range existingUsers {
		if user.Name == name {
			return true
		}
	}
	return false
}

func TestPasswordCorrect(t *testing.T){
	existingUsers := GetAllUsers()

	for _, user := range existingUsers {
		if name == user.Name && password == DecryptPassword(user.PwSalt) {
			return true
		}
	}
	return false
}

func TestEncryptPassword(t *testing.T) {
	return pw
}

func TestDecryptPassword(t *testing.T){
	return code
}

func TestGetBlog(t *testing.T){

	allBlogEntries := GetAllBlogEntries()

	for _, b := range allBlogEntries {
		if b.ID == blogID {
			blog = b
			break
		}
	}
	return blog
}

func TestGetBlogWithComments(t *testing.T){

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

/*func TestNewUserID(t *testing.T){

}*/

func TestNewBlogID(t *testing.T){
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

func TestNewCommentID(t *testing.T){
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

func TestSortBlogEntries(t *testing.T){

	sort.Slice(entries, func(i, j int) bool {
		date1, _ := time.Parse("02.01.2006 um 15:04:05", entries[i].Date)
		date2, _ := time.Parse("02.01.2006 um 15:04:05", entries[j].Date)

		date1int, _ := strconv.Atoi(date1.Format("20060102150405"))
		date2int, _ := strconv.Atoi(date2.Format("20060102150405"))
		return date1int > date2int
	})

	return entries
}

func TestSortComments(t *testing.T){

	sort.Slice(comments, func(i, j int) bool {
		date1, _ := time.Parse("02.01.2006 um 15:04:05", comments[i].Date)
		date2, _ := time.Parse("02.01.2006 um 15:04:05", comments[j].Date)

		date1int, _ := strconv.Atoi(date1.Format("20060102150405"))
		date2int, _ := strconv.Atoi(date2.Format("20060102150405"))
		return date1int > date2int
	})

	return comments
}