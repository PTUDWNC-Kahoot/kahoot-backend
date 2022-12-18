package app

import (
	"errors"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	// migrate tools
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"examples/kahootee/config"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

type mylog struct{}

func (*mylog) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Verbose should return true when verbose logging output is wanted
func (*mylog) Verbose() bool {
	return true
}

func StartMigrate() {
	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	databaseURL := cfg.PG.URL
	if databaseURL == "" {
		log.Fatalf("migrate: environment variable not declared: PG_URL")
	}

	databaseURL += "?sslmode=disable"

	for attempts > 0 {
		m, err = migrate.New("file://migrations", databaseURL)
		if err == nil {
			break
		}

		log.Printf("Migrate: postgres is trying to connect, attempts left: %d", attempts)

		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Printf("Migrate: postgres connect error: %s", err)
	}

	m.Log = &mylog{}

	defer m.Close()

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Migrate: up error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Migrate: no change")
		return
	}

	log.Printf("Migrate: up success")
}
