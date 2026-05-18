# Sudoku Solver

A simple Sudoku web-app written in Go, with [templ](https://templ.guide/) for HTML templates and rendering, [Tailwind CSS](https://tailwindcss.com/) for styles.

## Features

- Sudoku board generation
- Puzzle validation and response checking
- Automatic solving

## Running

Requires [Go](https://go.dev/) 1.26+.

```bash
git clone https://codeberg.com/org/sudoku-solver.git
cd sudoku-solver
go run ./cmd/server
```

HTTP server will be started in port 8080. Open `http://localhost:8080` in your browser to play.

## Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handlers/
│   │   ├── Check.go
│   │   ├── Home.go
│   │   ├── New.go
│   │   ├── routes.go
│   │   └── Solve.go
│   └── views/
│       ├── board.templ
│       ├── board_templ.go
│       ├── layout.templ
│       └── layout_templ.go
├── static/
│   ├── solve.js
│   └── styles.css
├── sudoku/
│   ├── board.go
│   ├── board_test.go
│   ├── generator.go
│   ├── generator_test.go
│   ├── solver.go
│   └── solver_test.go
├── go.mod
└── go.sum
```

## License

MIT
