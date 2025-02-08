package routes

import (
	"operator_text_channel/src/handlers"
	"operator_text_channel/src/services"

	"github.com/gorilla/mux"
)

func RegisterRoutes(service *services.Service) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/tags", handlers.CachingMiddleware(service, handlers.GetTags(service))).Methods("GET")
	r.HandleFunc("/tags", handlers.CreateTag(service)).Methods("POST")
	r.HandleFunc("/tags", handlers.DeleteTag(service)).Methods("DELETE")
	r.HandleFunc("/tags/name", handlers.GetTagName(service)).Methods("GET")

	r.HandleFunc("/appeals", handlers.GetAppeals(service)).Methods("GET")
	r.HandleFunc("/appeals", handlers.CreateAppeal(service)).Methods("POST")
	r.HandleFunc("/appeals", handlers.DeleteAppeal(service)).Methods("DELETE")
	r.HandleFunc("/appeals/tags", handlers.AddTagsToAppeal(service)).Methods("POST")
	r.HandleFunc("/appeals/tags", handlers.RemoveTagsFromAppeal(service)).Methods("DELETE")

	r.HandleFunc("/operators", handlers.CreateOperator(service)).Methods("POST")
	r.HandleFunc("/operators", handlers.DeleteOperator(service)).Methods("DELETE")
	r.HandleFunc("/operators/tags", handlers.GetOperatorTags(service)).Methods("GET")
	r.HandleFunc("/operators/appeals", handlers.GetOperatorAppeals(service)).Methods("GET")

	r.HandleFunc("/groups", handlers.CachingMiddleware(service, handlers.GetGroups(service))).Methods("GET")
	r.HandleFunc("/groups", handlers.CreateGroup(service)).Methods("POST")
	r.HandleFunc("/groups", handlers.DeleteGroup(service)).Methods("DELETE")
	r.HandleFunc("/groups/operators", handlers.GetGroupOperators(service)).Methods("GET")
	r.HandleFunc("/groups/operators", handlers.AddOperatorsToGroup(service)).Methods("POST")
	r.HandleFunc("/groups/operators", handlers.RemoveOperatorsFromGroup(service)).Methods("DELETE")
	r.HandleFunc("/groups/tags", handlers.GetGroupTags(service)).Methods("GET")
	r.HandleFunc("/groups/tags", handlers.AddTagsToGroup(service)).Methods("POST")
	r.HandleFunc("/groups/tags", handlers.RemoveTagsFromGroup(service)).Methods("DELETE")

	return r
}
