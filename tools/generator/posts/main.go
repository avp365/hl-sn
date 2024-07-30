package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/avp365/hl-sn/internal/pkg/generator"
)

func main() {

	f, err := os.Open("../../../.infra/srv/psql/pgadmin/data/friends.csv")

	if err != nil {

		log.Fatal(err)
	}

	r := csv.NewReader(f)

	pf, err := os.Create("../../../.infra/srv/psql/pgadmin/data/posts.csv")

	if err != nil {
		log.Fatal(err)
	}

	for {

		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		id := strings.Split(record[0], ";")[1]

		defer f.Close()

		var result string

		for i := 0; i < generator.RandInt(1, 30); i++ {

			if err != nil {
				log.Fatal(err)
			}
			result += fmt.Sprintf("%s;%s;%s\n", id, generator.GenerateText(500), generator.GenerateRndDate(2020, 2023).Format("2006-01-02 15:04:05.00"))

			if err != nil {
				log.Fatal(err)
			}
		}
		_, err = pf.WriteString(result)

		if err != nil {
			log.Fatal(err)
		}

	}

}
