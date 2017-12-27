//Matrikelnummern: 3229403, 9964427

package dataHandling

import (
	"Blog/utility"
	"Blog/config"
	"sort"
	"time"
	"strconv"
	"crypto/sha1"
	"fmt"
	"errors"
)

//testet ob übergebener Nutzername existiert
func UserExists(name string) bool {
	existingUsers := GetAllUsers()

	for i,_ := range existingUsers {
		if existingUsers[i].Name == name {
			return true
		}
	}
	return false
}

//testet ob Nutzer-Passwort komination korrekt ist
func PasswordCorrect(name, password string) bool {
	existingUsers := GetAllUsers()

	for i, _ := range existingUsers {
		if name == existingUsers[i].Name{
			if EncryptPassword(password, name) == existingUsers[i].PwSalt {
				return true
			}
			break
		}
	}
	return false
}

//verchlüsselt Passwort
func EncryptPassword(pw string, userName string) string {

	saltedPw := pw + userName

	h := sha1.New()

	h.Write([]byte(saltedPw))

	encryptedPw := h.Sum(nil)

	return fmt.Sprintf("%x", encryptedPw)
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
		date1, err1 := time.Parse("02.01.2006 um 15:04:05", entries[i].Date)
		date2, err2 := time.Parse("02.01.2006 um 15:04:05", entries[j].Date)

		date1int, err3 := strconv.Atoi(date1.Format("20060102150405"))
		date2int, err4 := strconv.Atoi(date2.Format("20060102150405"))

		if err1 != nil || err2 != nil ||
			err3 != nil || err4 != nil {
			utility.HandleError("control.SortBlogEntries -> Datumskonvertierung", errors.New("Date conversion error"))
		}

		return date1int > date2int
		})
	return entries
}

//sortiert Kommentare nach Datum (von neu nach alt)
func SortComments(comments []config.Comment) []config.Comment {

	sort.Slice(comments, func(i, j int) bool {
		date1, err1 := time.Parse("02.01.2006 um 15:04:05", comments[i].Date)
		date2, err2 := time.Parse("02.01.2006 um 15:04:05", comments[j].Date)

		date1int, err3 := strconv.Atoi(date1.Format("20060102150405"))
		date2int, err4 := strconv.Atoi(date2.Format("20060102150405"))

		if err1 != nil || err2 != nil ||
			err3 != nil || err4 != nil {
			utility.HandleError("control.SortComments -> Datumskonvertierung", errors.New("Date conversion error"))
		}

		return date1int > date2int
	})
	return comments
}