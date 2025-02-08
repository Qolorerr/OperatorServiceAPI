package handlers

import (
	"net/http"
	"operator_text_channel/src/services"
)

func CachingMiddleware(service *services.Service, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cacheKey := r.URL.RequestURI()

		cachedData, err := service.RDB.Get(service.CTX, cacheKey).Result()
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(cachedData))
			return
		}

		next.ServeHTTP(w, r)
	}
}
