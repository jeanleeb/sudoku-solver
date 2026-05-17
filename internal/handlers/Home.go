package handlers

import (
	"jeanleeb/sudoku-solver/internal/views"
	"net/http"
)

func (s *Service) Home(w http.ResponseWriter, r *http.Request) {
	if s.original == nil || s.current == nil {
		http.Redirect(w, r, "/new", http.StatusSeeOther)
		return
	}

	component := views.BoardView(s.original, s.current, s.errors)
	views.Layout("Sudoku", component).Render(r.Context(), w)
}
