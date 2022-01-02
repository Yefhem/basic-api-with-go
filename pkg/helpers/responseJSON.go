package helpers

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(rw http.ResponseWriter, code int, message interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(code)
	json.NewEncoder(rw).Encode(message)
}
