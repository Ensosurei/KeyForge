package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/tu-usuario/keyforge/backend/pkg/generator"
	"github.com/tu-usuario/keyforge/backend/pkg/verifier"
)

type VerifyRequest struct {
	Password string `json:"password"`
}

type PassphraseRequest struct {
	WordCount int    `json:"wordCount"`
	Separator string `json:"separator"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	path := r.URL.Path

	// Endpoint /api/verify
	if strings.HasSuffix(path, "/verify") {
		if r.Method != http.MethodPost {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		var req VerifyRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Cuerpo inválido", http.StatusBadRequest)
			return
		}

		result := verifier.VerifyPassword(req.Password)
		json.NewEncoder(w).Encode(result)
		return
	}

	// Endpoint /api/generate
	if strings.HasSuffix(path, "/generate") {
		if r.Method != http.MethodPost {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		var opts generator.GeneratorOptions
		if err := json.NewDecoder(r.Body).Decode(&opts); err != nil {
			http.Error(w, "Cuerpo inválido", http.StatusBadRequest)
			return
		}

		pass := generator.GeneratePassword(opts)
		json.NewEncoder(w).Encode(map[string]string{"password": pass})
		return
	}

	// Endpoint /api/passphrase
	// Busca la sección del endpoint /api/passphrase en api/index.go y actualízala así:
	if strings.HasSuffix(path, "/passphrase") {
		if r.Method != http.MethodPost {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		var opts generator.PassphraseOptions
		if err := json.NewDecoder(r.Body).Decode(&opts); err != nil {
			http.Error(w, "Cuerpo inválido", http.StatusBadRequest)
			return
		}

		phrase := generator.GeneratePassphrase(opts)
		json.NewEncoder(w).Encode(map[string]string{"passphrase": phrase})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"name":    "KeyForge API",
		"status":  "operational",
		"version": "1.0.0",
	})
}
