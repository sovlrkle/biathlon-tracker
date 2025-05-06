package controllers

func (e *eventController) GenerateReport() string {
	return e.reportUseCase.GenerateReport()
}
