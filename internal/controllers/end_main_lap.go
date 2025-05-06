package controllers

func (e *eventController) EndMainLap(message string) (string, error) {
	event, err := parse(message)
	if err != nil {
		return "", err
	}
	return event.String(), e.lapUseCase.EndLap(event.CompetitorID, event.Time)
}
