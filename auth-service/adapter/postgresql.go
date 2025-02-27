package adapter

import (
	"fmt"

	"github.com/revandpratama/reflect/auth-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresOption struct {
	db *gorm.DB
}

func (p *PostgresOption) Start(a *Adapter) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.ENV.DBHost, config.ENV.DBUser, config.ENV.DBPassword, config.ENV.DBName, config.ENV.DBPort, config.ENV.DBSSLMode)

	fmt.Println("Generated DSN:", dsn)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	a.Postgres = db
	p.db = db // Store reference for later use
	return nil
}

func (p *PostgresOption) Stop() error {
	if p.db != nil {
		sqlDB, err := p.db.DB()
		if err != nil {
			return fmt.Errorf("failed to get SQL DB from GORM: %w", err)
		}
		return sqlDB.Close()
	}
	return nil
}
