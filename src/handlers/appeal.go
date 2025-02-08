package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"operator_text_channel/src/services"
)

func GetAppeals(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("userId")
		if userId == "" {
			http.Error(w, "Missing appeal user id", http.StatusBadRequest)
			return
		}
		if err := uuid.Validate(userId); err != nil {
			http.Error(w, "Invalid appeal user id", http.StatusBadRequest)
			return
		}

		appeals, err := service.GetAppealsByUserId(userId)
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

func CreateAppeal(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestAppeal requestCreateAppeal
		err := json.NewDecoder(r.Body).Decode(&requestAppeal)
		if err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		appeal, err := service.CreateAppeal(requestAppeal.UserId, requestAppeal.TagIds)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(appeal)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResponse)
	}
}

func DeleteAppeal(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing appeal id", http.StatusBadRequest)
			return
		}
		if err := uuid.Validate(id); err != nil {
			http.Error(w, "Invalid appeal id", http.StatusBadRequest)
			return
		}

		err := service.DeleteAppeal(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func AddTagsToAppeal(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request requestChangeTagsInAppeal
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		err = service.AddTagsToAppeal(request.AppealId, request.TagIds)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func RemoveTagsFromAppeal(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request requestChangeTagsInAppeal
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		err = service.RemoveTagsFromAppeal(request.AppealId, request.TagIds)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
