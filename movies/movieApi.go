package movies

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BlackAWhite8/vkTest/bd"
	e "github.com/BlackAWhite8/vkTest/errors"
	_ "github.com/lib/pq"
)

type Movie struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	RealiseDate string `json:"realise_date"`
	Score       int    `json:"score"`
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	m := Movie{}
	_ = json.NewDecoder(r.Body).Decode(&m)
	/*проблема при заходе на страницу сайта с админ правами
	соединение теряется тк формат даты в виде пустой строки выдает ошибку, через постман все ок*/
	if m.RealiseDate == "" {
		return
	}
	conn := bd.Connection()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		e.DBConnError(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into movie (title, description, realise_date, score) values ($1,$2,$3,$4)", m.Title, m.Description, m.RealiseDate, m.Score)
	if err != nil {
		e.DBPostInfoError(err)
	}
	fmt.Println(result.RowsAffected())
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	m := Movie{}
	_ = json.NewDecoder(r.Body).Decode(&m)

	t := r.URL.Query().Get("title")

	conn := bd.Connection()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		e.DBConnError(err)
	}
	defer db.Close()

	result, err := db.Exec("update movie set title=$1, description=$2, realise_date=$3, score=$4 where title=$5", m.Title, m.Description, m.RealiseDate, m.Score, t)
	if err != nil {
		e.DBPostInfoError(err)
	}
	fmt.Println(result.RowsAffected())
}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	m := Movie{}
	_ = json.NewDecoder(r.Body).Decode(&m)
	/*проблема при заходе на страницу сайта с админ правами
	соединение теряется тк формат даты в виде пустой строки выдает ошибку, через постман все ок*/
	if m.RealiseDate == "" {
		return
	}
	conn := bd.Connection()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		e.DBConnError(err)
	}
	defer db.Close()

	result, err := db.Exec("delete from movie where title=$1 and description=$2 and realise_date=$3 and score=$4", m.Title, m.Description, m.RealiseDate, m.Score)
	if err != nil {
		e.DBPostInfoError(err)
	}
	fmt.Println(result.RowsAffected())
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sortBy := r.URL.Query().Get("sort_by")
	sortOrder := r.URL.Query().Get("order")
	if sortOrder == "" {
		sortOrder = "desc"
	}
	if sortBy == "" {
		sortBy = "score"
	}

	conn := bd.Connection()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		e.DBConnError(err)
	}
	defer db.Close()

	q := "select * from movie order by " + sortBy + " " + sortOrder
	rows, err := db.Query(q)
	if err != nil {
		e.DBGetInfoError(err)
	}
	defer rows.Close()

	movies := []Movie{}

	for rows.Next() {
		m := Movie{}
		err := rows.Scan(&m.Id, &m.Title, &m.Description, &m.RealiseDate, &m.Score)
		if err != nil {
			e.RowsDataError(err)
		}
		movies = append(movies, m)
	}
	err = json.NewEncoder(w).Encode(movies)
	if err != nil {
		e.EncodeError(err)
	}
}

func SearchMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	movies := []Movie{}

	conn := bd.Connection()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		e.DBConnError(err)
	}
	defer db.Close()

	pattern := "%" + r.URL.Query().Get("search") + "%"
	rows, err := db.Query("select * from movie where title like $1", pattern)
	if err != nil {
		e.DBGetInfoError(err)
	}
	defer rows.Close()

	for rows.Next() {
		m := Movie{}
		err := rows.Scan(&m.Id, &m.Title, &m.Description, &m.RealiseDate, &m.Score)
		if err != nil {
			e.RowsDataError(err)
		}
		movies = append(movies, m)
	}
	err = json.NewEncoder(w).Encode(movies)
	if err != nil {
		e.EncodeError(err)
	}
}
