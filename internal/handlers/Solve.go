package handlers

import (
	"encoding/json"
	"fmt"
	"jeanleeb/sudoku-solver/sudoku"
	"net/http"
)

func (s *Service) Solve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	fmt.Println("Starting to solve the puzzle...")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming not supported", http.StatusInternalServerError)
		return
	}

	board := s.current.Copy()

	ch := sudoku.SolveStepped(board)

	for step := range ch {
		data, err := json.Marshal(step)
		if err != nil {
			fmt.Printf("Error marshalling step: %v\n", err)
			continue
		}
		fmt.Fprintf(w, "event: step\ndata: %s\n\n", data)
		flusher.Flush()
	}

	fmt.Fprintf(w, "event: done\ndata: {}\n\n")
	flusher.Flush()

	s.current = board
}
