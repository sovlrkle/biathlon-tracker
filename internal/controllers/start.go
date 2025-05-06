package controllers

func (e *eventController) Start(message string) (string, error) {
	event, err := parse(message)
	if err != nil {
		return "", err
	}
	return event.String(), e.competitorUseCase.Start(event.CompetitorID, event.Time)
}
