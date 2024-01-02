package config

import "fmt"

type IDbConfig interface {
	Url() string
	MaxOpenConns() int
}

type db struct {
	host           string
	port           int
	protocol       string
	username       string
	password       string
	database       string
	sslMode        string
	maxConnections int
}

func (a *db) Url() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		a.host,
		a.port,
		a.username,
		a.password,
		a.database,
		a.sslMode,
	)
}

func (a *db) MaxOpenConns() int { return a.maxConnections }
