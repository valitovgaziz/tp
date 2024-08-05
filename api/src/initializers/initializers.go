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
	"gorm.io/gorm/logger"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var cfg *configs.PSQLConfig

func InitDBconnection() {
	slog.Info("Init DB connection")
	cleanenv.ReadEnv(&cfg)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Yekaterinburg",
		cfg.Host_db,
		cfg.Db_user,
		cfg.Db_password,
		cfg.Db_name,
		cfg.Db_port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		slog.Error("failed to connect database", "error", err)
		os.Exit(2)
	}
	psql.PSQL_GORM_DB = db
	sql, err := db.DB()
	if err != nil {
		slog.Error("failed to get database", "error", err)
		os.Exit(2)
	}
	err = sql.Ping()
	if err != nil {
		slog.Error("failed to ping database", "error", err)
		os.Exit(2)
	}
	slog.Info("connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)
}

func InitChiRouting() {
	slog.Info("Init routing")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome developer! Cool. Are you debugging?"))
	})
	err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r)
	if err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(2)
	}
	slog.Info("Server start on PORT: " + os.Getenv("SERVER_PORT"))
}
