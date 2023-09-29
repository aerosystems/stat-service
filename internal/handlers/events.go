package handlers

import (
	"github.com/aerosystems/stat-service/internal/helpers"
	"github.com/aerosystems/stat-service/internal/pagination"
	RPCClient "github.com/aerosystems/stat-service/internal/rpc_client"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

// GetEvents godoc
// @Summary Get Events
// @Description Get Events by project token
// @Tags events
// @Accept  json
// @Produce application/json
// @Security BearerAuth
// @Param projectToken query string true "Project Token"
// @Param limit query int false "Limit. Must be integer. Default 10"
// @Param offset query int false "Offset. Must be integer. Default 0"
// @Success 200 {object} Response{data=[]models.Event}
// @Failure 400 {object} ErrResponse
// @Failure 401 {object} ErrResponse
// @Failure 403 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /v1/events [get]
func (h *BaseHandler) GetEvents(c echo.Context) error {
	userId := c.Get("userId").(int)
	userRole := c.Get("userRole").(string)
	projectToken := c.QueryParam("projectToken")
	pagination, err := pagination.GetFromQuery(c.QueryParams())
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, err.Error(), err)
	}

	projectList, err := RPCClient.GetProjectList(userId)
	if err != nil {
		log.Println(err)
		return ErrorResponse(c, http.StatusInternalServerError, "could not get events", err)
	}
	if helpers.ContainsProjectToken(*projectList, projectToken) && !helpers.ContainsString([]string{"admin", "support"}, userRole) {
		return ErrorResponse(c, http.StatusForbidden, "access denied", nil)
	}

	res, err := h.eventRepo.GetByProjectToken(projectToken, "inspect", *pagination)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "could not get events", err)
	}
	return SuccessResponse(c, http.StatusOK, "OK", res)
}
