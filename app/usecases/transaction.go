// go:generate mockgen --destination=./mocks/mock_transaction_usecase.go --package=mocks --source=transaction_usecase.go
package usecases

import (
	"context"

	"backend-golang/app/models"
	e "backend-golang/pkg/customerror"
)

type transactionUsecase usecase

type TransactionInterface interface {
	CreateTransaction(ctx context.Context, req models.CreateTransactionRequest) (*models.Transaction, error)
	GetAllTransaction(ctx context.Context) ([]models.Transaction, error)
	GetDetailTransaction(ctx context.Context, id uint) (*models.Transaction, error)
}

func (u *transactionUsecase) CreateTransaction(ctx context.Context, req models.CreateTransactionRequest) (*models.Transaction, error) {

	if req.UserID == 0 {
		return nil, e.NewBadRequestError("user_id is required")
	}
	if req.PaketDataID == 0 {
		return nil, e.NewBadRequestError("paket_data_id is required")
	}

	user, err := u.Options.Repository.User.GetByID(ctx, req.UserID)
	if err != nil {
		return nil, e.NewNotFoundError("user not found")
	}

	paket, err := u.Options.Repository.PaketData.GetByID(ctx, req.PaketDataID)
	if err != nil {
		return nil, e.NewNotFoundError("paket data not found")
	}

	tx := models.Transaction{
		UserID:      user.ID,
		PaketDataID: paket.ID,
		Price:       paket.Price,
	}

	result, err := u.Options.Repository.Transaction.Create(ctx, tx)
	if err != nil {
		return nil, err
	}

	// 🔥 INI KUNCINYA
	fullData, err := u.Options.Repository.Transaction.GetByID(ctx, result.ID)
	if err != nil {
		return nil, err
	}

	return &fullData, nil
}

func (u *transactionUsecase) GetAllTransaction(ctx context.Context) ([]models.Transaction, error) {
	data, err := u.Options.Repository.Transaction.GetAll(ctx)
	return data, err
}

func (u *transactionUsecase) GetDetailTransaction(ctx context.Context, id uint) (*models.Transaction, error) {
	data, err := u.Options.Repository.Transaction.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
