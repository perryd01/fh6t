package server

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"

	"github.com/perryd01/fh6t/internal/session"
)

func New(store session.Store, hub *SSEHub) http.Handler {
	router := chi.NewRouter()
	api := humachi.New(router, huma.DefaultConfig("fh6t API", "1.0.0"))

	huma.Register(api, huma.Operation{
		OperationID: "list-sessions",
		Method:      http.MethodGet,
		Path:        "/api/sessions",
		Summary:     "List all recorded lap sessions",
	}, listSessionsHandler(store))

	huma.Register(api, huma.Operation{
		OperationID: "get-session-packets",
		Method:      http.MethodGet,
		Path:        "/api/sessions/{id}/packets",
		Summary:     "Get all packets for a session",
	}, getSessionPacketsHandler(store))

	router.Get("/api/events", sseHandler(hub))

	return router
}
