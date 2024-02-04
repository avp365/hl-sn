package generator

import (
	"math/rand"
	"time"

	"github.com/avp365/hl-sn/internal/entities"
	"github.com/avp365/hl-sn/internal/pkg/rufaker"
)

func GenerateRndUser() entities.User {
	user := entities.User{}
	firstName, secondName := generateRndName(randInt(0, 2))

	user.FirstName = firstName
	user.SecondName = secondName
	user.Birthdate = generateRndDate(1940, 2008)
	user.Biography = generateBio(randInt(0, 3))
	user.City = rufaker.GenerateCity()
	user.Password = "Password123456"

	return user

}

func generateRndName(isMen int) (string, string) {

	if isMen == 1 {
		return rufaker.FirstNameMale(), rufaker.LastNameMale()
	}

	return rufaker.FirstNameFemale(), rufaker.LastNameFemale()
}

func generateRndDate(startYear int, endYear int) time.Time {
	min := time.Date(startYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(endYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func generateBio(number int) string {
	b := make(map[int]string)

	b[0] = "А .... моя биография 1. моя биография 1. моя биография 1. моя биография 1. моя биография 1. моя биография 1. моя биография 1.моя биография 1. моя биография 1.моя биография 1. моя биография 1, моя биография 1. моя биография 1. моя биография 1. моя биография 1"
	b[1] = "Б .... моя биография 2, моя биография 2, моя биография 2, моя биография 2, моя биография 2, моя биография 2, моя биография 2, моя биография 2, моя биография 2, моя биография 2, моя биография 2, моя биография 2, моя биография 2, моя биография 2, моя биография 2"
	b[2] = "С .... моя биография 3, моя биография 3, моя биография 3 моя биография 3 моя биография 3 моя биография 3 моя биография 3 моя биография 3 моя биография 3 моя биография 3 моя биография 3 моя биография 3 моя биография 3 моя биография 3 моя биография 3 моя биография 2"

	return b[number]
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
