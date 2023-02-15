package my_own

import (
	"github.com/go-oauth2/oauth2/v4/manage"
	"time"
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

}
