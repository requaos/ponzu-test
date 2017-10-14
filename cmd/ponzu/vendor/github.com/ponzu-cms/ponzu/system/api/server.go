// Package api sets the various API handlers which provide an HTTP interface to
// Ponzu content, and include the types and interfaces to enable client-side
// interactivity with the system.
package api

import (
	"net/http"

	"github.com/requaos/access"
)

// Run adds Handlers to default http listener for API
func Run() {
	http.HandleFunc("/api/contents", access.GateKeeper(Record(CORS(Gzip(contentsHandler)))))

	http.HandleFunc("/api/content", access.GateKeeper(Record(CORS(Gzip(contentHandler)))))

	http.HandleFunc("/api/content/create", Record(CORS(createContentHandler)))

	http.HandleFunc("/api/content/update", access.GateKeeper(Record(CORS(updateContentHandler))))

	http.HandleFunc("/api/content/delete", access.GateKeeper(Record(CORS(deleteContentHandler))))

	http.HandleFunc("/api/search", access.GateKeeper(Record(CORS(Gzip(searchContentHandler)))))

	http.HandleFunc("/api/uploads", access.GateKeeper(Record(CORS(Gzip(uploadsHandler)))))
}
