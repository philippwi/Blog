package server

import (
	"fmt"
	"Blog/config"
	"net/http"
	"html/template"
	"Blog/dataHandling"
	"strconv"
)

var tpl *template.Template

func StartServer() {
	tpl = template.Must(template.ParseGlob("server/*.html"))
	fmt.Println("Server running: http://localhost" + config.Port)
	http.HandleFunc("/", loginPage)
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/viewblog", displayBlog)
	http.ListenAndServe(config.Port, nil)
}

func loginPage(wr http.ResponseWriter, rq *http.Request) {
	reqPath := rq.URL.Path
	if reqPath != "/" && reqPath != "/index" {
		//	return
	}

	if rq.Method == http.MethodPost {
		userName := rq.FormValue("usrnm")
		password := rq.FormValue("passw")

		if !dataHandling.UserExists(userName) { //Nutzer existiert nicht
			dataHandling.SaveUser(userName, password)
			http.Redirect(wr, rq, "/home", http.StatusFound)
		} else { //Nutzer existiert bereits
			if dataHandling.PasswordCorrect(userName, password) { //Passwort korrekt
				http.Redirect(wr, rq, "/home", http.StatusFound)
			} else { //Passwort falsch
				//http.Redirect(wr, rq, "/", http.st)
			}
		}
	}
	tpl.ExecuteTemplate(wr, "index.html", nil)
}

func homePage(wr http.ResponseWriter, rq *http.Request) {
	if rq.Method == http.MethodPost {
		author := "TestUser" //braucht aktuellen nutzer
		title := rq.FormValue("blgtitle")
		content := rq.FormValue("blgcont")
		dataHandling.SaveBlogEntry(author, title, content)
		http.Redirect(wr, rq, "/home", http.StatusFound)
	}
	tpl.ExecuteTemplate(wr, "home.html", dataHandling.GetBlogEntryList())
}

func displayBlog(wr http.ResponseWriter, rq *http.Request) {
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
		commentText := rq.FormValue("cmnt")
		dataHandling.SaveComment("TestUser", commentText, blog.ID)
		http.Redirect(wr, rq, "/viewblog?ID="+strconv.Itoa(blogID), http.StatusFound)
	}

	tpl.ExecuteTemplate(wr, "viewblog.html", data)
}
