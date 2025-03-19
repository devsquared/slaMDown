package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/devsquared/slaMDown/formatter"
	"github.com/devsquared/slaMDown/util"
)

func NewServer(ctx context.Context) error { //TODO: just throw the debug flag into a context and other needed info. Then pass that here.
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", IndexHandler())

	debug, ok := util.GetDebugContextKey(ctx)
	if !ok {
		return fmt.Errorf("unable to get debug from context") //TODO: could likely create a structured error type
	}
	port, ok := util.GetPortContextKey(ctx)
	if !ok {
		return fmt.Errorf("unable to get port from context")
	}

	if debug {
		formatter.PrintDebug("Spinning up local server at port: " + port)
	}

	// quickly validate port
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	return http.ListenAndServe(port, mux)
}

func IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Well hello there!"))
	}
}
