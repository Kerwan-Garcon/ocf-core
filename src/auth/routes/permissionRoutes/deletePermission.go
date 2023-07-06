package permissionController

import (
	"net/http"

	"soli/formations/src/auth/errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Delete permission godoc
//
//	@Summary		Suppression permission
//	@Description	Suppression d'un permission dans la base de données
//	@Tags			permissions
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID permission"
//	@Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
//
//	@Success		204	{object}	string
//
//	@Failure		400	{object}	errors.APIError	"Impossible de parser le json"
//	@Failure		404	{object}	errors.APIError	"Permission non trouvée - Impossible de la supprimer "
//
//	@Router			/permissions/{id} [delete]
func (p permissionController) DeletePermission(ctx *gin.Context) {

	id, parseErr := uuid.Parse(ctx.Param("id"))
	if parseErr != nil {
		ctx.JSON(http.StatusBadRequest, &errors.APIError{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: parseErr.Error(),
		})
		return
	}

	errorDelete := p.service.DeletePermission(id)
	if errorDelete != nil {
		ctx.JSON(http.StatusNotFound, &errors.APIError{
			ErrorCode:    http.StatusNotFound,
			ErrorMessage: "Permission not found",
		})
		return
	}
	ctx.JSON(http.StatusNoContent, "Done")
}
