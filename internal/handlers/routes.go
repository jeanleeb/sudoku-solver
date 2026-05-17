package handlers

import (
	"io"
	"jeanleeb/sudoku-solver/sudoku"
	"net/http"
)

type Service struct {
	current *sudoku.Board
}

func Register(mux *http.ServeMux, svc *Service) {
	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "OK")
	})
	mux.HandleFunc("GET /", svc.Home)
	mux.HandleFunc("GET /new", svc.New)
	mux.HandleFunc("POST /check", svc.Check)
}
