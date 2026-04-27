package repositories

import (
	"context"

	"backend-golang/app/constants"
	"backend-golang/app/models"

	e "backend-golang/pkg/customerror"
)

type transactionRepository repository

type TransactionInterface interface {
	GetAll(ctx context.Context) ([]models.Transaction, error)
	GetByID(ctx context.Context, id uint) (models.Transaction, error)
	Create(ctx context.Context, tx models.Transaction) (models.Transaction, error)
}

func (r *transactionRepository) GetAll(ctx context.Context) ([]models.Transaction, error) {
	var txs []models.Transaction

	err := r.Options.Postgres.
		WithContext(ctx).
		Preload("User").
		Preload("PaketData").
		Find(&txs).Error

	if err != nil {
		return nil, e.NewInternalServiceError(err.Error())
	}

	return txs, nil
}

func (r *transactionRepository) GetByID(ctx context.Context, id uint) (models.Transaction, error) {
	var tx models.Transaction

	err := r.Options.Postgres.
		WithContext(ctx).
		Preload("User").
		Preload("PaketData").
		First(&tx, id).Error

	if err != nil {
		if err.Error() == constants.GORM_ERR_NOT_FOUND {
			return tx, e.NewNotFoundError("transaction not found")
		}
		return tx, e.NewInternalServiceError(err.Error())
	}

	return tx, nil
}

func (r *transactionRepository) Create(ctx context.Context, tx models.Transaction) (models.Transaction, error) {
	err := r.Options.Postgres.WithContext(ctx).Create(&tx).Error
	if err != nil {
		return tx, e.NewInternalServiceError(err.Error())
	}
	return tx, nil
}
