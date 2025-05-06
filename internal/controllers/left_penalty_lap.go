package controllers

func (e *eventController) LeftPenaltyLap(message string) (string, error) {
	event, err := parse(message)
	if err != nil {
		return "", err
	}
	return event.String(), e.penaltyUseCase.LeftPenaltyLap(event.CompetitorID, event.Time)
}
