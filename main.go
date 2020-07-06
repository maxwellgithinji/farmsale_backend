package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" // For dev only, dont push to production
	// "os"

	// _ "github.com/maxwellgithinji/farmsale_backend/docs"
	_ "github.com/maxwellgithinji/farmsale_backend/docs"
	"github.com/maxwellgithinji/farmsale_backend/routes"
)

// @title Farmsale API
// @version 1.0
// @description This is a service for connecting farmers to customers
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email email@email.com

// @license.name MIT
// @license.url https://github.com/maxwellgithinji/farmsale_backend/blob/develop/LICENSE
//
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http https

func main() {
	// port := os.Getenv("PORT")
	http.Handle("/", routes.RouteHandlers())
	log.Fatal(http.ListenAndServe(":"+"8080", nil))
}
