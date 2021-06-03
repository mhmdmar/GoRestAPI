package Utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ServeEncodedJSONStruct(w http.ResponseWriter, obj interface{}) {
	err := json.NewEncoder(w).Encode(obj)
	if err != nil {
		fmt.Println(err)
	}
}
