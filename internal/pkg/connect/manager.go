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
	mr.InitDBPostrS1()
}

func (mr *Manager) InitDBPostr() {
	connPgx, err := pgxpool.New(context.Background(), configs.GetDBPostgrSettings().Url)

	if err != nil {
		log.Printf("Unable to connect to database master: %v\n", err)

	}

	mr.Conn.DBPostr = connPgx
}
func (mr *Manager) InitDBPostrS1() {
	connPgx, err := pgxpool.New(context.Background(), configs.GetDBPostgrS1Settings().Url)

	if err != nil {
		log.Printf("Unable to connect to database s1: %v\n", err)

	}

	mr.Conn.DBPostrS1 = connPgx
}
