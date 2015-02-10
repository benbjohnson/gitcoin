package gitcoin

import (
	"fmt"
	"net/http"
)

// Handler represents an HTTP handler for a gitcoin Target.
type Handler struct {
	Target *Target
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/target":
		if r.Method == "GET" {
			h.handleTarget(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	case "/gitcoins":
		panic("not yet implemented")
	case "/hash":
		if r.Method == "POST" {
			h.handleHash(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	default:
		http.NotFound(w, r)
	}
}

func (h *Handler) handleTarget(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%x\n", h.Target.Value)
}

func (h *Handler) handleHash(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	message := q.Get("message")
	// owner := q.Get("owner")

	// Check if the message was passed.
	if message == "" {
		http.Error(w, "message required", http.StatusBadRequest)
		return
	}

	// If the check failed then return an error.
	if !h.Target.Check(message) {
		http.Error(w, "hash too large", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%x\n", h.Target.Value)
}
