package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/perryd01/fh6t/internal/server"
	"github.com/perryd01/fh6t/internal/session"
	"github.com/perryd01/fh6t/internal/telemetry"
)

func main() {
	udpAddr := getenv("UDP_PORT", ":5300")
	httpAddr := getenv("HTTP_PORT", ":8080")

	inbound := make(chan telemetry.Packet, 256)
	managerCh := make(chan telemetry.Packet, 256)
	hubCh := make(chan telemetry.Packet, 256)

	store := session.NewMemoryStore()
	manager := session.NewManager(store, managerCh)
	hub := server.NewSSEHub(hubCh)
	handler := server.New(store, hub)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go fanOut(ctx, inbound, managerCh, hubCh)
	go manager.Run(ctx)
	go hub.Run(ctx)
	go listenUDP(ctx, udpAddr, inbound)

	srv := &http.Server{Addr: httpAddr, Handler: handler}
	go func() {
		<-ctx.Done()
		srv.Shutdown(context.Background())
	}()

	log.Printf("HTTP listening on %s, UDP on %s", httpAddr, udpAddr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP server: %v", err)
	}
}

func fanOut(ctx context.Context, in <-chan telemetry.Packet, outs ...chan<- telemetry.Packet) {
	for {
		select {
		case <-ctx.Done():
			return
		case p, ok := <-in:
			if !ok {
				return
			}
			for _, out := range outs {
				select {
				case out <- p:
				default:
				}
			}
		}
	}
}

func listenUDP(ctx context.Context, addr string, out chan<- telemetry.Packet) {
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		log.Fatalf("UDP listen %s: %v", addr, err)
	}
	defer conn.Close()

	go func() {
		<-ctx.Done()
		conn.Close()
	}()

	buf := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFrom(buf)
		if err != nil {
			select {
			case <-ctx.Done():
				return
			default:
				log.Printf("UDP read: %v", err)
				continue
			}
		}
		p, err := telemetry.Decode(buf[:n])
		if err != nil {
			log.Printf("decode: %v", err)
			continue
		}
		select {
		case out <- p:
		default:
			log.Println("inbound channel full — dropping packet")
		}
	}
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
