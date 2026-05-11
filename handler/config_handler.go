package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"alati-projekat/models"
	"alati-projekat/store"
)

var s = &store.Store{}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// ============ HEALTH ============

func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "OK"})
}

// ============ CONFIGS ============

func ConfigsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createConfig(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func ConfigByIDHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/configs/")
	parts := strings.Split(path, "/")

	if len(parts) < 2 {
		http.Error(w, "Potrebni su naziv i verzija", http.StatusBadRequest)
		return
	}

	naziv := parts[0]
	verzija := parts[1]

	switch r.Method {
	case http.MethodGet:
		getConfig(w, r, naziv, verzija)
	case http.MethodDelete:
		deleteConfig(w, r, naziv, verzija)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func createConfig(w http.ResponseWriter, r *http.Request) {
	var c models.Config
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Neispravan JSON", http.StatusBadRequest)
		return
	}
	if c.Naziv == "" || c.Verzija == "" {
		http.Error(w, "Naziv i verzija su obavezni", http.StatusBadRequest)
		return
	}
	err := s.DodavanjeConfig(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	writeJSON(w, http.StatusCreated, c)
}

func getConfig(w http.ResponseWriter, r *http.Request, naziv, verzija string) {
	c, err := s.FindByNazivIVerzija(naziv, verzija)
	if err != nil {
		http.Error(w, "Konfiguracija nije pronađena", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, c)
}

func deleteConfig(w http.ResponseWriter, r *http.Request, naziv, verzija string) {
	c, err := s.FindByNazivIVerzija(naziv, verzija)
	if err != nil {
		http.Error(w, "Konfiguracija nije pronađena", http.StatusNotFound)
		return
	}
	err = s.Delete(c.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// ============ GROUPS ============

func ConfigGroupsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createGroup(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func ConfigGroupByIDHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/config-groups/")
	path = strings.Trim(path, "/")
	parts := strings.Split(path, "/")

	if len(parts) == 0 || parts[0] == "" {
		http.Error(w, "Nedostaje naziv grupe", http.StatusBadRequest)
		return
	}

	// /config-groups/{naziv}/{verzija}
	if len(parts) == 2 {
		naziv := parts[0]
		verzija := parts[1]
		switch r.Method {
		case http.MethodGet:
			getGroup(w, r, naziv, verzija)
		case http.MethodDelete:
			deleteGroup(w, r, naziv, verzija)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	// /config-groups/{naziv}/{verzija}/configs
	if len(parts) == 3 && parts[2] == "configs" {
		naziv := parts[0]
		verzija := parts[1]
		switch r.Method {
		case http.MethodPut:
			addConfigToGroup(w, r, naziv, verzija)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	// /config-groups/{naziv}/{verzija}/configs/{configNaziv}
	if len(parts) == 4 && parts[2] == "configs" {
		naziv := parts[0]
		verzija := parts[1]
		configNaziv := parts[3]
		switch r.Method {
		case http.MethodDelete:
			removeConfigFromGroup(w, r, naziv, verzija, configNaziv)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	http.NotFound(w, r)
}

func createGroup(w http.ResponseWriter, r *http.Request) {
	var g models.KonfiguracionaGrupa
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		http.Error(w, "Neispravan JSON", http.StatusBadRequest)
		return
	}
	if g.Naziv == "" || g.Verzija == "" {
		http.Error(w, "Naziv i verzija su obavezni", http.StatusBadRequest)
		return
	}
	err := s.DodavanjeGrupe(g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	writeJSON(w, http.StatusCreated, g)
}

func getGroup(w http.ResponseWriter, r *http.Request, naziv, verzija string) {
	g, err := s.FindGrupaByNazivIVerzija(naziv, verzija)
	if err != nil {
		http.Error(w, "Grupa nije pronađena", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, g)
}

func deleteGroup(w http.ResponseWriter, r *http.Request, naziv, verzija string) {
	g, err := s.FindGrupaByNazivIVerzija(naziv, verzija)
	if err != nil {
		http.Error(w, "Grupa nije pronađena", http.StatusNotFound)
		return
	}
	err = s.DeleteGrupa(g.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func addConfigToGroup(w http.ResponseWriter, r *http.Request, naziv, verzija string) {
	var c models.Config
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Neispravan JSON", http.StatusBadRequest)
		return
	}
	err := s.DodavanjeConfigUGrupu(naziv, verzija, c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func removeConfigFromGroup(w http.ResponseWriter, r *http.Request, naziv, verzija, configNaziv string) {
	err := s.UklanjranjeConfigIzGrupe(naziv, verzija, configNaziv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
