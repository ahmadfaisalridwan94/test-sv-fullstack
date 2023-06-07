// Package controllers provides controllers for handling HTTP requests to the API.
package controllers

import (
	"net/http"

	"backend/api/responses"
)

// Home handles requests to the root endpoint of the API.
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	// Create a response using the Response struct.
	response := responses.Response{
		ResponseCode:    "00",
		ResponseStatus:  true,
		ResponseMessage: "Test Backend Sharing Vision - Ahmad Faisal Ridwan",
	}

	// Send the response as JSON to the client.
	responses.JSON(w, http.StatusOK, response)
}
