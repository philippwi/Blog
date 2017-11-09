//Matrikelnummern: 3229403, 9964427

package server

import (
	"net/http"
	"testing"
	"net/http/httptest"
	"Blog/utility"
	"Blog/config"
	"html/template"
)

func getRequest(t testing.TB, url string) *http.Request {

	tpl = template.Must(template.ParseGlob(utility.FixPath(config.HtmlDir) + "*.html"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func TestLoginPage(t *testing.T) {
	r := getRequest(t, "/login")

	rw := httptest.NewRecorder()

	LoginPage(rw, r)

}