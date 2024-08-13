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

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome developer! Cool."))
	})

	r.Post("/signin", auth.Register)
	r.Get("/allusers", admin.GetAllUser)

	r.Route("/auth", func(r chi.Router) {
		r.Route("/admin", func(r chi.Router) {
			r.Get("/allUsers", admin.GetAllUser)
		})
		r.Post("/login", auth.Login)
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
