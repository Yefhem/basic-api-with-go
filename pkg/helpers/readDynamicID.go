package helpers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ReadDynamicID(r *http.Request) (id int) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	CheckError(err)
	return id
}
