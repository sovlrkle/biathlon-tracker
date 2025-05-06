package domain

import (
	"time"
)

type status int

const (
	Registered status = iota + 1
	StartTimeSet
	OnStartLine
	Started
	OnFiringRange
	HitTargets
	LeftFiringRange
	PenaltyLap
	LeftPenaltyLap
	NotFinished
	Disqualified
)

func (s status) String() string {
	switch s {
	case NotFinished:
		return "NotFinished"
	case Disqualified:
		return "Disqualified"
	default:
		return ""
	}
}

type (
	Competitor struct {
		ID               int
		Status           status
		RegistrationTime time.Time
		PlannedStartTime time.Time
		StartTime        time.Time
		//LastEventTime    time.Time
		lapInfo
		penaltyInfo
		firingRangeInfo
	}

	lapInfo struct {
		CurrentLap    int
		LapStartTimes []time.Time
		LapDurations  []time.Duration
	}

	penaltyInfo struct {
		PenaltyStartTime     time.Time
		PenaltyDurations     []time.Duration
		TotalPenaltyDuration time.Duration
	}

	firingRangeInfo struct {
		CurrentFiringRange int
		Targets            map[int]bool
		Hits               int
		Misses             int
	}
)

func NewCompetitor(id int) *Competitor {
	return &Competitor{
		ID:     id,
		Status: Registered,
		lapInfo: lapInfo{
			LapStartTimes: make([]time.Time, 0),
			LapDurations:  make([]time.Duration, 0),
		},
		penaltyInfo: penaltyInfo{
			PenaltyDurations: make([]time.Duration, 0),
		},
	}
}
