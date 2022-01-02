package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Yefhem/basic-api-with-go/pkg/helpers"
	"github.com/Yefhem/basic-api-with-go/pkg/models"
)

func (h handler) AddAlbumHandler(rw http.ResponseWriter, r *http.Request) {

	body := helpers.ReadReqBody(r) // Read request body

	var album models.Album
	json.Unmarshal(body, &album)

	if result := h.DB.Create(&album); result.Error != nil {
		helpers.CheckError(result.Error)
	}

	helpers.ResponseJSON(rw, 201, "Created")
}
