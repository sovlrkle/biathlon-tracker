package controllers

func (e *eventController) EnterPenaltyLap(message string) (string, error) {
	event, err := parse(message)
	if err != nil {
		return "", err
	}

	return event.String(), e.penaltyUseCase.EnterPenaltyLap(event.CompetitorID, event.Time)
}
