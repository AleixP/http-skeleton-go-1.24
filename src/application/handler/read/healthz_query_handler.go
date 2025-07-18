package read

import (
	"encoding/json"
	"net/http"
)

type HealthzQueryHandler struct{}

func NewHealthzQueryHandler() *HealthzQueryHandler {
	return &HealthzQueryHandler{}
}
func (healthzQueryHandler *HealthzQueryHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	var response []string
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
