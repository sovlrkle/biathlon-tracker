package usecases

import (
	"biathlon-tracker/config"
	"biathlon-tracker/internal/repository"
	"time"
)

type ReportUseCase interface {
	GenerateReport() string
}

type FiringRangeUseCase interface {
	EnterFiringRange(competitorID int, firingRange int) error
	LeftFiringRange(competitorID int) error
	HitTarget(competitorID int, target int) error
}

type LapUseCase interface {
	EndLap(competitorID int, eventTime time.Time) error
}

type PenaltyUseCase interface {
	EnterPenaltyLap(competitorID int, eventTime time.Time) error
	LeftPenaltyLap(competitorID int, eventTime time.Time) error
}

type CompetitorUseCase interface {
	Register(competitorID int, eventTime time.Time) error
	CantContinue(competitorID int) error
	OnStartLine(competitorID int) error
	SetStartTime(competitorID int, startTime time.Time) error
	Start(competitorID int, startTime time.Time) error
}

type eventUseCase struct {
	repo repository.Repository
	cfg  *config.Config
}

func New(repo repository.Repository, cfg *config.Config) *eventUseCase {
	return &eventUseCase{
		repo: repo,
		cfg:  cfg,
	}
}
