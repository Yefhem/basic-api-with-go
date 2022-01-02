package handlers

import (
	"net/http"

	"github.com/Yefhem/basic-api-with-go/pkg/helpers"
	"github.com/Yefhem/basic-api-with-go/pkg/models"
)

func (h handler) GetAlbumHandler(rw http.ResponseWriter, r *http.Request) {

	id := helpers.ReadDynamicID(r) // Read dynamic id

	var album models.Album

	if result := h.DB.First(&album, id); result.Error != nil {
		helpers.CheckError(result.Error)
	}

	helpers.ResponseJSON(rw, 200, album)
}
