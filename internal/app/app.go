package app

import (
	"biathlon-tracker/config"
	"biathlon-tracker/internal/controllers"
	"biathlon-tracker/internal/repository"
	"biathlon-tracker/internal/usecases"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run(cfg *config.Config, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open events file: %w", err)
	}
	defer file.Close()

	repo := repository.New()
	useCase := usecases.New(repo, cfg)
	controller := controllers.New(useCase, useCase, useCase, useCase, useCase, cfg)

	var (
		output  strings.Builder
		message string
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		message = scanner.Text()

		eventId, err := getEventID(message)

		if err != nil {
			fmt.Printf("Error while parsing event ID: %v\n", err)
			continue
		}

		var res string

		handler := getHandler(eventId, controller)
		if handler == nil {
			fmt.Printf("Invalid event ID: %d\n", eventId)
			continue
		}

		res, err = handler(message)

		output.WriteString(res + "\n")
	}

	if err = scanner.Err(); err != nil {
		return fmt.Errorf("error reading events file: %w", err)
	}

	if err = writeToFile("output.log", output.String()); err != nil {
		return fmt.Errorf("failed to write to output.log: %w", err)
	}

	report := controller.GenerateReport()

	if err = writeToFile("report", report); err != nil {
		return fmt.Errorf("failed to write to report: %w", err)
	}

	return nil
}

func getHandler(eventID int, controller controllers.Controller) func(string) (string, error) {
	switch eventID {
	case 1:
		return controller.Register
	case 2:
		return controller.SetStartTime
	case 3:
		return controller.OnStartLine
	case 4:
		return controller.Start
	case 5:
		return controller.OnFiringRange
	case 6:
		return controller.HitTarget
	case 7:
		return controller.LeftFiringRange
	case 8:
		return controller.EnterPenaltyLap
	case 9:
		return controller.LeftPenaltyLap
	case 10:
		return controller.EndMainLap
	case 11:
		return controller.CantContinue
	default:
		return nil
	}
}

func getEventID(message string) (int, error) {
	parts := strings.Split(message, " ")
	if len(parts) < 2 {
		return 0, fmt.Errorf("invalid event format")
	}

	eventID, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid event ID: %w", err)
	}

	return eventID, nil
}

func writeToFile(filename, data string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err
}
