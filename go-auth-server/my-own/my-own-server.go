package my_own

import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"net/http"
	"time"
)

var (
	idVar       string
	secretVar   string
	redirectVar string
	portVar     int
)

func main() {
	manager := manage.NewManager()

	authCodeConfig := &manage.Config{
		AccessTokenExp:    time.Hour * 2,
		RefreshTokenExp:   time.Hour * 24 * 3,
		IsGenerateRefresh: true,
	}

	implicitConfig := &manage.Config{
		AccessTokenExp: time.Hour * 1,
	}

	passwordTokenConfig := &manage.Config{
		AccessTokenExp:    time.Hour * 2,
		RefreshTokenExp:   time.Hour * 24 * 7,
		IsGenerateRefresh: true,
	}

	clientTokenConfig := &manage.Config{
		AccessTokenExp: time.Hour * 2,
	}

	refreshTokenConfig := &manage.RefreshingConfig{
		IsGenerateRefresh: false,
	}

	manager.SetAuthorizeCodeTokenCfg(authCodeConfig)
	manager.SetImplicitTokenCfg(implicitConfig)
	manager.SetPasswordTokenCfg(passwordTokenConfig)
	manager.SetClientTokenCfg(clientTokenConfig)
	manager.SetRefreshTokenCfg(refreshTokenConfig)

	manager.MapAuthorizeGenerate(generates.NewAuthorizeGenerate())
	manager.MapAccessGenerate(generates.NewAccessGenerate())

	manager.MustTokenStorage(store.NewMemoryTokenStore())
	clientStore := store.NewClientStore()
	clientStore.Set(idVar, &models.Client{
		ID:     idVar,
		Secret: secretVar,
		Domain: redirectVar,
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)

	srv.SetAllowedResponseType(
		oauth2.Code,
		oauth2.Token,
	)

	srv.SetAllowedGrantType(
		oauth2.AuthorizationCode,
		oauth2.PasswordCredentials,
		oauth2.ClientCredentials,
		oauth2.Implicit,
		oauth2.Refreshing,
	)

	srv.SetClientInfoHandler(myClientInfoHandler)
	srv.SetClientAuthorizedHandler(myClientAuthorizedHandler)
	srv.SetClientScopeHandler(myClientScopeHandler)

	srv.SetUserAuthorizationHandler(myUserAuthorizationHandler)
	srv.SetPasswordAuthorizationHandler(myPasswordAuthorizationHandler)

	srv.SetRefreshingScopeHandler(myRefreshingScopeHandler)
	srv.SetRefreshingValidationHandler(myValidateURIHandler)

	srv.SetResponseTokenHandler(myResponseTokenHandler)
	srv.SetResponseErrorHandler(myResponseErrorHandler)

	srv.SetInternalErrorHandler(myInternalErrorHandler)
	srv.SetExtensionFieldsHandler(myExtensionFieldsHandler)
	srv.SetAccessTokenExpHandler(myAccessTokenExpHandler)
	srv.SetAuthorizeScopeHandler(myAuthorizeScopeHandler)
}

// TODO: implement handlers...

func myClientInfoHandler(r *http.Request) (clientID, clientSecret string, err error) {
	return server.ClientBasicHandler(r)
}

func myClientAuthorizedHandler(clientID string, grant oauth2.GrantType) (allowed bool, err error) {
	return
}

func myClientScopeHandler(tgr *oauth2.TokenGenerateRequest) (allowed bool, err error) {
	return
}

func myUserAuthorizationHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	return
}

func myPasswordAuthorizationHandler(ctx context.Context, clientID, username, password string) (userID string, err error) {
	return
}

func myRefreshingScopeHandler(tgr *oauth2.TokenGenerateRequest, scope string) (allowed bool, err error) {
	return
}

func myValidateURIHandler(ti oauth2.TokenInfo) (allowed bool, err error) {
	return
}

func myResponseTokenHandler(w http.ResponseWriter, data map[string]interface{}, header http.Header, code ...int) error {
	return nil
}

func myResponseErrorHandler(re *errors.Response) {
	return
}

func myInternalErrorHandler(err error) (re *errors.Response) {
	return
}

func myExtensionFieldsHandler(ti oauth2.TokenInfo) (fieldsValue map[string]interface{}) {
	return
}

func myAccessTokenExpHandler(w http.ResponseWriter, r *http.Request) (exp time.Duration, err error) {
	return
}

func myAuthorizeScopeHandler(w http.ResponseWriter, r *http.Request) (scope string, err error) {
	return
}
