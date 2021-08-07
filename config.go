//**********************************************************
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of goticket.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package goticket

import (
	"bytes"
	"net/http"
)

const (
	baseURL = "https://api.lexoffice.io"
)

// Request is to define the request data
type Request struct {
	Path, Method, Token string
	Body                []byte
}

// Send is to send a new request
func (r *Request) Send() (*http.Response, error) {

	// Set url
	url := baseURL + r.Path

	// Define client
	client := &http.Client{}

	// Request
	request, err := http.NewRequest(r.Method, url, bytes.NewBuffer(r.Body))
	if err != nil {
		return nil, err
	}

	// Define header
	request.Header.Set("Authorization", "Bearer "+r.Token)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	// Send request & get response
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Return data
	return response, nil

}
