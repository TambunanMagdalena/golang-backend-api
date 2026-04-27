package repositories

import (
	"context"

	"backend-golang/app/constants"
	"backend-golang/app/models"

	e "backend-golang/pkg/customerror"
)

type paketDataRepository repository

type PaketDataInterface interface {
	GetAll(ctx context.Context) ([]models.PaketData, error)
	GetByID(ctx context.Context, id uint) (models.PaketData, error)
	Create(ctx context.Context, paket models.PaketData) (models.PaketData, error)
	Update(ctx context.Context, paket models.PaketData, id uint) (models.PaketData, error)
	Delete(ctx context.Context, id uint) error
}

func (r *paketDataRepository) GetAll(ctx context.Context) ([]models.PaketData, error) {
	var pakets []models.PaketData

	err := r.Options.Postgres.WithContext(ctx).Find(&pakets).Error
	if err != nil {
		return nil, e.NewInternalServiceError(err.Error())
	}

	return pakets, nil
}

func (r *paketDataRepository) GetByID(ctx context.Context, id uint) (models.PaketData, error) {
	var paket models.PaketData

	err := r.Options.Postgres.WithContext(ctx).First(&paket, id).Error
	if err != nil {
		if err.Error() == constants.GORM_ERR_NOT_FOUND {
			return paket, e.NewNotFoundError("paket data not found")
		}
		return paket, e.NewInternalServiceError(err.Error())
	}

	return paket, nil
}

func (r *paketDataRepository) Create(ctx context.Context, paket models.PaketData) (models.PaketData, error) {
	err := r.Options.Postgres.WithContext(ctx).Create(&paket).Error
	if err != nil {
		return paket, e.NewInternalServiceError(err.Error())
	}
	return paket, nil
}

func (r *paketDataRepository) Update(ctx context.Context, paket models.PaketData, id uint) (models.PaketData, error) {
	err := r.Options.Postgres.WithContext(ctx).
		Model(&models.PaketData{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"name":          paket.Name,
			"price":         paket.Price,
			"quota":         paket.Quota,
			"active_period": paket.ActivePeriod,
		}).Error

	if err != nil {
		return paket, e.NewInternalServiceError(err.Error())
	}

	return paket, nil
}

func (r *paketDataRepository) Delete(ctx context.Context, id uint) error {
	result := r.Options.Postgres.WithContext(ctx).
		Delete(&models.PaketData{}, id)

	if result.RowsAffected == 0 {
		return e.NewNotFoundError("paket data not found")
	}

	if result.Error != nil {
		return e.NewInternalServiceError(result.Error.Error())
	}

	return nil
}