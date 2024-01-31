package connect

import (
	"context"

	"github.com/avp365/hl-sn/internal/pkg/configs"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

type Manager struct {
	Conn Connect
}

func NewManager() *Manager {
	return &Manager{}
}

func (mr *Manager) Init() {
	defer log.Infoln("Init Connect Manager ... Ok")
	mr.InitDBPostr()
}

func (mr *Manager) InitDBPostr() {
	connPgx, err := pgxpool.New(context.Background(), configs.GetDBPostgrSettings().Url)

	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)

	}

	mr.Conn.DBPostr = connPgx
}
