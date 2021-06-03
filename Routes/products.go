package Routes

import (
	"../Utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Product struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Picture string `json:"picture"`
	Price   int    `json:"price"`
	Rating  int    `json:"rating"`
}

var products = []Product{
	{
		Id:      "1",
		Title:   "Product 1",
		Picture: "/images/product1.jpg",
		Price:   3000,
		Rating:  3,
	},
	{
		Id:      "2",
		Title:   "Product 2",
		Picture: "/images/product2.jpg",
		Price:   10,
		Rating:  1,
	},
}

func GetProductById(id string) *Product {
	for _, product := range products {
		if product.Id == id {
			return &product
		}
	}
	return nil
}

func ReturnAllProducts(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: ReturnAllProducts")
	Utils.ServeEncodedJSONStruct(w, Response{Data: products})
}

func ReturnProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: ReturnProduct")
	ids, ok := r.URL.Query()["id"]
	if !ok || ids == nil || len(ids[0]) < 1 {
		Utils.ServeEncodedJSONStruct(w, Response{Error: "missing id"})
	} else {
		id := ids[0]
		product := GetProductById(id)
		if product == nil {
			Utils.ServeEncodedJSONStruct(w, Response{Error: fmt.Sprintf("product with id %s doesn't exist", id)})
		} else {
			Utils.ServeEncodedJSONStruct(w, Response{Data: product})
		}
	}
}

func HandleProductsRoutes(router *mux.Router, apiPrefix string) {
	fmt.Println(Utils.ConcatAPIPath(apiPrefix, "products"))
	router.HandleFunc(Utils.ConcatAPIPath(apiPrefix, "products"), ReturnAllProducts)
	router.HandleFunc(Utils.ConcatAPIPath(apiPrefix, "product"), ReturnProduct)
}
