package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	userIdChan    = make(chan int)
	postIdChan    = make(chan int)
	commentIdChan = make(chan int)
)

type server struct {
	addr string
}

func newServer(addr string) *server {
	return &server{
		addr: addr,
	}
}

func (s *server) start() {
	go startIdGeneration(userIdChan, 100)
	go startIdGeneration(postIdChan, 1000)
	go startIdGeneration(commentIdChan, 10000)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.Timeout(10 * time.Second))

	r.Route("/users", func(r chi.Router) {
		r.Get("/", listAllUsers)
		r.Post("/", createUser)

		r.Route("/{userId}", func(r chi.Router) {
			r.Get("/", getUser)
			r.Put("/", updateUser)
			r.Patch("/", editUser)
			r.Delete("/", deleteUser)

			// nested route
			r.Get("/posts", getUserPosts)
		})
	})

	r.Route("/posts", func(r chi.Router) {
		r.Get("/", listAllPosts)
		r.Post("/", createPost)

		r.Route("/{postId}", func(r chi.Router) {
			r.Get("/", getPost)
			r.Put("/", updatePost)
			r.Patch("/", editPost)
			r.Delete("/", deletePost)

			// nested route
			r.Get("/comments", getPostComments)
		})
	})

	r.Route("/comments", func(r chi.Router) {
		r.Get("/", listAllComments)
		r.Post("/", createComment)

		r.Route("/{commentId}", func(r chi.Router) {
			// r.Use(commentCtx)
			r.Get("/", getComment)
			r.Put("/", updateComment)
			r.Patch("/", editComment)
			r.Delete("/", deleteComment)
		})
	})

	log.Fatal(http.ListenAndServe(s.addr, r))
}

func writeJSON(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(v)
}

// can also be simply "any" instead of the types
func readJSON[T *User | *UserPatch | *Post | *PostPatch | *Comment | *CommentPatch](r io.Reader, c T) error {
	err := json.NewDecoder(r).Decode(c)
	return err
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func startIdGeneration(ch chan<- int, i int) {
	for {
		i += 1
		ch <- i
	}
}
