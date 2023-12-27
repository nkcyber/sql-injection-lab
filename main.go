package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/nkcyber/sql-injection-lab/db"
	"github.com/nkcyber/sql-injection-lab/handlers"
	"golang.org/x/exp/slog"
)

func main() {
	seedPath := flag.String("seedPath", "./example_seed.sql", "The path to the SQL script with seed data. The script will be executed on server initalization.")
	port := flag.Int("port", 8080, "the port to listen and serve HTTP on")
	flag.Parse()

	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db, err := db.NewDocuments(*seedPath)
	if err != nil {
		log.Error(fmt.Sprintf("could not initalize database: %v", err))
		return
	}
	h := handlers.New(log, db)

	server := &http.Server{
		Addr:         fmt.Sprintf("localhost:%d", *port),
		Handler:      h,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	fmt.Printf("Listening on %v\n", server.Addr)
	server.ListenAndServe()
}
