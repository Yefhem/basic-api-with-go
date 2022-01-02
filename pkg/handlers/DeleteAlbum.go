package handlers

import (
	"net/http"

	"github.com/Yefhem/basic-api-with-go/pkg/helpers"
	"github.com/Yefhem/basic-api-with-go/pkg/models"
)

func (h handler) DeleteAlbumHandler(rw http.ResponseWriter, r *http.Request) {

	id := helpers.ReadDynamicID(r) // Read dynamic id

	var album models.Album

	if result := h.DB.First(&album, id); result.Error != nil {
		helpers.CheckError(result.Error)
	}

	h.DB.Delete(&album)

	helpers.ResponseJSON(rw, 200, "Deleted")
}
