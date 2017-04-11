package httpserv

import (
	"context"
	"encoding/json"
	"io/ioutil"
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
	var content map[string]string

	log.Printf("chopHandler: %s: %s", request.Host, request.Method)

	marshalled, err := ioutil.ReadAll(request.Body)
	request.Body.Close()
	if err != nil {
		log.Printf("error: unable to read body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(marshalled, &content); err != nil {
		log.Printf("error: unable to unmarshal body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("debug: url = %s", content["url"])
}

// unchopHandler redirects to the original URL if possible
func (s *server) unchopHandler(w http.ResponseWriter, request *http.Request) {

}
