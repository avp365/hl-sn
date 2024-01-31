package configs

import "github.com/avp365/hl-sn/internal/pkg/env"

func GetDBPostgrSettings() DBPostr {
	return DBPostr{
		Url: env.GetEnv("POSTGRESS_URL", ""),
	}
}
