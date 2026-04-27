package controllers

import (
	"backend-golang/app/usecases"
	"backend-golang/pkg/config"
)

type Main struct {
	User        UserInterface
	PaketData   PaketDataInterface
	Transaction TransactionInterface
}

type controller struct {
	Options Options
}

type Options struct {
	Config   *config.Config
	UseCases *usecases.Main
}

func Init(opts Options) *Main {
	ctrl := &controller{opts}

	m := &Main{
		User:        (*userController)(ctrl),
		PaketData:   (*paketDataController)(ctrl),
		Transaction: (*transactionController)(ctrl),
	}

	return m
}
