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
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/ready", readyHandler)
	mux.HandleFunc("/uuid", uuidHandler)

	// Placeholder rute za API koji ce kasnije biti implementiran.
	mux.HandleFunc("/configs", configsHandler)
	mux.HandleFunc("/configs/", configByIDHandler)

	// Placeholder rute za konfiguracione grupe.
	mux.HandleFunc("/config-groups", configGroupsHandler)
	mux.HandleFunc("/config-groups/", configGroupByIDHandler)
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

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "Backend server radi",
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"status": "OK",
	})
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"status": "READY",
	})
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

	writeJSON(w, http.StatusOK, map[string]string{
		"uuid": id,
	})
}

func configsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusNotImplemented, map[string]string{
			"message": "GET /configs ruta postoji, API implementacija ceka",
		})

	case http.MethodPost:
		writeJSON(w, http.StatusNotImplemented, map[string]string{
			"message": "POST /configs ruta postoji, API implementacija ceka",
		})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func configByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/configs/")

	if id == "" {
		http.Error(w, "Missing config id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusNotImplemented, map[string]string{
			"message": "GET /configs/{id} ruta postoji, API implementacija ceka",
			"id":      id,
		})

	case http.MethodPut:
		writeJSON(w, http.StatusNotImplemented, map[string]string{
			"message": "PUT /configs/{id} ruta postoji, API implementacija ceka",
			"id":      id,
		})

	case http.MethodDelete:
		writeJSON(w, http.StatusNotImplemented, map[string]string{
			"message": "DELETE /configs/{id} ruta postoji, API implementacija ceka",
			"id":      id,
		})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func configGroupsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusNotImplemented, map[string]string{
			"message": "GET /config-groups ruta postoji, API implementacija ceka",
		})

	case http.MethodPost:
		writeJSON(w, http.StatusNotImplemented, map[string]string{
			"message": "POST /config-groups ruta postoji, API implementacija ceka",
		})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func configGroupByIDHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/config-groups/")
	path = strings.Trim(path, "/")

	if path == "" {
		http.Error(w, "Missing config group id", http.StatusBadRequest)
		return
	}

	parts := strings.Split(path, "/")
	groupID := parts[0]

	if len(parts) == 1 {
		handleSingleConfigGroup(w, r, groupID)
		return
	}

	if len(parts) == 2 && parts[1] == "configs" {
		handleConfigGroupConfigs(w, r, groupID)
		return
	}

	if len(parts) == 3 && parts[1] == "configs" {
		configID := parts[2]
		handleConfigInGroup(w, r, groupID, configID)
		return
	}

	http.NotFound(w, r)
}

func handleSingleConfigGroup(w http.ResponseWriter, r *http.Request, groupID string) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusNotImplemented, map[string]string{
			"message":  "GET /config-groups/{id} ruta postoji, API implementacija ceka",
			"group_id": groupID,
		})

	case http.MethodPut:
		writeJSON(w, http.StatusNotImplemented, map[string]string{
			"message":  "PUT /config-groups/{id} ruta postoji, API implementacija ceka",
			"group_id": groupID,
		})

	case http.MethodDelete:
		writeJSON(w, http.StatusNotImplemented, map[string]string{
			"message":  "DELETE /config-groups/{id} ruta postoji, API implementacija ceka",
			"group_id": groupID,
		})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleConfigGroupConfigs(w http.ResponseWriter, r *http.Request, groupID string) {
	switch r.Method {
	case http.MethodPost:
		writeJSON(w, http.StatusNotImplemented, map[string]string{
			"message":  "POST /config-groups/{id}/configs ruta postoji za dodavanje konfiguracije u grupu",
			"group_id": groupID,
		})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleConfigInGroup(w http.ResponseWriter, r *http.Request, groupID string, configID string) {
	switch r.Method {
	case http.MethodDelete:
		writeJSON(w, http.StatusNotImplemented, map[string]string{
			"message":   "DELETE /config-groups/{id}/configs/{configId} ruta postoji za uklanjanje konfiguracije iz grupe",
			"group_id":  groupID,
			"config_id": configID,
		})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
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
