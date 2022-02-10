package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Denuha/anekdot-service/internal/config"
	delivery "github.com/Denuha/anekdot-service/internal/delivery/http"
	"github.com/Denuha/anekdot-service/internal/repository"
	"github.com/Denuha/anekdot-service/internal/repository/client"
	"github.com/Denuha/anekdot-service/internal/server"
	"github.com/Denuha/anekdot-service/internal/service"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
)

func Run() {
	log := logrus.New()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := sql.Open("pgx", cfg.DBConnStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	if err := goose.Run("up", db, "migrations", ""); err != nil {
		log.Fatal(err)
	}

	pgClient := client.NewPostgresClient(db)
	repos := repository.NewRepositories(pgClient)
	services := service.NewServices(cfg, repos)
	handlers := delivery.NewHandlers(services)

	srv := server.NewServer(*cfg, handlers.Init(cfg, log))
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error occurred while running http server: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Stop(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	fmt.Println("hello")
}
