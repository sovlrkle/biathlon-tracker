package usecases

import (
	"biathlon-tracker/internal/domain"
	"fmt"
)

var _ FiringRangeUseCase = (*eventUseCase)(nil)

func (e *eventUseCase) EnterFiringRange(competitorID int, firingRange int) error {
	competitor, err := e.repo.Get(competitorID)
	if err != nil {
		return err
	}

	competitor.Status = domain.OnFiringRange
	competitor.CurrentFiringRange = firingRange
	competitor.Targets = make(map[int]bool)
	return nil
}

func (e *eventUseCase) HitTarget(competitorID int, target int) error {
	competitor, err := e.repo.Get(competitorID)
	if err != nil {
		return err
	}

	if competitor.Status != domain.OnFiringRange && competitor.Status != domain.HitTargets {
		return fmt.Errorf("competitor(%d) must be on firing range before hitting target", competitorID)
	}
	competitor.Status = domain.HitTargets

	if !competitor.Targets[target] {
		competitor.Targets[target] = true
		competitor.Hits++
	}
	return nil
}

func (e *eventUseCase) LeftFiringRange(competitorID int) error {
	competitor, err := e.repo.Get(competitorID)
	if err != nil {
		return err
	}

	if competitor.Status != domain.OnFiringRange && competitor.Status != domain.HitTargets {
		return fmt.Errorf("competitor(%d) must be on firing range before leaving it", competitorID)
	}

	competitor.Misses += 5 - len(competitor.Targets)
	competitor.CurrentFiringRange = -1
	competitor.Status = domain.LeftFiringRange
	return nil
}
