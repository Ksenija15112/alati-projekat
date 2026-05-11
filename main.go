package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"alati-projekat/handler"
)

const serverAddress = ":8080"

func main() {
	mux := http.NewServeMux()
	registerRoutes(mux)

	server := &http.Server{
		Addr:         serverAddress,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go startServer(server)
	waitForShutdown(server)
}

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/ready", readyHandler)
	mux.HandleFunc("/uuid", uuidHandler)

	mux.HandleFunc("/configs", handler.ConfigsHandler)
	mux.HandleFunc("/configs/", handler.ConfigByIDHandler)

	mux.HandleFunc("/config-groups", handler.ConfigGroupsHandler)
	mux.HandleFunc("/config-groups/", handler.ConfigGroupByIDHandler)
}

func startServer(server *http.Server) {
	log.Println("Server pokrenut na http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("Greska pri pokretanju servera:", err)
	}
}

func waitForShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("Server se gasi...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal("Greska pri gasenju servera:", err)
	}
	log.Println("Server je uspesno ugasen.")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"message": "Backend server radi"})
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "READY"})
}

func uuidHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id, err := generateUUID()
	if err != nil {
		http.Error(w, "Failed to generate UUID", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"uuid": id})
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("Greska pri pisanju JSON odgovora:", err)
	}
}

func generateUUID() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	bytes[6] = (bytes[6] & 0x0f) | 0x40
	bytes[8] = (bytes[8] & 0x3f) | 0x80
	return hex.EncodeToString(bytes[0:4]) + "-" +
		hex.EncodeToString(bytes[4:6]) + "-" +
		hex.EncodeToString(bytes[6:8]) + "-" +
		hex.EncodeToString(bytes[8:10]) + "-" +
		hex.EncodeToString(bytes[10:16]), nil
}

// da ne bude unused import
var _ = strings.TrimPrefix
