package configs

import "time"

type DBPostr struct {
	Url             string
	MaxOpenConn     int
	ConnMaxLifeTime time.Duration
}
