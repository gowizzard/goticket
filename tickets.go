//**********************************************************
//
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package goticket

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type TicketBody struct {
	Alert       bool                `json:"alert"`
	Autorespond bool                `json:"autorespond"`
	Source      string              `json:"source"`
	Name        string              `json:"name"`
	Email       string              `json:"email"`
	Phone       string              `json:"phone"`
	Subject     string              `json:"subject"`
	Ip          string              `json:"ip"`
	Message     string              `json:"message"`
	Attachments []map[string]string `json:"attachments,omitempty"`
}

// CreateTicket is to create a new ticket via json api
func CreateTicket(data TicketBody, baseUrl, token string) (string, error) {

	// Create new client
	client := &http.Client{}

	// Convert json data
	convert, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Create url
	url := baseUrl + "/api/tickets.json"

	// Create new request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(convert))
	if err != nil {
		return "", err
	}

	// Define header for request
	request.Header.Set("X-API-Key", token)

	// Send request & get response
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	// Close response body
	defer response.Body.Close()

	// Read response
	read, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	// Return response
	return string(read), nil

}
