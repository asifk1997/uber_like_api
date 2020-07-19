package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/asifk1997/uber_like_api/api/auth"
	"github.com/asifk1997/uber_like_api/api/models"
	"github.com/asifk1997/uber_like_api/api/responses"
	"github.com/asifk1997/uber_like_api/api/utils/formaterror"
	"github.com/gorilla/mux"
)

func (server *Server) CreateBooking(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	bookingUserEntry := models.BookingUserEntry{}
	err = json.Unmarshal(body, &bookingUserEntry)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = bookingUserEntry.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	bookingCreated, err := bookingUserEntry.SaveBooking(server.DB, uid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, bookingCreated.ID))
	responses.JSON(w, http.StatusCreated, bookingCreated)
}

func (server *Server) GetBookingsForUser(w http.ResponseWriter, r *http.Request) {

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	log.Println("uid", uid)
	booking := models.Booking{}

	bookings, err := booking.FindAllBookingsForAUser(server.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, bookings)
}

func (server *Server) GetBooking(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	log.Println("uid", uid)
	booking := models.Booking{}

	bookingReceived, err := booking.FindBookingByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	log.Println("booking ", bookingReceived, uid)
	if bookingReceived.UserID != uid {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	responses.JSON(w, http.StatusOK, bookingReceived)
}
