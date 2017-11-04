//Matrikelnummern: 3229403, 9964427

package dataHandling

import (
	"Blog/utility"
	"Blog/config"
	"sort"
	"time"
	"strconv"
)

//testet ob übergebener Nutzername existiert
func UserExists(name string) bool {
	existingUsers := GetAllUsers()

	for _, user := range existingUsers {
		if user.Name == name {
			return true
		}
	}
	return false
}

//testet ob Nutzer-Passwort komination korrekt ist
func PasswordCorrect(name, password string) bool {
	existingUsers := GetAllUsers()

	for _, user := range existingUsers {
		if name == user.Name && password == DecryptPassword(user.PwSalt) {
			return true
		}
	}
	return false
}

//verchlüsselt Passwort
func EncryptPassword(pw string) string {
	return pw
}

//entschlüsselt Passwort
func DecryptPassword(code string) string {
	return code
}

//generiert freie ID für einen Blogeintrag
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

//sortiert Blogeinträge nach Datum (von neu nach alt)
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

//sortiert Kommentare nach Datum (von neu nach alt)
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