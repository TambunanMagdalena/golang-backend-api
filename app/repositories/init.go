package repositories

import (
	
	"backend-golang/pkg/config"

	"gorm.io/gorm"
)

type Main struct {
	PaketData PaketDataInterface
	Transaction TransactionInterface
	User UserInterface
}


type repository struct {
	Options Options
}

type Options struct {
	Postgres *gorm.DB
	Config   *config.Config
}

func Init(opts Options) *Main {
	repo := &repository{opts}

	m := &Main{
		PaketData: (*paketDataRepository)(repo),
		Transaction: (*transactionRepository)(repo),
		User: (*userRepository)(repo),
	}

	return m
}
