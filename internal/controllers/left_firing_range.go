package controllers

func (e *eventController) LeftFiringRange(message string) (string, error) {
	event, err := parse(message)
	if err != nil {
		return "", err
	}

	return event.String(), e.firingRangeUseCase.LeftFiringRange(event.CompetitorID)
}
