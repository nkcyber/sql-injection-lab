package handlers

import (
	"fmt"
	"net/http"

	"github.com/nkcyber/sql-injection-lab/components"
	"github.com/nkcyber/sql-injection-lab/db"
	"golang.org/x/exp/slog"
)

func New(log *slog.Logger, db *db.Documents) *DefaultHandler {
	return &DefaultHandler{
		Log:       log,
		Documents: db,
	}
}

type DefaultHandler struct {
	Log       *slog.Logger
	Documents *db.Documents
}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	h.Log.Info(fmt.Sprintf("SERVE HTTP GET '%v'", r.URL))
	// TODO: handle routes that aren't the document viewer
	h.GetDocumentViewer(w, r)
}

func (h *DefaultHandler) GetDocumentViewer(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM documents WHERE securityCode = 'A';"
	ds, err := h.Documents.UnsafeQuery(query)
	if err != nil {
		h.Log.Error(fmt.Sprintf("error querying all documents: %v\n", err))
	}

	// render page to client
	components.DocumentViewer(query, ds).Render(r.Context(), w)
}
