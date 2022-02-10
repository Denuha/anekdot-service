package app

import (
	"database/sql"
	"fmt"

	"github.com/Denuha/anekdot-service/internal/config"
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

	fmt.Println("hello")
}
