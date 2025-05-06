package usecases

import (
	"biathlon-tracker/internal/domain"
	"time"
)

var _ LapUseCase = (*eventUseCase)(nil)

func (e *eventUseCase) EndLap(competitorID int, eventTime time.Time) error {
	competitor, err := e.repo.Get(competitorID)
	if err != nil {
		return err
	}

	competitor.CurrentLap++
	competitor.Status = domain.Started

	if len(competitor.LapStartTimes) > 0 {
		lastLapStart := competitor.LapStartTimes[len(competitor.LapStartTimes)-1]
		duration := eventTime.Sub(lastLapStart)
		competitor.LapDurations = append(competitor.LapDurations, duration)
	}

	if competitor.CurrentLap < e.cfg.Laps {
		competitor.LapStartTimes = append(competitor.LapStartTimes, eventTime)
	}

	return nil
}
