package sql

import "time"

type Database struct {
	Username     string
	Password     string
	Host         string
	Port         string
	DBName       string
	Param1       string
	Param2       string
	Param3       string
	MaxOpenConns int
	MaxIdleConns int
	MaxLifetime  time.Duration
}
