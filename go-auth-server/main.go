package main

import (
	"context"
	"flag"

	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/go-session/session"
)

var (
	dumpVar     bool
	idVar       string
	secretVar   string
	redirectVar string // redirect url
	portVar     int
)

// setup all vars
func init() {
	flag.BoolVar(&dumpVar, "d", true, "Dump requests and responses.")
	flag.StringVar(&idVar, "i", "222222", "The client id being passed in")
	flag.StringVar(&secretVar, "s", "22222222", "The client secret being passed in")
	flag.StringVar(&redirectVar, "r", "http://localhost:9094", "The domain of the redirect url")
	flag.IntVar(&portVar, "p", 9096, "the base port for the server")
}

func myPasswordAuthorizationHandler(ctx context.Context, clientID, username, password string) (userID string, err error) {
	if username == "test" && password == "test" {
		userID = "test"
		return userID, nil
	}

	return "", errors.New("not authorized user")
}

func main() {
	flag.Parse()
	manager := manage.NewDefaultManager()

	// set authorizeCodeToken config
	// manage.DefaultAuthorizeCodeTokenCfg has sample expire time
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token by uuid
	// accessToken, refreshToken, err will be returned
	manager.MapAccessGenerate(generates.NewAccessGenerate())
	//// generate JWT access token with keyId, key, Method
	//manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))

	// client memory store
	// NewClientStore returns map[string]oauth2.ClientInfo
	clientStore := store.NewClientStore()
	clientStore.Set(idVar, &models.Client{
		ID:     idVar,
		Secret: secretVar,
		Domain: redirectVar,
	})
	manager.MapClientStorage(clientStore)

	// same code
	// srv := server.NewServer(server.NewConfig(), manager)
	srv := server.NewDefaultServer(manager)

	srv.SetPasswordAuthorizationHandler(myPasswordAuthorizationHandler)

	srv.SetUserAuthorizationHandler(myUserAuthorizationHandler)

	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})

	log.Fatal(http.ListenAndServe(":9096", nil))
}

func dumpRequest(writer io.Writer, header string, r *http.Request) error {
	// some dump request...

	return nil

}
func myUserAuthorizationHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	// if dump request, ignore the error
	if dumpVar {
		_ = dumpRequest(os.Stdout, "userAuthorizeHandler", r)
	}

	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		return
	}
	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		// some response write ...
		return

	}
	userID = uid.(string)
	store.Delete("LoggedInUserID")
	store.Save()
	return
}
