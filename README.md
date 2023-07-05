# Facteur Go SDK

> Send emails in your Go applications with Facteur.

## Installation

```bash
go get github.com/facteurdev/facteur-go
```

## Usage

```go
package main

import "github.com/facteurdev/facteur-go"

func main() {
	payload := &facteur.SendEmailPayload{
    From: "no-reply@example.com",
    To: "ayn@rand.com",
    Subject: "Who is John Galt? From Golang!",
    Text: "I started my life with a single absolute: that the world was mine to shape in the image of my highest values and never to be given up to a lesser standard, no matter how long or hard the struggle.",
    HTML: "<p>I started my life with <b>a single absolute</b>: that the world was mine to shape in the image of my highest values and never to be given up to a lesser standard, no matter how long or hard the struggle.</p>",
	}

	f := facteur.NewFacteur("<YOUR_API_KEY>")
	err := f.SendEmail(payload)
	if err != nil {
		fmt.Println(err)
	}
}
```
