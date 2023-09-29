package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *BaseHandler) GetEvents(c echo.Context) error {
	projectToken := c.QueryParam("projectToken")
	size, err := strconv.Atoi(c.QueryParam("size"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "size parameter must be integer", err)
	}
	from, err := strconv.Atoi(c.QueryParam("from"))
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, "from parameter must be integer", err)
	}
	res, err := h.eventRepo.GetByProjectToken(projectToken, size, from)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err)
	}
	return SuccessResponse(c, http.StatusOK, "OK", res)
}
