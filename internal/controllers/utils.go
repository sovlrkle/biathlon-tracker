package controllers

import (
	"biathlon-tracker/internal/domain"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parse(message string) (*domain.Event, error) {
	parts := strings.Split(message, "] ")
	if len(parts) < 2 {
		return nil, ErrInvalidEvent
	}

	t, err := parseTime(strings.Trim(parts[0], "["))
	if err != nil {
		return nil, err
	}

	parts = strings.Fields(parts[1])
	if len(parts) < 2 {
		return nil, ErrInvalidEvent
	}

	id, err := parseInt(parts[0])
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidEvent, err)
	}

	competitorID, err := parseInt(parts[1])
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidEvent, err)
	}

	params := parts[2:]
	return domain.NewEvent(t, id, competitorID, params), nil
}

func parseTime(s string) (time.Time, error) {
	return time.Parse("15:04:05.000", s)
}

func parseInt(s string) (int, error) {
	return strconv.Atoi(s)
}
