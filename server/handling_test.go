//Matrikelnummern: 3229403, 9964427

package server

import (
	"net/http"
	"testing"
	"net/http/httptest"
	"Blog/utility"
	"html/template"
	"Blog/config"

)

func getRequest(t testing.TB, url string) *http.Request {

	tpl = template.Must(template.ParseGlob(utility.FixPath(config.HtmlDir) + "*.html"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func TestChangePw(t *testing.T) {
	r := getRequest(t, "/changepw")

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(ChangePw)

	handler.ServeHTTP(rec, r)

	if sc := rec.Code; sc != http.StatusOK {
		t.Errorf("ChangePw-Handler status code = %v statt %v",
			sc, http.StatusOK)
	}
}

func TestLoginPage(t *testing.T) {
	r := getRequest(t, "/login")

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginPage)

	handler.ServeHTTP(rec, r)

	if sc := rec.Code; sc != http.StatusOK {
		t.Errorf("LoginPage-Handler status code = %v statt %v",
			sc, http.StatusOK)
	}
}

func TestViewblogPage(t *testing.T) {
	testBlogID := "0" //Test-Blog 0 sollte vorhanden sein
	r := getRequest(t, "/viewblog?ID="+testBlogID)

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(ViewblogPage)

	handler.ServeHTTP(rec, r)

	if sc := rec.Code; sc != http.StatusOK {
		t.Errorf("ViewblogPage-Handler status code = %v statt %v",
			sc, http.StatusOK)
	}
}

func TestEdtBlg(t *testing.T) {
	testBlogID := "0" //Test-Blog 0 sollte vorhanden sein
	r := getRequest(t, "/editblog?ID="+testBlogID)

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(EdtBlg)

	handler.ServeHTTP(rec, r)

	if sc := rec.Code; sc != http.StatusOK {
		t.Errorf("EdtBlg-Handler status code = %v statt %v",
			sc, http.StatusOK)
	}
}

func TestDltBlog(t *testing.T) {
	r := getRequest(t, "/deleteblog")

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(DltBlog)

	handler.ServeHTTP(rec, r)

	if sc := rec.Code; sc != http.StatusOK {
		t.Errorf("DltBlog-Handler status code = %v statt %v",
			sc, http.StatusOK)
	}
}

func TestLogout(t *testing.T) {
	r := getRequest(t, "/logout")

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(Logout)

	handler.ServeHTTP(rec, r)

	if sc := rec.Code; sc != http.StatusFound {
		t.Errorf("Logout-Handler status code = %v statt %v",
			sc, http.StatusFound)
	}
}

func TestIsUserLoggedIn(t *testing.T)  {
	existentUser := "TestUser"
	nonexistentUser := "x"

	rec1, rec2, rec3 := httptest.NewRecorder(), httptest.NewRecorder(), httptest.NewRecorder()

	http.SetCookie(rec1, &http.Cookie{
		Name:  "user",
		Value: utility.EncryptCookie(existentUser),
	})
	http.SetCookie(rec2, &http.Cookie{
		Name:  "user",
		Value: utility.EncryptCookie(nonexistentUser),
	})

	req1 := &http.Request{Header: http.Header{"Cookie": rec1.HeaderMap["Set-Cookie"]}}
	req2 := &http.Request{Header: http.Header{"Cookie": rec2.HeaderMap["Set-Cookie"]}}
	req3 := &http.Request{Header: http.Header{"Cookie": rec3.HeaderMap["Set-Cookie"]}}

	if !IsUserLoggedIn(req1){
		t.Errorf("User '%v' wurde nicht erkannt", existentUser)
	}
	if IsUserLoggedIn(req2){
		t.Errorf("User '%v' wurde fälschlicherweise erkannt", nonexistentUser)
	}
	if IsUserLoggedIn(req3){
		t.Error("User wurde erkannt, obwohl kein Cookie vorhanden")
	}
}

func TestGetCurrentUsername(t *testing.T)  {
	usr := "TestUser12345;-{]}?"

	rec1, rec2 := httptest.NewRecorder(), httptest.NewRecorder()

	http.SetCookie(rec1, &http.Cookie{
		Name:  "user",
		Value: utility.EncryptCookie(usr),
	})

	req1 := &http.Request{Header: http.Header{"Cookie": rec1.HeaderMap["Set-Cookie"]}}
	req2 := &http.Request{Header: http.Header{"Cookie": rec2.HeaderMap["Set-Cookie"]}}

	curUsr1 := GetCurrentUsername(req1)
	curUsr2 := GetCurrentUsername(req2)

	if  curUsr1 != usr {
		t.Errorf("Cookie1-Wert ist '%v' statt '%v'", curUsr1, usr)
	}

	if  curUsr2 != ""{
		t.Error("Cookie2-Wert ist nicht leer")
	}
}

func TestGetCurrentNick(t *testing.T) {
	nick := "TestNick12345;-{]}?"

	rec1, rec2 := httptest.NewRecorder(), httptest.NewRecorder()

	http.SetCookie(rec1, &http.Cookie{
		Name:  "nick",
		Value: utility.EncryptCookie(nick),
	})

	req1 := &http.Request{Header: http.Header{"Cookie": rec1.HeaderMap["Set-Cookie"]}}
	req2 := &http.Request{Header: http.Header{"Cookie": rec2.HeaderMap["Set-Cookie"]}}

	curNick1 := GetCurrentNick(req1)
	curNick2 := GetCurrentNick(req2)

	if  curNick1 != nick {
		t.Errorf("Cookie1-Wert ist '%v' statt '%v'", curNick1, nick)
	}

	if  curNick2 != ""{
		t.Error("Cookie2-Wert ist nicht leer")
	}
}
