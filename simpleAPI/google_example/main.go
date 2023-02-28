package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"html/template"
	"io"
	"log"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("secret"))

const (
	baseURL  = "127.0.0.1"
	basePORT = "8080"
)

type myUser struct {
	Email string
	Name  string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RenderMainViewHandler)
	r.HandleFunc("/auth", RenderAuthViewHandler)
	r.HandleFunc("/auth/callback", AuthenticateHandler)
	r.HandleFunc("/view", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		email, _ := session.Values["user"].(string)
		name, _ := session.Values["username"].(string)
		user := &myUser{email, name}
		RenderTemplate(w, "view.html", user)
	})

	srv := &http.Server{
		Handler: r,
		Addr:    baseURL + ":" + basePORT,
	}
	log.Fatal(srv.ListenAndServe())
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, _ := template.ParseFiles(name)
	tmpl.Execute(w, data)
}

func RenderMainViewHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "main.html", nil)
}

func RenderAuthViewHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Options = &sessions.Options{
		Path:   "/auth",
		MaxAge: 300,
	}
	state := RandToken()
	session.Values["state"] = state
	session.Save(r, w)
	RenderTemplate(w, "auth.html", GetLoginURL(state))
}

func AuthenticateHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session")
	state := session.Values["state"]

	delete(session.Values, "state")
	session.Save(r, w)

	if state != r.FormValue("state") {
		http.Error(w, "Invalid session state", http.StatusUnauthorized)
		return
	}

	token, err := OAuthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	fmt.Println(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := OAuthConfig.Client(oauth2.NoContext, token)
	//client := OAuthConfig.Client(context.TODO(), token)
	userInfoResp, err := client.Get(UserInfoAPIEndpoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer userInfoResp.Body.Close()
	userInfo, err := io.ReadAll(userInfoResp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var authUser User
	json.Unmarshal(userInfo, &authUser)

	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400,
	}
	session.Values["user"] = authUser.Email
	session.Values["username"] = authUser.Name
	session.Save(r, w)

	http.Redirect(w, r, "/view", http.StatusFound)
}
