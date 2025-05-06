package usecases

import (
	"biathlon-tracker/internal/domain"
	"fmt"
	"time"
)

var _ CompetitorUseCase = (*eventUseCase)(nil)

func (e *eventUseCase) Register(competitorID int, eventTime time.Time) error {
	competitor, err := e.repo.Get(competitorID)
	if err != nil {
		competitor = domain.NewCompetitor(competitorID)
		e.repo.Add(competitor)
	}
	competitor.Status = domain.Registered
	competitor.RegistrationTime = eventTime
	competitor.PlannedStartTime = e.cfg.StartTime.Add(time.Duration(competitorID-1) * e.cfg.Delta)
	return nil
}

func (e *eventUseCase) SetStartTime(competitorID int, startTime time.Time) error {
	competitor, err := e.repo.Get(competitorID)
	if err != nil {
		return err
	}
	if competitor.Status != domain.Registered {
		return fmt.Errorf("competitor(%d) must be registered before setting start time", competitorID)
	}
	competitor.Status = domain.StartTimeSet
	competitor.PlannedStartTime = startTime
	return nil
}

func (e *eventUseCase) OnStartLine(competitorID int) error {
	competitor, err := e.repo.Get(competitorID)
	if err != nil {
		return err
	}
	if competitor.Status != domain.Registered && competitor.Status != domain.StartTimeSet {
		return fmt.Errorf("competitor(%d) must be registered or have start time set by a draw before being on start line", competitorID)
	}
	competitor.Status = domain.OnStartLine
	return nil
}

func (e *eventUseCase) Start(competitorID int, startTime time.Time) error {
	competitor, err := e.repo.Get(competitorID)
	if err != nil {
		return err
	}

	if competitor.Status != domain.OnStartLine {
		return fmt.Errorf("competitor(%d) must be on start line before starting", competitorID)
	}

	if startTime.After(competitor.PlannedStartTime.Add(e.cfg.Delta)) {
		competitor.Status = domain.Disqualified
	} else {
		competitor.Status = domain.Started
		competitor.StartTime = startTime
		competitor.LapStartTimes = append(competitor.LapStartTimes, startTime)
	}
	return nil
}

func (e *eventUseCase) CantContinue(competitorID int) error {
	competitor, err := e.repo.Get(competitorID)
	if err != nil {
		return err
	}
	competitor.Status = domain.NotFinished
	return nil
}
