package configs

import "github.com/avp365/hl-sn/srv/sn/internal/pkg/env"

func GetDBPostgrSettings() DBPostr {
	return DBPostr{
		Url:         env.GetEnv("POSTGRESS_URL", ""),
		MaxOpenConn: 20,
	}
}

func GetDBPostgrS1Settings() DBPostr {
	return DBPostr{
		Url:         env.GetEnv("POSTGRESS_SLAVE_1_URL", ""),
		MaxOpenConn: 20,
	}
}

func GetTokenSecretKey() string {
	return env.GetEnv("TOKEN_SECRET_KEY", "")
}

func GetRedisUrl() string {
	return env.GetEnv("REDIS_URL", "")
}
