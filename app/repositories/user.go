package repositories

import (
	"context"

	"backend-golang/app/constants"
	"backend-golang/app/models"

	e "backend-golang/pkg/customerror"
)

type userRepository repository

type UserInterface interface {
	GetAll(ctx context.Context) ([]models.User, error)
	GetByID(ctx context.Context, id uint) (models.User, error)
	GetByPhone(ctx context.Context, phone string) (models.User, error)
	Create(ctx context.Context, user models.User) (models.User, error)
	Update(ctx context.Context, user models.User, id uint) (models.User, error)
	Delete(ctx context.Context, id uint) error
}

func (r *userRepository) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User

	err := r.Options.Postgres.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, e.NewInternalServiceError(err.Error())
	}

	return users, nil
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (models.User, error) {
	var user models.User

	err := r.Options.Postgres.WithContext(ctx).First(&user, id).Error
	if err != nil {
		if err.Error() == constants.GORM_ERR_NOT_FOUND {
			return user, e.NewNotFoundError("user not found")
		}
		return user, e.NewInternalServiceError(err.Error())
	}

	return user, nil
}

func (r *userRepository) GetByPhone(ctx context.Context, phone string) (models.User, error) {
	var user models.User

	err := r.Options.Postgres.WithContext(ctx).
		Where("phone_number = ?", phone).
		First(&user).Error

	if err != nil {
		if err.Error() == constants.GORM_ERR_NOT_FOUND {
			return user, e.NewNotFoundError("user not found")
		}
		return user, e.NewInternalServiceError(err.Error())
	}

	return user, nil
}

func (r *userRepository) Create(ctx context.Context, user models.User) (models.User, error) {
	err := r.Options.Postgres.WithContext(ctx).Create(&user).Error
	if err != nil {
		return user, e.NewInternalServiceError(err.Error())
	}
	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user models.User, id uint) (models.User, error) {
	err := r.Options.Postgres.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"name":         user.Name,
			"phone_number": user.PhoneNumber,
		}).Error

	if err != nil {
		return user, e.NewInternalServiceError(err.Error())
	}

	return user, nil
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	result := r.Options.Postgres.WithContext(ctx).
		Delete(&models.User{}, id)

	if result.RowsAffected == 0 {
		return e.NewNotFoundError("user not found")
	}

	if result.Error != nil {
		return e.NewInternalServiceError(result.Error.Error())
	}

	return nil
}