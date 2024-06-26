package main

import (
	"github.com/avp365/hl-sn/internal/pkg/connect"
	"github.com/avp365/hl-sn/internal/repositories"
	router "github.com/avp365/hl-sn/internal/routers"
)

func main() {

	mg := connect.NewManager()
	mg.Init()

	repositories.InitUserRepository(mg.Conn.DBPostr, mg.Conn.DBPostrS1)

	router.Run()

}
