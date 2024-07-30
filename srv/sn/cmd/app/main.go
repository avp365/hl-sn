package main

import (
	"github.com/avp365/hl-sn/srv/sn/internal/pkg/connect"
	"github.com/avp365/hl-sn/srv/sn/internal/repositories"
	router "github.com/avp365/hl-sn/srv/sn/internal/routers"
)

func main() {

	mg := connect.NewManager()
	mg.Init()

	repositories.InitUserRepository(mg.Conn.DBPostr, mg.Conn.DBPostrS1)
	repositories.InitPostRepository(mg.Conn.DBPostr, mg.Conn.DBPostrS1)

	router.Run()

}
