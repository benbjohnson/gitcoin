package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/benbjohnson/gitcoin"
)

func main() {
	addr := flag.String("addr", ":10000", "bind address")
	flag.Parse()

	target := gitcoin.NewTarget()

	h := &gitcoin.Handler{Target: target}
	log.Printf("listening on http://localhost%s", *addr)
	log.Fatal(http.ListenAndServe(*addr, h))
}
