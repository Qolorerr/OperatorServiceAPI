package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"operator_text_channel/src/services"
	"strconv"
	"time"
)

func GetTags(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limitStr := r.URL.Query().Get("limit")
		offsetStr := r.URL.Query().Get("offset")

		limit := -1
		offset := 0

		if limitStr != "" {
			l, err := strconv.Atoi(limitStr)
			if err != nil {
				http.Error(w, "Invalid pagination", http.StatusBadRequest)
				return
			}
			limit = l
		}
		if offsetStr != "" {
			o, err := strconv.Atoi(offsetStr)
			if err != nil {
				http.Error(w, "Invalid pagination", http.StatusBadRequest)
				return
			}
			offset = o
		}

		tags, err := service.GetTags(limit, offset)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(tags)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cacheKey := r.URL.RequestURI()
		err = service.RDB.Set(service.CTX, cacheKey, jsonResponse, time.Minute*5).Err()
		if err != nil {
			log.Println("Error caching response:", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}

func GetTagName(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing tag id", http.StatusBadRequest)
			return
		}
		if err := uuid.Validate(id); err != nil {
			http.Error(w, "Invalid tag id", http.StatusBadRequest)
			return
		}

		tag, err := service.GetTagById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(map[string]interface{}{"name": tag.Name})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}

func CreateTag(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestTag requestCreateTag
		err := json.NewDecoder(r.Body).Decode(&requestTag)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		tag, err := service.CreateTag(requestTag.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(*tag)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResponse)
	}
}

func DeleteTag(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing tag id", http.StatusBadRequest)
			return
		}
		if err := uuid.Validate(id); err != nil {
			http.Error(w, "Invalid tag id", http.StatusBadRequest)
			return
		}

		err := service.DeleteTag(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
