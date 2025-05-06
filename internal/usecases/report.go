package usecases

import (
	"biathlon-tracker/internal/domain"
	"fmt"
	"sort"
	"strings"
	"time"
)

func (e *eventUseCase) GenerateReport() string {
	competitors := e.repo.GetAll()
	sort.Slice(competitors, func(i, j int) bool {
		return competitors[i].StartTime.Before(competitors[j].StartTime)
	})

	var report strings.Builder
	for _, competitor := range competitors {
		var status string
		if competitor.Status == domain.NotFinished || competitor.Status == domain.Disqualified {
			status = competitor.Status.String()
		} else {
			status = e.calcDiffStartTime(competitor)
		}

		laps := e.calcLapsData(competitor)

		penaltyTime, penaltySpeed := e.calcPenalty(competitor)
		penaltyStr := e.formatPenalty(penaltyTime, penaltySpeed)

		shotsStat := fmt.Sprintf("%d/%d", competitor.Hits, competitor.Hits+competitor.Misses)

		report.WriteString(fmt.Sprintf(
			"[%s] %d [%s] %s %s\n",
			status,
			competitor.ID,
			strings.Join(laps, ", "),
			penaltyStr,
			shotsStat,
		))
	}
	return report.String()
}

func (e *eventUseCase) formatPenalty(time time.Duration, speed float64) string {
	if time > 0 {
		return fmt.Sprintf("{%s, %.3f}", e.formatDuration(time), speed)
	}
	return "{,}"
}

func (e *eventUseCase) calcPenalty(c *domain.Competitor) (time.Duration, float64) {
	totalPenalty := time.Duration(0)
	for _, pt := range c.PenaltyDurations {
		totalPenalty += pt
	}

	if totalPenalty > 0 && c.Misses > 0 {
		penaltyDistance := e.cfg.PenaltyLen * c.Misses
		return totalPenalty, float64(penaltyDistance) / totalPenalty.Seconds()
	}
	return 0, 0.0
}

func (e *eventUseCase) calcLapsData(c *domain.Competitor) []string {
	var res []string
	for _, lap := range c.LapDurations {
		res = append(res, fmt.Sprintf("{%s, %.3f}", e.formatDuration(lap), float64(e.cfg.LapLen)/lap.Seconds()))
	}
	for len(res) < e.cfg.Laps {
		res = append(res, "{,}")
	}
	return res
}

func (e *eventUseCase) calcDiffStartTime(competitor *domain.Competitor) string {
	switch competitor.Status {
	case domain.NotFinished, domain.Disqualified:
		return ""
	default:
		return e.formatDuration(competitor.StartTime.Sub(competitor.PlannedStartTime))
	}
}

func (e *eventUseCase) formatDuration(d time.Duration) string {
	return fmt.Sprintf("%02d:%02d:%02d.%03d", int(d.Hours()), int(d.Minutes())%60, int(d.Seconds())%60, int(d.Milliseconds())%1000)
}
