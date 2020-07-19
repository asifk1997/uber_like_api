package controllers

import (
	"net/http"

	"github.com/asifk1997/uber_like_api/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
