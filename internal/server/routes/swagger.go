package routes

import (
	"net/http"

	_ "auth-service/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func InitSwaggerHandler(mux *http.ServeMux) {
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
}
