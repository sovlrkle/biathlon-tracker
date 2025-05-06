package controllers

import "fmt"

func (e *eventController) CantContinue(message string) (string, error) {
	event, err := parse(message)
	if err != nil {
		return "", err
	}

	if len(event.Params) < 1 {
		return "", fmt.Errorf("%v: only comment is required", ErrInvalidParams)
	}

	return event.String(), e.competitorUseCase.CantContinue(event.CompetitorID)
}
