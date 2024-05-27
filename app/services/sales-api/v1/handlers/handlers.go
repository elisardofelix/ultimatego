package handlers

import (
	"github.com/elisardofelix/ultimatego/app/services/sales-api/v1/handlers/hackgrp"
	v1 "github.com/elisardofelix/ultimatego/business/web/v1"
	"github.com/elisardofelix/ultimatego/foundation/web"
)

type Routes struct{}

// Add implements the RouterAdder interface.
func (Routes) Add(app *web.App, cfg v1.APIMuxConfig) {
	hackgrp.Routes(app)
}
