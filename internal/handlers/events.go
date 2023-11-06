package handlers

import (
	"github.com/aerosystems/stat-service/internal/helpers"
	RPCClient "github.com/aerosystems/stat-service/internal/rpc_client"
	RangeService "github.com/aerosystems/stat-service/pkg/range_service"
	"github.com/labstack/echo/v4"
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
// @Param startTime query string false "Start time in RFC3339 format. Default NOW - 24 hours"
// @Param endTime query string false "End time in RFC3339 format. Default NOW"
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
	pagination, err := RangeService.GetLimitPaginationFromQuery(c.QueryParams())
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, 400801, err.Error(), err)
	}
	timeRange, err := RangeService.GetTimeRangeFromQuery(c.QueryParams())
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, 400802, err.Error(), err)
	}

	projectList, err := RPCClient.GetProjectList(userId)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, 500801, "could not get events", err)
	}
	if helpers.ContainsProjectToken(*projectList, projectToken) && !helpers.ContainsString([]string{"staff"}, userRole) {
		return ErrorResponse(c, http.StatusForbidden, 403801, "access denied", nil)
	}

	res, total, err := h.eventRepo.GetByProjectToken(projectToken, "inspect", *timeRange, *pagination)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, 500801, "could not get events", err)
	}
	if len(res) == 0 {
		return ErrorResponse(c, http.StatusNotFound, 404801, "events not found", nil)
	}
	return c.JSON(http.StatusOK, Response{
		Message: "events successfully found",
		Total:   total,
		Data:    res,
	})
}
