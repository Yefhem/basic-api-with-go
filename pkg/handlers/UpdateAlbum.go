package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Yefhem/basic-api-with-go/pkg/helpers"
	"github.com/Yefhem/basic-api-with-go/pkg/models"
)

func (h handler) PutAlbumHandler(rw http.ResponseWriter, r *http.Request) {

	id := helpers.ReadDynamicID(r) // Read dynamic id
	body := helpers.ReadReqBody(r) // Read request body

	var updateAlbum models.Album
	json.Unmarshal(body, &updateAlbum)

	var album models.Album

	if result := h.DB.First(&album, id); result.Error != nil {
		helpers.CheckError(result.Error)
	}

	album.Title = updateAlbum.Title
	album.Singer = updateAlbum.Singer
	album.Year = updateAlbum.Year

	h.DB.Save(&album)

	helpers.ResponseJSON(rw, 200, "Updated")
}
