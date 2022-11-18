// A mock-puzzle to exercise all of the common bits

package day00

import (
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {
	results := common.ParseFile(puzzle.InputFile, "\n", parseUser)

	var users []user
	for _, result := range results {
		users = append(users, result.(user))
	}

	var oldest user

	for index, user := range users {
		user.Age = float64(time.Since(user.Birthday).Hours() / 24 / 365)

		log.Debugf("User %v = { name: %v, email: %v, birthday: %v, age: %v }\n",
			index, user.Name, user.Email, user.Birthday.Format(common.ShortDateFormat), user.Age)

		if user.Age > oldest.Age {
			oldest = user
		}
	}

	answer := common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(users),
		Part2: solvePart2(users),
	}

	return answer
}

func solvePart1(users []user) string {
	log.Debug("Solving part 1.")
	var oldest user

	for _, user := range users {
		user.Age = float64(time.Since(user.Birthday).Hours() / 24 / 365)

		if user.Age > oldest.Age {
			log.Debug("Older user detected.")
			log.Tracef("oldest = { name: %v, email: %v, birthday: %v, age: %v }\n",
				user.Name, user.Email, user.Birthday.Format(common.ShortDateFormat), user.Age)

			oldest = user
		}
	}

	log.Debug("Part 1 solved.")
	return oldest.Name
}

func solvePart2(users []user) string {
	log.Debug("Solving part 2.")
	var oldest user

	for _, user := range users {
		user.Age = float64(time.Since(user.Birthday).Hours() / 24 / 365)

		if user.Age > oldest.Age {
			log.Debug("Older user detected.")
			log.Tracef("oldest = { name: %v, email: %v, birthday: %v, age: %v }\n",
				user.Name, user.Email, user.Birthday.Format(common.ShortDateFormat), user.Age)

			oldest = user
		}
	}

	log.Debug("Part 2 solved.")
	return strconv.Itoa(int(oldest.Age))
}

func parseUser(text string) interface{} {
	if text == "" {
		return nil
	}

	tokens := common.Split(text, ",")
	name := tokens[0]
	email := tokens[1]
	birthday, err := time.Parse(common.ShortDateFormat, tokens[2])

	common.Check(err)

	result := user{
		Name:     name,
		Email:    email,
		Birthday: birthday,
	}

	return result
}
