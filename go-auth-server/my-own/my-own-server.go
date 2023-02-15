package my_own

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
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

}
