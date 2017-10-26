//Matrikelnummern: 3229403, 9964427

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
	fmt.Println("Server running: https://localhost" + config.Port)
	http.HandleFunc("/", loginPage)
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/viewblog", viewblogPage)
	http.HandleFunc("/logout", Logout)
	http.ListenAndServeTLS(config.Port, config.ServerDir+"cert.pem",config.ServerDir+"key.pem", nil)
}

func loginPage(wr http.ResponseWriter, rq *http.Request) {
	if rq.Method == http.MethodPost {
		userName := rq.FormValue("usrnm")
		password := rq.FormValue("passw")

		if !dataHandling.UserExists(userName) { //Nutzer existiert nicht
			dataHandling.SaveUser(userName, password)
			cookie := http.Cookie{Name: "user", Value: userName, Expires: time.Now().Add(10 * time.Minute)}
			http.SetCookie(wr, &cookie)
			http.Redirect(wr, rq, "/home", http.StatusFound)
		} else { //Nutzer existiert bereits
			if dataHandling.PasswordCorrect(userName, password) { //Passwort korrekt
				cookie := http.Cookie{Name: "user", Value: userName, Expires: time.Now().Add(10 * time.Minute)}
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
	var currentUser string

	if !IsUserLoggedIn(rq) {
		currentUser = ""
	}else{
		currentUser = GetCurrentUsername(rq)
	}

	if rq.Method == http.MethodPost {
		if !IsUserLoggedIn(rq) {
			http.Redirect(wr, rq, "/", http.StatusFound)
			return
		}
		author := currentUser
		title := rq.FormValue("blgtitle")
		content := rq.FormValue("blgcont")
		dataHandling.SaveBlogEntry(author, title, content)
		http.Redirect(wr, rq, "/home", http.StatusFound)
	}

	pageData := config.HomeData{
		CurrentUser: currentUser,
		BlogEntries: dataHandling.GetAllBlogEntries(),
	}

	tpl.ExecuteTemplate(wr, "home.html", pageData)
}

func viewblogPage(wr http.ResponseWriter, rq *http.Request) {
	var currentUser string

	if !IsUserLoggedIn(rq) {
		currentUser = ""
	}else{
		currentUser = GetCurrentUsername(rq)
	}

	blogID, _ := strconv.Atoi(rq.URL.Query()["ID"][0])

	blog, blogComments := dataHandling.GetBlogWithComments(blogID)

	pageData := config.ViewblogData{
		CurrentUser: currentUser,
		Blog: blog,
		BlogComments: blogComments}

	if rq.Method == http.MethodPost {
		var author string
		if currentUser == ""{
			author = rq.FormValue("nicknm")+" (Leser)"
		}else{
			author = currentUser
		}
		commentText := rq.FormValue("cmnt")
		dataHandling.SaveComment(author, commentText, blog.ID)
		http.Redirect(wr, rq, "/viewblog?ID="+strconv.Itoa(blogID), http.StatusFound)
	}
	tpl.ExecuteTemplate(wr, "viewblog.html", pageData)
}

func IsUserLoggedIn(rq *http.Request) bool {
	cookie, err := rq.Cookie("user")

	if err != nil {
		return false
	}

	if !dataHandling.UserExists(cookie.Value) {
		return false
	}

	return true
}

func GetCurrentUsername(rq *http.Request) string {
	cookie, err := rq.Cookie("user")
	if err != nil {
		return "error"
	}
	return cookie.Value
}

func Logout(wr http.ResponseWriter, rq *http.Request) {
	cookie := http.Cookie{Name: "user", Value: "", Expires: time.Now()}
	http.SetCookie(wr, &cookie)

	http.Redirect(wr, rq, "/", http.StatusFound)
}
