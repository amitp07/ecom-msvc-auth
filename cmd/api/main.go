package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/amitp07/ecom-msvc-auth/internal/data"
	"github.com/amitp07/ecom-msvc-auth/internal/db"
	"github.com/amitp07/ecom-msvc-auth/internal/env"
)

type config struct {
	port int
	env  string
}

type application struct {
	cfg    *config
	logger *log.Logger
	models data.Models
}

func main() {

	db := db.InitDb()

	cfg := config{
		port: env.Env.Port,
		env:  env.Env.AppEnvironment,
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := application{
		cfg:    &cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", env.Env.Port),
		Handler: app.run(),
	}

	app.logger.Printf("%s server is starting on port %d ...\n", env.Env.AppEnvironment, app.cfg.port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
