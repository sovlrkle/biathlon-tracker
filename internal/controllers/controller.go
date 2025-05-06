package controllers

import (
	"biathlon-tracker/config"
	"biathlon-tracker/internal/usecases"

	"github.com/pkg/errors"
)

type Controller interface {
	Register(message string) (string, error)
	SetStartTime(message string) (string, error)
	OnStartLine(message string) (string, error)
	Start(message string) (string, error)
	OnFiringRange(message string) (string, error)
	HitTarget(message string) (string, error)
	LeftFiringRange(message string) (string, error)
	EnterPenaltyLap(message string) (string, error)
	LeftPenaltyLap(message string) (string, error)
	EndMainLap(message string) (string, error)
	CantContinue(message string) (string, error)
	GenerateReport() string
}

var _ Controller = (*eventController)(nil)

var (
	ErrInvalidEvent  = errors.New("invalid event")
	ErrInvalidParams = errors.New("invalid params")
)

type eventController struct {
	reportUseCase      usecases.ReportUseCase
	firingRangeUseCase usecases.FiringRangeUseCase
	lapUseCase         usecases.LapUseCase
	penaltyUseCase     usecases.PenaltyUseCase
	competitorUseCase  usecases.CompetitorUseCase
	cfg                *config.Config
}

func New(
	reportUseCase usecases.ReportUseCase,
	firingRangeUseCase usecases.FiringRangeUseCase,
	lapUseCase usecases.LapUseCase,
	penaltyUseCase usecases.PenaltyUseCase,
	competitorUseCase usecases.CompetitorUseCase,
	cfg *config.Config,
) *eventController {
	return &eventController{
		reportUseCase:      reportUseCase,
		firingRangeUseCase: firingRangeUseCase,
		lapUseCase:         lapUseCase,
		penaltyUseCase:     penaltyUseCase,
		competitorUseCase:  competitorUseCase,
		cfg:                cfg,
	}
}
