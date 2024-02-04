package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/avp365/hl-sn/internal/pkg/generator"
	"github.com/avp365/hl-sn/internal/pkg/password"
)

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	pass, err := password.HashPassword("123456")

	if err != nil {
		log.Fatal(err)
	}

	for g := 0; g < 10; g++ {
		g := g
		wg.Add(1)
		go func() {

			defer wg.Done()

			f, err := os.Create("../../.infra/srv/psql/data/users-" + strconv.Itoa(g) + ".txt")
			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			var result string

			for i := 0; i < 100000; i++ {

				user := generator.GenerateRndUser()
				fmt.Printf("user name=%s i=%d \n", user.FirstName, i)

				if err != nil {
					log.Fatal(err)
				}
				result += fmt.Sprintf("%s;%s;%s;%s;%s;%s \n", user.FirstName, user.SecondName, user.Birthdate.Format("2006-01-02"), user.Biography, user.City, pass)

				if err != nil {
					log.Fatal(err)
				}
			}
			_, err = f.WriteString(result)

			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	wg.Wait()
}
