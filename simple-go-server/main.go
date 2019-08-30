package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()
	handler := newHTTPHandler()
	srv := &http.Server{
		Addr:    ":8100",
		Handler: handler,
	}
	fmt.Println("starting new http server on :8100")
	err := runServer(ctx, srv, srv.ListenAndServe)
	if err != nil {
		log.Fatal(err)
	}
}

func runServer(ctx context.Context, srv *http.Server, f func() error) error {
	errCh := make(chan error)
	go func() {
		select {
		case errCh <- f():
		case <-ctx.Done():
		}
	}()
	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		err := srv.Shutdown(context.Background())
		return err
	}
}

func newHTTPHandler() http.Handler {
	r := mux.NewRouter()
	r.NewRoute().
		Methods(http.MethodGet).
		Path("/").
		Handler(&httpHandler{})
	return r
}

type httpHandler struct{}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := h.handle(w, req)
	if err == nil {
		return
	}
	log.Fatal(err)
}

func (h *httpHandler) handle(w http.ResponseWriter, req *http.Request) error {
	resp, err := json.Marshal(`{request:"success"}`)
	if err != nil {
		return err
	}
	fmt.Println("req headers:", req.Header)
	w.Header().Set("Content-Type", "application/json")

	//Ignore write error,we don't care
	_, _ = w.Write(resp)
	return nil
}
