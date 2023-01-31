package system

import (
	"net/http"

	"62tech.co/service/model"
	"62tech.co/service/utl/response"

	"github.com/labstack/echo/v4"
)

func (m *Module) HandleRest(group *echo.Group) {
	group.GET("/health", m.HealthCheck).Name = "user-list"
}

// @Summary Get Health Information
// @Description get the status of server.
// @Tags System
// @Accept */*
// @Produce json
// @Success 200 {interface} model.Response{}
// @Router /api/v1/system/health [get]
func (m *Module) HealthCheck(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
	)
	resp := map[string]interface{}{
		"health": "ok",
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: nil,
		Data:    resp,
	})
}
