package handlers

import (
	"github.com/aerosystems/stat-service/internal/models"
	"github.com/aerosystems/stat-service/internal/services"
	RangeService "github.com/aerosystems/stat-service/pkg/range_service"
	"github.com/google/uuid"
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
// @Failure 404 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /v1/events [get]
func (h *BaseHandler) GetEvents(c echo.Context) error {
	accessTokenClaims := c.Get("accessTokenClaims").(*services.AccessTokenClaims)
	projectToken := c.QueryParam("projectToken")
	pagination, err := RangeService.GetLimitPaginationFromQuery(c.QueryParams())
	if err != nil {
		return h.ErrorResponse(c, http.StatusBadRequest, err.Error(), err)
	}
	timeRange, err := RangeService.GetTimeRangeFromQuery(c.QueryParams())
	if err != nil {
		return h.ErrorResponse(c, http.StatusBadRequest, err.Error(), err)
	}

	if !h.eventService.IsAccess(projectToken, uuid.MustParse(accessTokenClaims.UserUuid)) {
		return h.ErrorResponse(c, http.StatusForbidden, "access denied", nil)
	}

	res, total, err := h.eventService.GetByProjectToken(projectToken, models.InspectEvent, *timeRange, *pagination)
	if err != nil {
		return h.ErrorResponse(c, http.StatusInternalServerError, "could not get events", err)
	}
	if total == 0 {
		return h.ErrorResponse(c, http.StatusNotFound, "events not found", nil)
	}
	return c.JSON(http.StatusOK, Response{
		Message: "events successfully found",
		Total:   total,
		Data:    res,
	})
}
