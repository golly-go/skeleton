package controllers

import (
	// Slimloans custom helpers

	"net/http"

	"github.com/slimloans/golly"
	"github.com/slimloans/golly/middleware"
)

// Routes is the entry point to the systems routes
func Routes(a golly.Application) {
	a.Routes().
		Use(middleware.Recoverer, middleware.RequestLogger).
		Use(middleware.Cors(middleware.CorsOptions{
			AllowedHeaders:   a.Config.GetStringSlice("cors.headers"),
			AllowedOrigins:   a.Config.GetStringSlice("cors.origins"),
			AllowedMethods:   a.Config.GetStringSlice("cors.methods"),
			AllowCredentials: true,
		})).
		Namespace("/", func(r *golly.Route) {
			r.Get("/status", func(wctx golly.WebContext) { wctx.RenderStatus(http.StatusOK) })
		})
}