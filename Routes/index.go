package Routes

import (
	"../Utils"
	"github.com/gorilla/mux"
	"net/http"
)

type Response struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func ApiNotFound(w http.ResponseWriter, _ *http.Request) {
	Utils.ServeEncodedJSONStruct(w, Response{Error: "API not found"})
}

func HandleRoutes(router *mux.Router, apiPrefix string) {
	HandleProductsRoutes(router, apiPrefix)
	HandleUserRoutes(router, apiPrefix)
	router.PathPrefix(apiPrefix).HandlerFunc(ApiNotFound)
}
