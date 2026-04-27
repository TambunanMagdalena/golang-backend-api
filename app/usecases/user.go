// go:generate mockgen --destination=./mocks/mock_user_usecase.go --package=mocks --source=user_usecase.go
package usecases

import (
	"context"

	"backend-golang/app/models"
	e "backend-golang/pkg/customerror"
)

type userUsecase usecase

type UserInterface interface {
	CreateUser(ctx context.Context, req models.CreateUserRequest) (*models.User, error)
	GetAllUser(ctx context.Context) ([]models.User, error)
	GetDetailUser(ctx context.Context, id uint) (*models.User, error)
	UpdateUser(ctx context.Context, req models.UpdateUserRequest, id uint) (*models.User, error)
	DeleteUser(ctx context.Context, id uint) error
}

func (u *userUsecase) CreateUser(ctx context.Context, req models.CreateUserRequest) (*models.User, error) {

	if req.Name == "" {
		return nil, e.NewBadRequestError("name is required")
	}
	if req.PhoneNumber == "" {
		return nil, e.NewBadRequestError("phone_number is required")
	}

	// 🔥 CHECK DUPLICATE
	_, err := u.Options.Repository.User.GetByPhone(ctx, req.PhoneNumber)
	if err == nil {
		return nil, e.NewBadRequestError("phone number already exists")
	}

	user := models.User{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}

	result, err := u.Options.Repository.User.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (u *userUsecase) GetAllUser(ctx context.Context) ([]models.User, error) {
	return u.Options.Repository.User.GetAll(ctx)
}

func (u *userUsecase) GetDetailUser(ctx context.Context, id uint) (*models.User, error) {
	data, err := u.Options.Repository.User.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (u *userUsecase) UpdateUser(ctx context.Context, req models.UpdateUserRequest, id uint) (*models.User, error) {

	if req.Name == "" {
		return nil, e.NewBadRequestError("name is required")
	}
	if req.PhoneNumber == "" {
		return nil, e.NewBadRequestError("phone_number is required")
	}

	// 🔥 CHECK EXIST
	_, err := u.Options.Repository.User.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}

	result, err := u.Options.Repository.User.Update(ctx, user, id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (u *userUsecase) DeleteUser(ctx context.Context, id uint) error {
	return u.Options.Repository.User.Delete(ctx, id)
}