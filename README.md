# goticket

[![Go](https://github.com/gowizzard/goticket/actions/workflows/go.yml/badge.svg)](https://github.com/gowizzard/goticket/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/gowizzard/goticket)](https://goreportcard.com/report/github.com/gowizzard/goticket) 

Is a library for the ticket software osTicket. [Here](https://osticket.com/) you can find their website. With this small library you can create a new ticket in json format via the Rest API. A small tutorial can be found here.

## Install

```console
go get github.com/gowizzard/goticket
```

## How to use?

With a function you can create a ticket directly in your system. How to create an API token can be found [here](https://docs.osticket.com/en/latest/Developer%20Documentation/API%20Docs.html).

### Create ticket

To give you a better understanding of the data to be transferred, I'll show you the structure for creating the JSON data. Here you can see which values are needed. Below you will find an example.

```go
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
```

With this function you create a ticket.

```go
// Ticket body
body := &goticket.TicketBody{
    Alert:       true,
    Autorespond: true,
    Source:      "API",
    Name:        "Jonas Kwiedor",
    Email:       "jonas.kwiedor@jj-ideenschmiede.de",
    Phone:       "",
    Subject:     "Sehr sauer",
    Ip:          "data:text/html,Message <b>Was ist da los????</b>",
    Message:     "",
    Attachments: nil,
}

// Add file
file := make(map[string]string)
file["files.txt"] = "data:text/plain;charset=utf-8,content"

// Add file to attachment
body.Attachments = append(body.Attachments, file)

// Create test ticket
ticket, err := goticket.CreateTicket(body, "https://base.url.de", "token")
if err != nil {
    fmt.Println(err)
}

// Print ticket response
fmt.Println(ticket)
```

## Help
If you have any questions or comments, please contact us by e-mail at [jonas.kwiedor@jj-ideenschmiede.de](mailto:jonas.kwiedor@jj-ideenschmiede.de)