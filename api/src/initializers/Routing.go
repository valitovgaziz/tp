package initializers

import (
	"api/src/admin"
	"api/src/auth"
	"log/slog"
	"os"
	"time"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var Done = make(chan bool)

func InitChiRouting() {
	slog.Info("Init routing")
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.RequestID)
	r.Use(middleware.CleanPath)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.NoCache)
	r.Use(middleware.Recoverer)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})

	// public Routes
	r.Group(func(r chi.Router) {
		r.Post("/signup", auth.Register) // register
		r.Post("/signin", auth.Login)     // signin
	})

	// Private Routes
	// Require Authentication
	r.Group(func(r chi.Router) {
		r.Use(auth.AuthMiddleware)
		r.Get("/allUsers", admin.GetAllUser) // all users get
	})

	// up server on os.Getenv("SERVER_PORT") port on gorutin
	go func() {
		defer close(Done)
		err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r)
		if err != nil {
			slog.Error("Can't start server: ", "error", err)
		}
	}()
}
