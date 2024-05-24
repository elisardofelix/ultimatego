package handlers

import (
	"github.com/dimfeld/httptreemux/v5"
	"github.com/elisardofelix/ultimatego/app/services/sales-api/v1/handlers/hackgrp"
	v1 "github.com/elisardofelix/ultimatego/business/web/v1"
)

type Routes struct{}

// Add implements the RouterAdder interface.
func (Routes) Add(mux *httptreemux.ContextMux, cfg v1.APIMuxConfig) {
	hackgrp.Routes(mux)
}
