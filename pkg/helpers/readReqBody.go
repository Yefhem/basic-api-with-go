package helpers

import (
	"io/ioutil"
	"net/http"
)

func ReadReqBody(r *http.Request) (body []byte) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	CheckError(err)
	return body
}
