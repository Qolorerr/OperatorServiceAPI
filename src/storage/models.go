package storage

import "time"

type RedisConfig struct {
	Addr        string
	Password    string
	User        string
	DB          int
	MaxRetries  int
	DialTimeout time.Duration
	Timeout     time.Duration
}

type PGConfig struct {
	Addr     string
	Port     int
	User     string
	Password string
	DbName   string
}
