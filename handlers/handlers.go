package handlers

import (
	"net/http"

	"github.com/nkcyber/sql-injection-lab/components"
	"golang.org/x/exp/slog"
)

func New(log *slog.Logger) *DefaultHandler {
	return &DefaultHandler{
		Log: log,
	}
}

type DefaultHandler struct {
	Log *slog.Logger
}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	h.Get(w, r)
}

func (h *DefaultHandler) Get(w http.ResponseWriter, r *http.Request) {
	h.Log.Info("Default Handler Get")
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
