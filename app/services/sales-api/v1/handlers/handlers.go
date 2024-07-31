package handlers

import (
	"github.com/elisardofelix/ultimatego/app/services/sales-api/v1/handlers/checkgrp"
	"github.com/elisardofelix/ultimatego/app/services/sales-api/v1/handlers/hackgrp"
	"github.com/elisardofelix/ultimatego/app/services/sales-api/v1/handlers/usergrp"
	v1 "github.com/elisardofelix/ultimatego/business/web/v1"
	"github.com/elisardofelix/ultimatego/foundation/web"
)

type Routes struct{}

// Add implements the RouterAdder interface to add all routes.
func (Routes) Add(app *web.App, apiCfg v1.APIMuxConfig) {
	hackgrp.Routes(app, hackgrp.Config{
		Auth: apiCfg.Auth,
	})

	checkgrp.Routes(app, checkgrp.Config{
		Build: apiCfg.Build,
		Log:   apiCfg.Log,
		DB:    apiCfg.DB,
	})

	usergrp.Routes(app, usergrp.Config{
		Build: apiCfg.Build,
		Log:   apiCfg.Log,
		DB:    apiCfg.DB,
		Auth:  apiCfg.Auth,
	})
}
