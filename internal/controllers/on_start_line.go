package controllers

func (e *eventController) OnStartLine(message string) (string, error) {
	event, err := parse(message)
	if err != nil {
		return "", err
	}
	return event.String(), e.competitorUseCase.OnStartLine(event.CompetitorID)
}
