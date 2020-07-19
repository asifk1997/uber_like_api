package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/asifk1997/uber_like_api/api/models"
	"github.com/asifk1997/uber_like_api/api/responses"
)

func (server *Server) GetNearbyCabs(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	getCab := models.UserCarrentLocationAndRange{}
	err = json.Unmarshal(body, &getCab)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	log.Println(getCab)
	cab := models.Cab{}
	cabs, err := cab.FindNearByCabs(server.DB, getCab)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, cabs)
}
