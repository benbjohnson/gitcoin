package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	host := flag.String("host", "localhost:10000", "remote host")
	flag.Parse()
	log.SetFlags(0)

	// Grab the first arg as the message.
	message := flag.Arg(0)

	// Construct URL.
	q := url.Values{}
	q.Set("message", message)
	u := url.URL{
		Scheme:   "http",
		Host:     *host,
		Path:     "/hash",
		RawQuery: q.Encode(),
	}

	// Send message to server.
	resp, err := http.Post(u.String(), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the whole body.
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Check for failure.
	switch resp.StatusCode {
	case http.StatusOK:
		fmt.Print(string(b))
	case http.StatusConflict:
		log.Fatal("Sorry, that hash is too small.")
	default:
		log.Fatal("ERROR!", string(b))
	}
}
