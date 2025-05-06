package controllers

func (e *eventController) Register(message string) (string, error) {
	event, err := parse(message)
	if err != nil {
		return "", err
	}

	return event.String(), e.competitorUseCase.Register(event.CompetitorID, event.Time)
}
