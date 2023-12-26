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
	h.Get(w, r)
}

func (h *DefaultHandler) Get(w http.ResponseWriter, r *http.Request) {
	h.Log.Info("Default Handler Get")
	ds, err := h.Documents.QueryAll()
	if err != nil {
		h.Log.Error(fmt.Sprintf("error querying all documents: %v\n", err))
	}
	for _, document := range ds {
		h.Log.Info(fmt.Sprintf("Document:\n\tName: %v\n\tCode: %v\n\tContent: %v\n", document.Name, document.SecurityCode, document.Content))
	}
	h.Log.Info(fmt.Sprintf("# of documents: %v", len(ds)))
	// var props ViewProps
	// var err error
	// if err != nil {
	// 	h.Log.Error("failed to get counts", slog.Any("error", err))
	// 	http.Error(w, "failed to get counts", http.StatusInternalServerError)
	// 	return
	// }
	h.View(w, r)
}

// func (h *DefaultHandler) Post(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()

// 	// Decide the action to take based on the button that was pressed.
// 	var it services.IncrementType
// 	if r.Form.Has("global") {
// 		it = services.IncrementTypeGlobal
// 	}
// 	if r.Form.Has("session") {
// 		it = services.IncrementTypeSession
// 	}

// 	counts, err := h.CountService.Increment(r.Context(), it, session.ID(r))
// 	if err != nil {
// 		h.Log.Error("failed to increment", slog.Any("error", err))
// 		http.Error(w, "failed to increment", http.StatusInternalServerError)
// 		return
// 	}

// 	// Display the view.
// 	h.View(w, r, ViewProps{
// 		Counts: counts,
// 	})
// }

// type ViewProps struct {
// 	Counts services.Counts
// }

func (h *DefaultHandler) View(w http.ResponseWriter, r *http.Request) {
	components.Index().Render(r.Context(), w)
}