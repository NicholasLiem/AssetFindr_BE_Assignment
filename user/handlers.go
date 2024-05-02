package user

import (
	"encoding/json"
	"net/http"
)

func HelloHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(rw).Encode("Hello")
	if err != nil {
		return
	}
}
