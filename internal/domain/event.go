package domain

import (
	"fmt"
	"strings"
	"time"
)

type Event struct {
	Time         time.Time
	ID           int
	CompetitorID int
	Params       []string
}

func NewEvent(time time.Time, id int, competitorID int, params []string) *Event {
	return &Event{
		Time:         time,
		ID:           id,
		CompetitorID: competitorID,
		Params:       params,
	}
}

func (e *Event) String() string {
	switch e.ID {
	case 1:
		return fmt.Sprintf("[%s] The competitor(%d) registered", e.Time.Format("15:04:05.000"), e.CompetitorID)
	case 2:
		return fmt.Sprintf("[%s] The start time for the competitor(%d) was set by a draw to %s", e.Time.Format("15:04:05.000"), e.CompetitorID, e.Params[0])
	case 3:
		return fmt.Sprintf("[%s] The competitor(%d) is on the start line", e.Time.Format("15:04:05.000"), e.CompetitorID)
	case 4:
		return fmt.Sprintf("[%s] The competitor(%d) has started", e.Time.Format("15:04:05.000"), e.CompetitorID)
	case 5:
		return fmt.Sprintf("[%s] The competitor(%d) is on the firing range(%s)", e.Time.Format("15:04:05.000"), e.CompetitorID, e.Params[0])
	case 6:
		return fmt.Sprintf("[%s] The target(%s) has been hit by competitor(%d)", e.Time.Format("15:04:05.000"), e.Params[0], e.CompetitorID)
	case 7:
		return fmt.Sprintf("[%s] The competitor(%d) left the firing range", e.Time.Format("15:04:05.000"), e.CompetitorID)
	case 8:
		return fmt.Sprintf("[%s] The competitor(%d) entered the penalty laps", e.Time.Format("15:04:05.000"), e.CompetitorID)
	case 9:
		return fmt.Sprintf("[%s] The competitor(%d) left the penalty laps", e.Time.Format("15:04:05.000"), e.CompetitorID)
	case 10:
		return fmt.Sprintf("[%s] The competitor(%d) ended the main lap", e.Time.Format("15:04:05.000"), e.CompetitorID)
	case 11:
		return fmt.Sprintf("[%s] The competitor(%d) can`t continue: %s", e.Time.Format("15:04:05.000"), e.CompetitorID, strings.Join(e.Params, " "))
	//case 32:
	//	return fmt.Sprintf("[%s] The competitor(%d) is disqualified", e.Time.Format("15:04:05.000"), e.CompetitorID)
	//case 33:
	//	return fmt.Sprintf("[%s] The competitor(%d) has finished", e.Time.Format("15:04:05.000"), e.CompetitorID)
	default:
		return fmt.Sprintf("[%s] Unknown event(%d) for competitor(%d)", e.Time.Format("15:04:05.000"), e.ID, e.CompetitorID)
	}
}
