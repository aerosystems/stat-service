package main

import (
	"github.com/aerosystems/stat-service/internal/handlers"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Log         *logrus.Logger
	BaseHandler *handlers.BaseHandler
}

func NewConfig(
	log *logrus.Logger,
	baseHandler *handlers.BaseHandler,
) *Config {
	return &Config{
		Log:         log,
		BaseHandler: baseHandler,
	}
}
