package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
)

var ErrInvalidConfig = errors.New("invalid config")

type Config struct {
	Laps        int    `json:"laps"`
	LapLen      int    `json:"lapLen"`
	PenaltyLen  int    `json:"penaltyLen"`
	FiringLines int    `json:"firingLines"`
	Start       string `json:"start"`
	StartDelta  string `json:"startDelta"`
	StartTime   time.Time
	Delta       time.Duration
}

func New(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(cfg); err != nil {
		return nil, err
	}

	if err = validateConfig(cfg); err != nil {
		return nil, err
	}

	cfg.StartTime, err = time.Parse("15:04:05.000", cfg.Start)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidConfig, err)
	}

	cfg.Delta, err = parseDuration(cfg.StartDelta)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidConfig, err)
	}

	return cfg, nil
}

func parseDuration(durationStr string) (time.Duration, error) {
	t, err := time.Parse("15:04:05", durationStr)
	if err != nil {
		return 0, err
	}

	return time.Duration(t.Hour())*time.Hour +
		time.Duration(t.Minute())*time.Minute +
		time.Duration(t.Second())*time.Second, nil
}

func validateConfig(cfg *Config) error {
	if cfg.Laps <= 0 {
		return fmt.Errorf("%w: Laps should be > 0, but it is: %d", ErrInvalidConfig, cfg.Laps)
	}
	if cfg.LapLen <= 0 {
		return fmt.Errorf("%w: LapLen should be > 0, but it is: %d", ErrInvalidConfig, cfg.LapLen)
	}
	if cfg.PenaltyLen <= 0 {
		return fmt.Errorf("%w: PenaltyLen should be > 0, but it is: %d", ErrInvalidConfig, cfg.LapLen)
	}
	if cfg.FiringLines <= 0 {
		return fmt.Errorf("%w: FiringLines should be > 0, but it is: %d", ErrInvalidConfig, cfg.LapLen)
	}

	return nil
}
