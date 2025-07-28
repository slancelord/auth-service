package router

import (
	"net/http"

	_ "auth-service/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func InitSwaggerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
}
