package logger

import "go.uber.org/zap"

var (
	CreateUserJourney = zap.String("journey", "create-user")
)
