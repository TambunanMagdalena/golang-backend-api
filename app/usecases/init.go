package usecases

import (
	"backend-golang/app/repositories"
	"backend-golang/pkg/config"
)

type Main struct {
	PaketData   PaketDataInterface
	Transaction TransactionInterface
	User        UserInterface
}

type usecase struct {
	Options Options
}

type Options struct {
	Repository *repositories.Main
	Config     *config.Config
}

func Init(opts Options) *Main {
	ucs := &usecase{
		Options: opts,
	}

	m := &Main{
		PaketData:   (*paketDataUsecase)(ucs),
		Transaction: (*transactionUsecase)(ucs),
		User:        (*userUsecase)(ucs),
	}

	return m
}
