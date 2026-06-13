package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/perryd01/fh6t/internal/session"
	"github.com/perryd01/fh6t/internal/telemetry"
)

type SessionSummary struct {
	ID          string `json:"id"`
	StartedAt   string `json:"startedAt"`
	EndedAt     string `json:"endedAt,omitempty"`
	PacketCount int    `json:"packetCount"`
}
type ListSessionsInput struct{}
type ListSessionsOutput struct {
	Body []SessionSummary
}

func listSessionsHandler(store session.Store) func(context.Context, *ListSessionsInput) (*ListSessionsOutput, error) {
	return func(ctx context.Context, _ *ListSessionsInput) (*ListSessionsOutput, error) {
		sessions, err := store.List()
		if err != nil {
			return nil, err
		}

		sessionSummaries := make([]SessionSummary, 0)
		for _, s := range sessions {
			tempSummary := SessionSummary{
				ID:          s.ID,
				StartedAt:   s.StartedAt.Format(time.RFC3339),
				PacketCount: len(s.Packets),
			}

			if s.EndedAt != nil {
				tempSummary.EndedAt = s.EndedAt.Format(time.RFC3339)
			}

			sessionSummaries = append(sessionSummaries, tempSummary)
		}

		output := ListSessionsOutput{
			Body: sessionSummaries,
		}

		return &output, nil

	}
}

type GetSessionPacketsInput struct {
	ID string `path:"id"`
}
type GetSessionPacketsOutput struct {
	Body []telemetry.Packet
}

func getSessionPacketsHandler(store session.Store) func(context.Context, *GetSessionPacketsInput) (*GetSessionPacketsOutput, error) {
	return func(ctx context.Context, in *GetSessionPacketsInput) (*GetSessionPacketsOutput, error) {
		s, err := store.Get(in.ID)
		if err != nil {
			return nil, huma.Error404NotFound("session not found", err)
		}
		return &GetSessionPacketsOutput{Body: s.Packets}, nil
	}
}

func sseHandler(hub *SSEHub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "streaming unsupported", http.StatusInternalServerError)
			return
		}

		client := hub.Subscribe()
		defer hub.Unsubscribe(client)

		for {
			select {
			case <-r.Context().Done():
				return
			case data := <-client:
				fmt.Fprintf(w, "data: %s\n\n", data)
				flusher.Flush()
			}
		}
	}
}
