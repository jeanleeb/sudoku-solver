package handlers

import (
	"jeanleeb/sudoku-solver/sudoku"
	"net/http"
)

func (s *Service) New(w http.ResponseWriter, r *http.Request) {
	puzzle, _ := sudoku.Generate()
	s.current = puzzle
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
