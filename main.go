package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/nkcyber/sql-injection-lab/db"
	"github.com/nkcyber/sql-injection-lab/handlers"
	"golang.org/x/exp/slog"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	db, err := db.NewDocuments()
	if err != nil {
		log.Error(fmt.Sprintf("could not initalize database: %v", err))
		return
	}
	h := handlers.New(log, db)

	server := &http.Server{
		Addr:         "localhost:9000",
		Handler:      h,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	fmt.Printf("Listening on %v\n", server.Addr)
	server.ListenAndServe()
}
