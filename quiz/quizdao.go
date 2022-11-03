package quiz

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "pgpass"
	dbname   = "toolquiz"
)

type iQuizDAO interface {
	getRandom() (tool Tool, err error)
	getHint(questionNum int) (hint string, err error)
	getAnswer(questionNum int, answer string) (correct bool, err error)
}

type quizDAOPostgres struct{}

func getQuizDAO() iQuizDAO {
	return quizDAOPostgres{}
}

func (quizDAOPostgres) getRandom() (tool Tool, err error) {
	db := getConnection()
	rows, err := db.Query("SELECT question_num, image_url FROM tool ORDER BY random() LIMIT 1")
	defer rows.Close()
	defer db.Close()
	t := Tool{}

	for rows.Next() {
		err = rows.Scan(&t.QuestionNum, &t.ImageURL)
	}
	checkError(err)
	return t, err
}

func (quizDAOPostgres) getHint(questionNum int) (hint string, err error) {
	db := getConnection()
	defer db.Close()
	err = db.QueryRow("SELECT hint FROM tool WHERE question_num = $1", questionNum).Scan(&hint)
	checkError(err)
	return hint, err
}

func (quizDAOPostgres) getAnswer(questionNum int, answer string) (correct bool, err error) {
	db := getConnection()
	rows, err := db.Query("SELECT answer FROM tool WHERE question_num = $1 and answer = $2", questionNum, strings.ToLower(answer))
	checkError(err)
	defer rows.Close()
	defer db.Close()

	correct = false
	for rows.Next() {
		correct = true
	}

	return correct, err
}

func getConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		db.Close()
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		panic(err)
	}
	return db
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
