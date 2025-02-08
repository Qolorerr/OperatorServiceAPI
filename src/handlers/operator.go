package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"operator_text_channel/src/services"
)

func CreateOperator(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		operator, err := service.CreateOperator()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(operator)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResponse)
	}
}

func DeleteOperator(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing operator id", http.StatusBadRequest)
			return
		}
		if err := uuid.Validate(id); err != nil {
			http.Error(w, "Invalid operator id", http.StatusBadRequest)
			return
		}

		err := service.DeleteOperator(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func GetOperatorTags(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing operator id", http.StatusBadRequest)
			return
		}
		if err := uuid.Validate(id); err != nil {
			http.Error(w, "Invalid operator id", http.StatusBadRequest)
			return
		}

		tags, err := service.GetOperatorTags(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(tags)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}

func GetOperatorAppeals(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing operator id", http.StatusBadRequest)
			return
		}
		if err := uuid.Validate(id); err != nil {
			http.Error(w, "Invalid operator id", http.StatusBadRequest)
			return
		}

		appeals, err := service.GetOperatorAppeals(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(appeals)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}
