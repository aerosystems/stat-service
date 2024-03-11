package main

import (
	"github.com/aerosystems/stat-service/internal/config"
	"github.com/sirupsen/logrus"
)

type App struct {
	log *logrus.Logger
	cfg *config.Config
}
