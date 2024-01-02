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
	ip := flag.String("ip", "localhost", "The ip address to listen and serve HTTP on")
	port := flag.Int("port", 8080, "The port to listen and serve HTTP on")
	seedPath := flag.String("seedPath", "./example_seed.sql", "The path to the SQL script with seed data; The script will be executed on server initalization")
	flag.Parse()

	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db, err := db.NewDocuments(*seedPath)
	if err != nil {
		log.Error(fmt.Sprintf("could not initalize database: %v", err))
		return
	}
	h := handlers.New(log, db)

	server := &http.Server{
		Addr:         fmt.Sprintf("%v:%d", *ip, *port),
		Handler:      h,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	fmt.Printf("Listening on %v\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Error(fmt.Sprintf("exited: %v", err))
	}
}
