package ient

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider"
	entgen "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/storage/ient/gen"
)

const (
	defaultRetry   = 5
	defaultTimeout = 500 * time.Millisecond
)

type PersonStorage struct {
	log *slog.Logger
	db  *entgen.Client
}

var _ provider.PersonProviderI = (*PersonStorage)(nil)

func NewDriver(cfg *Config) (*entgen.Client, error) {
	for attempt := 0; attempt < defaultRetry; attempt++ {
		if db, err := connect(cfg); err == nil {
			return db, nil
		}

		time.Sleep(defaultTimeout)
	}
	return nil, fmt.Errorf("ient: exceeded max retry limit to connect")
}

func NewPersonStorage(log *slog.Logger, db *entgen.Client) *PersonStorage {
	return &PersonStorage{log: log, db: db}
}

func connect(cfg *Config) (*entgen.Client, error) {
	return entgen.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			cfg.Hostname, cfg.Port, cfg.Username, cfg.DBName, cfg.Password))
}
