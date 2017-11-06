//Matrikelnummern: 3229403, 9964427

package server

import (
	"Blog/config"
	"net/http"
	"html/template"
	"Blog/dataHandling"
	"strconv"
	"time"
	"Blog/utility"
)

//Variable zur HTML Template Nutzung
var tpl *template.Template

//Variable zum Speichern der festgelegten Sitzungsdauer
var sesExp time.Duration

//initialisiert nötige Einstellungen und startet HTTPS Server
func StartServer(sessionExp int, port string) {
	sesExp = time.Duration(sessionExp) * time.Minute
	tpl = template.Must(template.ParseGlob(utility.FixPath(config.HtmlDir) + "*.html"))
	http.HandleFunc("/", LoginPage)
	http.HandleFunc("/changepw", ChangePw)
	http.HandleFunc("/home", HomePage)
	http.HandleFunc("/viewblog", ViewblogPage)
	http.HandleFunc("/editblog", EdtBlg)
	http.HandleFunc("/deleteblog", DltBlog)
	http.HandleFunc("/logout", Logout)
	http.ListenAndServeTLS(":"+port, utility.FixPath(config.ServerDir)+"cert.pem", config.ServerDir+"key.pem", nil)
}

//Anmeldeseite mit Login-Eingabe-Verarbeitung
func LoginPage(wr http.ResponseWriter, rq *http.Request) {
	if rq.Method == http.MethodPost {
		userName := rq.FormValue("usrnm")
		password := rq.FormValue("passw")

		if !dataHandling.UserExists(userName) { //Nutzer existiert nicht
			dataHandling.SaveUser(userName, password)
			cookie := http.Cookie{Name: "user", Value: userName, Expires: time.Now().Add(sesExp)}
			http.SetCookie(wr, &cookie)
			http.Redirect(wr, rq, "/home", http.StatusFound)
		} else { //Nutzer existiert bereits
			if dataHandling.PasswordCorrect(userName, password) { //Passwort korrekt
				cookie := http.Cookie{Name: "user", Value: userName, Expires: time.Now().Add(sesExp)}
				http.SetCookie(wr, &cookie)
				http.Redirect(wr, rq, "/home", http.StatusFound)
			} else { //Passwort falsch
				//http.Redirect(wr, rq, "/", http.st)
			}
		}
	}
	tpl.ExecuteTemplate(wr, "index.html", nil)
}

//Homepage zur Darstellung der Blogübersicht mit Möglichkeit der Blogerstellung
func HomePage(wr http.ResponseWriter, rq *http.Request) {
	var currentUser string

	if !IsUserLoggedIn(rq) {
		currentUser = ""
	} else {
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
		tpl.ExecuteTemplate(wr, "message.html", config.Message{
			MsgText:  "Blog gespeichert",
			Redirect: "home"})
	}

	pageData := config.HomeData{
		CurrentUser: currentUser,
		BlogEntries: dataHandling.GetAllBlogEntries(),
	}

	tpl.ExecuteTemplate(wr, "home.html", pageData)
}

//Darstellung eines bestimmten Blogs mit Bearbeitungs-, Kommentar- und Löschmöglichkeit
func ViewblogPage(wr http.ResponseWriter, rq *http.Request) {
	var currentUser string

	if !IsUserLoggedIn(rq) {
		currentUser = ""
	} else {
		currentUser = GetCurrentUsername(rq)
	}

	blogID, err := strconv.Atoi(rq.URL.Query()["ID"][0])

	if err != nil{}

	blog, blogComments := dataHandling.GetBlogWithComments(blogID)

	pageData := config.ViewblogData{
		CurrentUser:  currentUser,
		Blog:         blog,
		BlogComments: blogComments}

	if rq.Method == http.MethodPost {
		var author string
		if currentUser == "" {
			author = rq.FormValue("nicknm") + " (Leser)"
		} else {
			author = currentUser
		}
		commentText := rq.FormValue("cmnt")
		dataHandling.SaveComment(author, commentText, blog.ID)
		http.Redirect(wr, rq, "/viewblog?ID="+strconv.Itoa(blogID), http.StatusFound)
	}

	tpl.ExecuteTemplate(wr, "viewblog.html", pageData)
}

//Änderung des Bloginhaltes
func EdtBlg(wr http.ResponseWriter, rq *http.Request) {
	var currentUser string

	if !IsUserLoggedIn(rq) {
		currentUser = ""
	} else {
		currentUser = GetCurrentUsername(rq)
	}

	blogID, _ := strconv.Atoi(rq.URL.Query()["ID"][0])

	blogContent := dataHandling.GetBlog(blogID).Content

	pageData := config.ChangeblogData{
		CurrentUser: currentUser,
		BlogContent: blogContent,
	}

	if rq.Method == http.MethodPost {
		if !IsUserLoggedIn(rq) {
			tpl.ExecuteTemplate(wr, "message.html", config.Message{
				MsgText:  "Session expired",
				Redirect: "logout"})
		} else {
			newContent := rq.FormValue("blgcont")
			dataHandling.ChangeBlogEntry(newContent, blogID)
			http.Redirect(wr, rq, "/viewblog?ID="+strconv.Itoa(blogID), http.StatusFound)
		}
	}

	tpl.ExecuteTemplate(wr, "editblog.html", pageData)
}

//Löschung eines Blogeintrages
func DltBlog(wr http.ResponseWriter, rq *http.Request) {

	if !IsUserLoggedIn(rq) {
		tpl.ExecuteTemplate(wr, "message.html", config.Message{
			MsgText:  "Session expired",
			Redirect: "logout"})
		return
	}

	blogID, _ := strconv.Atoi(rq.URL.Query()["ID"][0])

	dataHandling.DeleteBlogEntry(blogID)
	http.Redirect(wr, rq, "/home", http.StatusFound)
}

//Nutzerpasswort prüfen und ändern
func ChangePw(wr http.ResponseWriter, rq *http.Request) {
	var currentUser string

	if !IsUserLoggedIn(rq) {
		tpl.ExecuteTemplate(wr, "message.html", config.Message{
			MsgText:  "Session expired",
			Redirect: "logout"})
		return
	} else {
		currentUser = GetCurrentUsername(rq)
	}

	if rq.Method == http.MethodPost {
		if !IsUserLoggedIn(rq) {
			tpl.ExecuteTemplate(wr, "message.html", config.Message{
				MsgText:  "Session expired",
				Redirect: "logout"})
		} else {
			oldPw := rq.FormValue("currpw")
			newPw1 := rq.FormValue("newpw1")
			newPw2 := rq.FormValue("newpw2")

			if !dataHandling.PasswordCorrect(currentUser, oldPw) {
				tpl.ExecuteTemplate(wr, "message.html", config.Message{
					MsgText:  "Passwort falsch",
					Redirect: "changepw"})
			} else {
				if newPw1 == newPw2 {
					dataHandling.ChangeUserPassword(currentUser, newPw1)
					tpl.ExecuteTemplate(wr, "message.html", config.Message{
						MsgText:  "Passwort geändert",
						Redirect: "home"})
				} else {
					tpl.ExecuteTemplate(wr, "message.html", config.Message{
						MsgText:  "Passwörter stimmen nicht überein",
						Redirect: "changepw"})
				}
			}
		}
	}
	tpl.ExecuteTemplate(wr, "changepw.html", nil)
}

//Sitzung eines angemeldeten Nutzers beenden
func Logout(wr http.ResponseWriter, rq *http.Request) {
	cookie := http.Cookie{Name: "user", Value: "", Expires: time.Now()}
	http.SetCookie(wr, &cookie)

	http.Redirect(wr, rq, "/", http.StatusFound)
}

//testet ob Aufrufer ein angemeldeter Nutzer ist
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

//liefert Name des aktuell angemelden Nutzers
func GetCurrentUsername(rq *http.Request) string {
	cookie, err := rq.Cookie("user")
	if err != nil {
		return "error"
	}
	return cookie.Value
}
