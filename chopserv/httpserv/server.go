package httpserv

import (
	"context"
	"log"
	"net/http"

	"github.com/dougfort/chopper/chopserv/types"
)

type server struct {
	ctx context.Context
	cfg types.Config
}

// Serve HTTP
// we could use a more sophisticated multiplexer like Gorilla, but the
// interface is so simple I think this will suffice
func Serve(
	ctx context.Context,
	cfg types.Config,
) {
	s := &server{
		ctx: ctx,
		cfg: cfg,
	}

	http.HandleFunc("/chop", s.chopHandler)
	http.HandleFunc("/unchop", s.unchopHandler)

	// ListenAndServe always returns an error
	err := http.ListenAndServe(cfg.Address, nil)
	log.Printf("debug: http.ListenAndServe returned: %s", err)
}

// chopHandler 'chops' the URL to a smaller size
func (s *server) chopHandler(w http.ResponseWriter, request *http.Request) {
	log.Printf("chopHandler: %s: %s", request.Host, request.Method)
}

// unchopHandler redirects to the original URL if possible
func (s *server) unchopHandler(w http.ResponseWriter, request *http.Request) {

}
