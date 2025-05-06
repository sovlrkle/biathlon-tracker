package usecases

import (
	"biathlon-tracker/internal/domain"
	"fmt"
	"time"
)

var _ PenaltyUseCase = (*eventUseCase)(nil)

func (e *eventUseCase) EnterPenaltyLap(competitorID int, eventTime time.Time) error {
	competitor, err := e.repo.Get(competitorID)
	if err != nil {
		return err
	}
	if competitor.Status != domain.LeftFiringRange && competitor.Status != domain.LeftPenaltyLap {
		return fmt.Errorf("competitor(%d) must left firing range or previous penalty lap before entering a penalty lap", competitorID)
	}
	competitor.Status = domain.PenaltyLap
	competitor.PenaltyStartTime = eventTime
	return nil
}

func (e *eventUseCase) LeftPenaltyLap(competitorID int, eventTime time.Time) error {
	competitor, err := e.repo.Get(competitorID)
	if err != nil {
		return err
	}
	if competitor.Status != domain.PenaltyLap {
		return fmt.Errorf("competitor(%d) must be in penalty lap before leaving it", competitorID)
	}
	if !competitor.PenaltyStartTime.IsZero() {
		duration := eventTime.Sub(competitor.PenaltyStartTime)
		competitor.PenaltyDurations = append(competitor.PenaltyDurations, duration)
		competitor.TotalPenaltyDuration += duration
		competitor.PenaltyStartTime = time.Time{}
		competitor.Status = domain.Started
	}
	return nil
}
