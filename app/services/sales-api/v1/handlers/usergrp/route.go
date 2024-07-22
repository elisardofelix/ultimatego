package usergrp

import (
	"net/http"

	"github.com/elisardofelix/ultimatego/business/core/user"
	"github.com/elisardofelix/ultimatego/business/core/user/stores/userdb"
	"github.com/elisardofelix/ultimatego/business/web/v1/auth"
	"github.com/elisardofelix/ultimatego/foundation/logger"
	"github.com/elisardofelix/ultimatego/foundation/web"
	"github.com/jmoiron/sqlx"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build string
	Log   *logger.Logger
	DB    *sqlx.DB
	Auth  *auth.Auth
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	usrCore := user.NewCore(cfg.Log, userdb.NewStore(cfg.Log, cfg.DB))

	hdl := New(usrCore, cfg.Auth)
	app.Handle(http.MethodPost, version, "/users", hdl.Create)
}
