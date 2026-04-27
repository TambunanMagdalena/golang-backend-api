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

type userController controller

type UserInterface interface {
	Create(ctx echo.Context) error
	Get(ctx echo.Context) error
	GetDetail(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

// Create User godoc
// @Summary Create User
// @Description Create new user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body models.CreateUserRequest true "Create User"
// @Success 201 {object} models.User
// @Router /v1/users [post]
func (c *userController) Create(ctx echo.Context) error {
	var (
		reqBody models.CreateUserRequest
		resBody *models.User
		err     error
	)

	if err = ctx.Bind(&reqBody); err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"invalid request"})
	}

	resBody, err = c.Options.UseCases.User.CreateUser(ctx.Request().Context(), reqBody)
	if err != nil {
		return helpers.Response(ctx, e.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, resBody, nil)
}

// @Summary Get All Users
// @Description Retrieve all users
// @Tags Users
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /v1/users [get]
func (c *userController) Get(ctx echo.Context) error {
	data, err := c.Options.UseCases.User.GetAllUser(ctx.Request().Context())
	if err != nil {
		return helpers.Response(ctx, e.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// @Summary Get User by ID
// @Description Retrieve user detail by ID
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /v1/users/{id} [get]
func (c *userController) GetDetail(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := c.Options.UseCases.User.GetDetailUser(ctx.Request().Context(), uint(id))
	if err != nil {
		return helpers.Response(ctx, e.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// @Summary Update User
// @Description Update user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body models.UpdateUserRequest true "Update User"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /v1/users/{id} [patch]
func (c *userController) Update(ctx echo.Context) error {
	var (
		reqBody models.UpdateUserRequest
		resBody *models.User
		err     error
	)

	id, _ := strconv.Atoi(ctx.Param("id"))

	if err = ctx.Bind(&reqBody); err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"invalid request"})
	}

	resBody, err = c.Options.UseCases.User.UpdateUser(ctx.Request().Context(), reqBody, uint(id))
	if err != nil {
		return helpers.Response(ctx, e.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, resBody, nil)
}

// @Summary Delete User
// @Description Delete user by ID
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /v1/users/{id} [delete]
func (c *userController) Delete(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.Options.UseCases.User.DeleteUser(ctx.Request().Context(), uint(id))
	if err != nil {
		return helpers.Response(ctx, e.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, nil, nil)
}
