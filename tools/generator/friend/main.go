package main

import (
	"fmt"
	"log"
	"os"

	"github.com/avp365/hl-sn/internal/pkg/generator"
)

func main() {

	userId := 1456

	f, err := os.Create("../../../.infra/srv/psql/pgadmin/data/friends.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var result string

	for i := 0; i < 2000; i++ {

		if err != nil {
			log.Fatal(err)
		}
		result += fmt.Sprintf("%d;%d;%s \n", userId, 1456+i, generator.GenerateRndDate(2023, 2024).Format("2006-01-02 15:04:05.00"))

		if err != nil {
			log.Fatal(err)
		}
	}
	_, err = f.WriteString(result)

	if err != nil {
		log.Fatal(err)
	}

}
