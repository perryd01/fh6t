package server

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"

	"github.com/perryd01/fh6t/internal/session"
	"github.com/perryd01/fh6t/internal/telemetry"
)

// sseBody tricks Huma into documenting the /api/events endpoint with the
// correct content type and schema. The real handler is registered on chi
// because SSE streams don't fit Huma's request/response model.
type sseBody struct {
	telemetry.Packet
}

func (sseBody) ContentType(string) string { return "text/event-stream" }

type SSEOutput struct {
	Body sseBody
}

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

	huma.Register(api, huma.Operation{
		OperationID: "stream-events",
		Method:      http.MethodGet,
		Path:        "/api/events",
		Summary:     "Stream live telemetry packets as server-sent events",
	}, func(_ context.Context, _ *struct{}) (*SSEOutput, error) { return nil, nil })

	router.Get("/api/events", sseHandler(hub))

	return router
}
