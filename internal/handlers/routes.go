package handlers

import (
	"io"
	"jeanleeb/sudoku-solver/sudoku"
	"net/http"
)

type Service struct {
	original *sudoku.Board
	current  *sudoku.Board
	errors   [9][9]bool
	isSolved bool
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
	mux.HandleFunc("GET /solve/stream", svc.Solve)
	mux.Handle("GET /static/", http.StripPrefix("/static", http.FileServer((http.Dir("static")))))
}
