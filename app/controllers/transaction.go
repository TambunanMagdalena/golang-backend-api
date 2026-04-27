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

type transactionController controller

type TransactionInterface interface {
	Create(ctx echo.Context) error
	Get(ctx echo.Context) error
	GetDetail(ctx echo.Context) error
}

// @Summary Create Transaction
// @Description User membeli paket data
// @Tags Transactions
// @Accept json
// @Produce json
// @Param request body models.CreateTransactionRequest true "Create Transaction"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /v1/transactions [post]
func (c *transactionController) Create(ctx echo.Context) error {
	var (
		reqBody models.CreateTransactionRequest
		resBody *models.Transaction
		err     error
	)

	if err = ctx.Bind(&reqBody); err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"invalid request"})
	}

	resBody, err = c.Options.UseCases.Transaction.CreateTransaction(ctx.Request().Context(), reqBody)
	if err != nil {
		return helpers.Response(ctx, e.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, resBody, nil)
}

// @Summary Get All Transactions
// @Description Ambil semua data transaksi
// @Tags Transactions
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /v1/transactions [get]
func (c *transactionController) Get(ctx echo.Context) error {
	var (
		data []models.Transaction
		err  error
	)

	data, err = c.Options.UseCases.Transaction.GetAllTransaction(ctx.Request().Context())
	if err != nil {
		return helpers.Response(ctx, e.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// @Summary Get Transaction by ID
// @Description Ambil detail transaksi berdasarkan ID
// @Tags Transactions
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /v1/transactions/{id} [get]
func (c *transactionController) GetDetail(ctx echo.Context) error {
	var (
		data *models.Transaction
		err  error
		id   int
	)

	id, _ = strconv.Atoi(ctx.Param("id"))

	data, err = c.Options.UseCases.Transaction.GetDetailTransaction(ctx.Request().Context(), uint(id))
	if err != nil {
		return helpers.Response(ctx, e.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}
