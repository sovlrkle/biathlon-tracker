package repository

import (
	"biathlon-tracker/internal/domain"

	"github.com/pkg/errors"
)

type Repository interface {
	Get(id int) (*domain.Competitor, error)
	Add(competitor *domain.Competitor)
	GetAll() []*domain.Competitor
}

var ErrCompetitorNotFound = errors.New("competitor not found")

type competitorRepository struct {
	competitors map[int]*domain.Competitor
}

func New() *competitorRepository {
	return &competitorRepository{
		competitors: make(map[int]*domain.Competitor),
	}
}

func (r *competitorRepository) Get(id int) (*domain.Competitor, error) {
	competitor, exists := r.competitors[id]
	if !exists {
		return nil, ErrCompetitorNotFound
	}
	return competitor, nil
}

func (r *competitorRepository) Add(competitor *domain.Competitor) {
	r.competitors[competitor.ID] = competitor
}

func (r *competitorRepository) GetAll() []*domain.Competitor {
	result := make([]*domain.Competitor, 0, len(r.competitors))
	for _, competitor := range r.competitors {
		result = append(result, competitor)
	}
	return result
}
