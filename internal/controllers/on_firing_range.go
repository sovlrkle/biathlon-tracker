package controllers

import (
	"fmt"
	"strconv"
)

func (e *eventController) OnFiringRange(message string) (string, error) {
	event, err := parse(message)
	if err != nil {
		return "", err
	}

	if len(event.Params) != 1 {
		return "", fmt.Errorf("%v: only firingRange is required", ErrInvalidParams)
	}
	firingRange, err := strconv.Atoi(event.Params[0])
	return event.String(), e.firingRangeUseCase.EnterFiringRange(event.CompetitorID, firingRange)
}
