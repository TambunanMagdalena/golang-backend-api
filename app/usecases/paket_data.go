package usecases

import (
	"context"

	"backend-golang/app/models"
	e "backend-golang/pkg/customerror"
)

type paketDataUsecase usecase

type PaketDataInterface interface {
	CreatePaket(ctx context.Context, req models.CreatePaketDataRequest) (*models.PaketData, error)
	GetAllPaket(ctx context.Context) ([]models.PaketData, error)
	GetDetailPaket(ctx context.Context, id uint) (*models.PaketData, error)
	UpdatePaket(ctx context.Context, req models.UpdatePaketDataRequest, id uint) (*models.PaketData, error)
	DeletePaket(ctx context.Context, id uint) error
}

func (u *paketDataUsecase) CreatePaket(ctx context.Context, req models.CreatePaketDataRequest) (*models.PaketData, error) {

	if req.Name == "" {
		return nil, e.NewBadRequestError("name is required")
	}
	if req.Price <= 0 {
		return nil, e.NewBadRequestError("price must be greater than 0")
	}
	if req.Quota <= 0 {
		return nil, e.NewBadRequestError("quota must be greater than 0")
	}
	if req.ActivePeriod <= 0 {
		return nil, e.NewBadRequestError("active_period must be greater than 0")
	}

	paket := models.PaketData{
		Name:         req.Name,
		Price:        req.Price,
		Quota:        req.Quota,
		ActivePeriod: req.ActivePeriod,
	}

	result, err := u.Options.Repository.PaketData.Create(ctx, paket)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (u *paketDataUsecase) GetAllPaket(ctx context.Context) ([]models.PaketData, error) {
	return u.Options.Repository.PaketData.GetAll(ctx)
}

func (u *paketDataUsecase) GetDetailPaket(ctx context.Context, id uint) (*models.PaketData, error) {
	data, err := u.Options.Repository.PaketData.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (u *paketDataUsecase) UpdatePaket(ctx context.Context, req models.UpdatePaketDataRequest, id uint) (*models.PaketData, error) {

	if req.Name == "" {
		return nil, e.NewBadRequestError("name is required")
	}
	if req.Price <= 0 {
		return nil, e.NewBadRequestError("price must be greater than 0")
	}
	if req.Quota <= 0 {
		return nil, e.NewBadRequestError("quota must be greater than 0")
	}
	if req.ActivePeriod <= 0 {
		return nil, e.NewBadRequestError("active_period must be greater than 0")
	}

	_, err := u.Options.Repository.PaketData.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	paket := models.PaketData{
		Name:         req.Name,
		Price:        req.Price,
		Quota:        req.Quota,
		ActivePeriod: req.ActivePeriod,
	}

	result, err := u.Options.Repository.PaketData.Update(ctx, paket, id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (u *paketDataUsecase) DeletePaket(ctx context.Context, id uint) error {
	return u.Options.Repository.PaketData.Delete(ctx, id)
}