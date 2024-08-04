package initializers

import (
	"api/src/configs"
	"api/src/storages/psql"
	"fmt"
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var cfg *configs.PSQLConfig

func InitDBconnection() {
	cleanenv.ReadEnv(&cfg)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Yekaterinburg",
		cfg.Host_db,
		cfg.Db_user,
		cfg.Db_password,
		cfg.Db_name,
		cfg.Db_port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect database", "error", err)
		os.Exit(2)
	}
	psql.PSQL_GORM_DB = db
	slog.Info("connected to database")
}

func InitChiRouting() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome developer"))
	})	
	http.ListenAndServe(":8000", r)
}
