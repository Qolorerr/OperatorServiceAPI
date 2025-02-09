package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"operator_text_channel/src/services"
	"time"
)

func GetGroups(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groups, err := service.GetGroups()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(groups)
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

func CreateGroup(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestGroup requestCreateGroup
		err := json.NewDecoder(r.Body).Decode(&requestGroup)
		if err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		group, err := service.CreateGroup(requestGroup.TagIds)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(group)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResponse)
	}
}

func DeleteGroup(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing group id", http.StatusBadRequest)
			return
		}
		if err := uuid.Validate(id); err != nil {
			http.Error(w, "Invalid group id", http.StatusBadRequest)
			return
		}
		if err := service.DeleteGroup(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func GetGroupOperators(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing group id", http.StatusBadRequest)
			return
		}
		if err := uuid.Validate(id); err != nil {
			http.Error(w, "Invalid group id", http.StatusBadRequest)
			return
		}

		operators, err := service.GetGroupOperators(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(operators)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}

func AddOperatorsToGroup(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request requestChangeOperatorsInGroup
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}
		if err = service.AddOperatorsToGroup(request.GroupId, request.OperatorIds); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func RemoveOperatorsFromGroup(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request requestChangeOperatorsInGroup
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}
		if err = service.RemoveOperatorsFromGroup(request.GroupId, request.OperatorIds); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func GetGroupTags(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing group id", http.StatusBadRequest)
			return
		}
		if err := uuid.Validate(id); err != nil {
			http.Error(w, "Invalid group id", http.StatusBadRequest)
			return
		}

		tags, err := service.GetGroupTags(id)
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

func AddTagsToGroup(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request requestChangeTagsInGroup
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}
		if err = service.AddTagsToGroup(request.GroupId, request.TagIds); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func RemoveTagsFromGroup(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request requestChangeTagsInGroup
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}
		if err = service.RemoveTagsFromGroup(request.GroupId, request.TagIds); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
