package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/wellingtonchida/searchstreet/internal/service"
	"github.com/wellingtonchida/searchstreet/types"
)

type LogginCep struct {
	cep service.Service
}

func NewLogginService(svc service.Service) service.Service {
	return &LogginCep{
		cep: svc,
	}
}

func (l *LogginCep) GetCep(ctx context.Context, zipCode string) (cep *types.Address, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestId": ctx.Value("requestId"),
			"method":    "GetCep",
			"zipCode":   zipCode,
			"took":      time.Since(begin),
			"error":     err,
		}).Info("GetCep")
	}(time.Now())
	return l.cep.GetCep(ctx, zipCode)
}

func (l *LogginCep) GetCepFromDb(ctx context.Context, zipCode string) (cep types.Address, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestId": ctx.Value("requestId"),
			"method":    "GetCepFromDb",
			"zipCode":   zipCode,
			"took":      time.Since(begin),
			"error":     err,
		}).Info("GetCepFromDb")
	}(time.Now())
	return l.cep.GetCepFromDb(ctx, zipCode)
}

func (l *LogginCep) InsertCep(ctx context.Context, address *types.Address) (err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestId": ctx.Value("requestId"),
			"method":    "InsertCep",
			"zipCode":   address.ZipCode,
			"took":      time.Since(begin),
			"error":     err,
		}).Info("InsertCep")
	}(time.Now())
	return l.cep.InsertCep(ctx, address)
}
