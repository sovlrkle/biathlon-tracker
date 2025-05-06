package controllers

import "fmt"

func (e *eventController) SetStartTime(message string) (string, error) {
	event, err := parse(message)
	if err != nil {
		return "", err
	}
	if len(event.Params) != 1 {
		return "", fmt.Errorf("%v: only startTime is required", ErrInvalidParams)
	}
	startTime, err := parseTime(event.Params[0])
	if err != nil {
		return "", err
	}
	return event.String(), e.competitorUseCase.SetStartTime(event.CompetitorID, startTime)
}
