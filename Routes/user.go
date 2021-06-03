package Routes

import (
	"../Utils"
	"fmt"
	"github.com/gorilla/mux"
)

type User struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Picture string `json:"picture"`
	Price   int    `json:"price"`
	Rating  int    `json:"rating"`
}

func HandleUserRoutes(router *mux.Router, apiPrefix string) {
	fmt.Println(Utils.ConcatAPIPath(apiPrefix, "user"))
	router.HandleFunc(Utils.ConcatAPIPath(apiPrefix, "user"), ReturnAllProducts)
}
