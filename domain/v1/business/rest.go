package business

import (
	"strconv"

	business_model "62tech.co/service/domain/v1/business/model"
	business_usecase "62tech.co/service/domain/v1/business/usecase"
	"62tech.co/service/model"

	"net/http"

	"62tech.co/service/utl/response"

	"github.com/labstack/echo/v4"
)

func (m *Module) HandleRest(group *echo.Group) {
	group.GET("/list", m.businessList).Name = "business-list"
	group.GET("/detail/:param", m.businessDetail).Name = "business-detail"
	group.POST("/create", m.createBusiness).Name = "business-create"
	group.PUT("/update/:param", m.businessUpdate).Name = "business-update"
	group.DELETE("/delete/:param", m.productDelete).Name = "business-delete"
	group.GET("/init", m.InitBusiness).Name = "business-initialize"
}

// @Summary Initialization Unit Business
// @Description Initialization Unit Business
// @Tags Business
// @Accept */*
// @Produce json
// @Success 200 {interface} model.Response{}
// @Router /api/v1/business/init [get]
func (m *Module) InitBusiness(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
	)

	// usecase get user list
	resp, err := business_usecase.BusinessInitialize(m.Config)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Error:   err,
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: nil,
		Data:    resp,
	})
}

// @Summary Get List Product
// @Description get the status of server.
// @Tags Product Management - Product
// @Accept */*
// @Produce json
// @Param limit query int false "Filter Limit"
// @Param page query int false "Filter Page"
// @Success 200 {interface} model.Response{}
// @Router /api/v1/business/list [get]
func (m *Module) businessList(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
	)

	// usecase get user list
	resp, err := business_usecase.BusinessesList(m.Config)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Error:   err,
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: nil,
		Data:    resp,
	})
}

// @Summary Create Business
// @Description get the status of server.
// @Tags Business
// @Accept */*
// @Produce json
// @Param Body body business_model.Business true "Payload"
// @Success 200 {interface} model.Response{}
// @Router /api/v1/business/create [post]
func (m *Module) createBusiness(c echo.Context) error {

	var (
		requestId = c.Get("request_id").(string)
		payload   = business_model.Business{}
	)

	err := c.Bind(&payload)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err,
		})
	}

	resp, err := business_usecase.CreateNewBusiness(m.Config, &payload)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: "error",
			Error:   err.Error(),
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: "created business success",
		Data:    resp,
	})
}

// @Summary Update Business
// @Description get the status of server.
// @Tags Business
// @Accept */*
// @Produce json
// @Param param path string true "Business Id"
// @Param Body body business_model.Business true "Payload"
// @Success 200 {interface} model.Response{}
// @Router /api/v1/business/update/{param} [put]
func (m *Module) businessUpdate(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
		payload   = business_model.Business{}
	)

	id := c.Param("param")
	payload.ID = id

	err := c.Bind(&payload)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err,
		})
	}

	resp, err := business_usecase.BusinessUpdate(m.Config, &payload)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err.Error(),
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: "your data has been updated",
		Data:    resp,
	})
}

// @Summary Delete Business
// @Description get the status of server.
// @Tags Business
// @Accept */*
// @Produce json
// @Param param path string true "Business Id"
// @Success 200 {interface} model.Response{}
// @Router /api/v1/business/delete/{param} [delete]
func (m *Module) productDelete(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
		payload   = business_model.Business{}
	)

	id := c.Param("param")
	payload.ID = id

	err := c.Bind(&payload)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err,
		})
	}

	_, err = business_usecase.BusinessDelete(m.Config, &payload)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err.Error(),
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: "your data has been updated",
		Data:    nil,
	})
}

// @Summary Detail of Business
// @Description get the status of server.
// @Tags Business
// @Accept */*
// @Produce json
// @Param param path string true "Business Id"
// @Success 200 {interface} model.Response{}
// @Router /api/v1/business/detail/{param} [get]
func (m *Module) businessDetail(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
	)
	id, _ := strconv.Atoi(c.Param("param"))
	resp, err := business_usecase.BusinessDetail(m.Config, id)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: nil,
		Data:    resp,
		Error:   err,
	})
}
