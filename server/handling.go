package server

import (
	"fmt"
	"Blog/config"
	"net/http"
	"html/template"
	"Blog/dataHandling"
	"strconv"
	"time"
)

var tpl *template.Template

func StartServer() {
	tpl = template.Must(template.ParseGlob("server/*.html"))
	fmt.Println("Server running: http://localhost" + config.Port)
	http.HandleFunc("/", loginPage)
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/viewblog", displayBlog)
	http.HandleFunc("/logout", Logout)
	http.ListenAndServe(config.Port, nil)
}

func loginPage(wr http.ResponseWriter, rq *http.Request) {
		if rq.Method == http.MethodPost {
		userName := rq.FormValue("usrnm")
		password := rq.FormValue("passw")

		if !dataHandling.UserExists(userName) { //Nutzer existiert nicht
			dataHandling.SaveUser(userName, password)
			cookie := http.Cookie{Name: "user", Value: userName, Expires: time.Now().Add(365*24*time.Hour)}
			http.SetCookie(wr, &cookie)
			http.Redirect(wr, rq, "/home", http.StatusFound)
		} else { //Nutzer existiert bereits
			if dataHandling.PasswordCorrect(userName, password) { //Passwort korrekt
				cookie := http.Cookie{Name: "user", Value: userName, Expires: time.Now().Add(365*24*time.Hour)}
				http.SetCookie(wr, &cookie)
				http.Redirect(wr, rq, "/home", http.StatusFound)
			} else { //Passwort falsch
				//http.Redirect(wr, rq, "/", http.st)
			}
		}
	}
	tpl.ExecuteTemplate(wr, "index.html", nil)
}

func homePage(wr http.ResponseWriter, rq *http.Request) {
	if !IsUserLoggedIn(rq){
		http.Redirect(wr, rq, "/", http.StatusFound)
		return
	}

	if rq.Method == http.MethodPost {
		if !IsUserLoggedIn(rq){
			http.Redirect(wr, rq, "/", http.StatusFound)
			return
		}
		author := GetCurrentUsername(rq.Cookie("user"))
		title := rq.FormValue("blgtitle")
		content := rq.FormValue("blgcont")
		dataHandling.SaveBlogEntry(author, title, content)
		http.Redirect(wr, rq, "/home", http.StatusFound)
	}
	tpl.ExecuteTemplate(wr, "home.html", dataHandling.GetBlogEntryList())
}

func displayBlog(wr http.ResponseWriter, rq *http.Request) {
	if !IsUserLoggedIn(rq){
		http.Redirect(wr, rq, "/", http.StatusFound)
		return
	}
	blogID, _ := strconv.Atoi(rq.URL.Query()["ID"][0])

	allBlogs := dataHandling.GetBlogEntryList()
	allComments := dataHandling.GetCommentList()

	blog := config.BlogEntry{}
	blogComments := []config.Comment{}

	for _, b := range allBlogs.BlogEntries {
		if b.ID == blogID {
			blog = b
			break
		}
	}

	for _, c := range allComments.Comments {
		if c.BlogID == blogID {
			blogComments = append(blogComments, c)
		}
	}

	data := config.ViewblogData{Blog: blog, BlogComments: blogComments}

	if rq.Method == http.MethodPost {
		if !IsUserLoggedIn(rq){
			http.Redirect(wr, rq, "/", http.StatusFound)
			return
		}
		author := GetCurrentUsername(rq.Cookie("user"))
		commentText := rq.FormValue("cmnt")
		dataHandling.SaveComment(author, commentText, blog.ID)
		http.Redirect(wr, rq, "/viewblog?ID="+strconv.Itoa(blogID), http.StatusFound)
	}

	tpl.ExecuteTemplate(wr, "viewblog.html", data)
}

func IsUserLoggedIn(rq *http.Request) bool{
	cookie, err := rq.Cookie("user")

	if err != nil{
		return false
	}

	if !dataHandling.UserExists(cookie.Value){
		return false
	}

	return true
}

func GetCurrentUsername(cookie *http.Cookie, err error) string{
	if err != nil{
		return "error"
	}
	return cookie.Value
}

func Logout(wr http.ResponseWriter, rq *http.Request) {
	cookie, err := rq.Cookie("user")

	if err != nil {
		return
	}

	geeehhht nichtttt


	cookie.Expires = time.Now().Add(365*24*time.Hour)

	http.Redirect(wr, rq, "/", http.StatusFound)
}