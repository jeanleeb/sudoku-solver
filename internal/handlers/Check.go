package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func (s *Service) Check(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var errors [9][9]bool

	for row := range 9 {
		for col := range 9 {
			name := fmt.Sprintf("cell-%d-%d", row, col)
			valStr := r.FormValue(name)
			if s.original.Get(row, col) != 0 {
				continue
			}
			if s.original.Get(row, col) == 0 && valStr == "" {
				s.current.Set(row, col, 0)
				errors[row][col] = false
				continue
			}

			val, err := strconv.Atoi(valStr)
			if err != nil || val < 1 || val > 9 {
				errors[row][col] = true
				continue
			}
			s.current.Set(row, col, val)
		}
	}

	for row := range 9 {
		for col := range 9 {
			val := s.current.Get(row, col)
			if val == 0 {
				continue
			}
			s.current.Set(row, col, 0)
			if !s.current.IsValidPlacement(row, col, val) {
				errors[row][col] = true
			}
			s.current.Set(row, col, val)
		}
	}

	s.errors = errors

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
