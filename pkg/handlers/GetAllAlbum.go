package handlers

import (
	"net/http"

	"github.com/Yefhem/basic-api-with-go/pkg/helpers"
	"github.com/Yefhem/basic-api-with-go/pkg/models"
)

func (h handler) GetAllAlbumHanlder(rw http.ResponseWriter, r *http.Request) {

	var albums []models.Album

	if result := h.DB.Find(&albums); result.Error != nil {
		helpers.CheckError(result.Error)
	}

	helpers.ResponseJSON(rw, 200, albums)
}
