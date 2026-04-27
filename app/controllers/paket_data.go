package controllers

import (
	"net/http"
	"strconv"

	"backend-golang/app/constants"
	"backend-golang/app/helpers"
	"backend-golang/app/models"

	e "backend-golang/pkg/customerror"

	"github.com/labstack/echo/v4"
)

type paketDataController controller

type PaketDataInterface interface {
	Create(ctx echo.Context) error
	Get(ctx echo.Context) error
	GetDetail(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

// @Summary Create Paket Data
// @Description Membuat paket data baru
// @Tags Paket Data
// @Accept json
// @Produce json
// @Param request body models.CreatePaketDataRequest true "Create Paket Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /v1/paket-data [post]
func (c *paketDataController) Create(ctx echo.Context) error {
	var reqBody models.CreatePaketDataRequest

	if err := ctx.Bind(&reqBody); err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"invalid request"})
	}

	data, err := c.Options.UseCases.PaketData.CreatePaket(ctx.Request().Context(), reqBody)
	if err != nil {
		return helpers.Response(ctx, e.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// @Summary Get All Paket Data
// @Description Ambil semua paket data
// @Tags Paket Data
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /v1/paket-data [get]
func (c *paketDataController) Get(ctx echo.Context) error {
	data, err := c.Options.UseCases.PaketData.GetAllPaket(ctx.Request().Context())
	if err != nil {
		return helpers.Response(ctx, e.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// @Summary Get Paket Data by ID
// @Description Ambil detail paket data berdasarkan ID
// @Tags Paket Data
// @Produce json
// @Param id path int true "Paket Data ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /v1/paket-data/{id} [get]
func (c *paketDataController) GetDetail(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := c.Options.UseCases.PaketData.GetDetailPaket(ctx.Request().Context(), uint(id))
	if err != nil {
		return helpers.Response(ctx, e.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// @Summary Update Paket Data
// @Description Update paket data berdasarkan ID
// @Tags Paket Data
// @Accept json
// @Produce json
// @Param id path int true "Paket Data ID"
// @Param request body models.UpdatePaketDataRequest true "Update Paket Data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /v1/paket-data/{id} [patch]
func (c *paketDataController) Update(ctx echo.Context) error {
	var reqBody models.UpdatePaketDataRequest

	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.Bind(&reqBody); err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"invalid request"})
	}

	data, err := c.Options.UseCases.PaketData.UpdatePaket(ctx.Request().Context(), reqBody, uint(id))
	if err != nil {
		return helpers.Response(ctx, e.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// @Summary Delete Paket Data
// @Description Hapus paket data berdasarkan ID
// @Tags Paket Data
// @Produce json
// @Param id path int true "Paket Data ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /v1/paket-data/{id} [delete]
func (c *paketDataController) Delete(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.Options.UseCases.PaketData.DeletePaket(ctx.Request().Context(), uint(id))
	if err != nil {
		return helpers.Response(ctx, e.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, nil, nil)
}
