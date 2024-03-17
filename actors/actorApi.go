package actors

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BlackAWhite8/vkTest/bd"
	e "github.com/BlackAWhite8/vkTest/errors"
	_ "github.com/lib/pq"
)

type Actor struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Sex      string   `json:"sex"`
	Birthday string   `json:"birthday"`
	Movies   []string `json:"movies"`
}

func GetActors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	conn := bd.Connection()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		e.DBConnError(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from actor")
	if err != nil {
		e.DBGetInfoError(err)
	}
	defer rows.Close()

	actors := []Actor{}
	m := ""
	for rows.Next() {
		a := Actor{}

		err := rows.Scan(&a.Id, &a.Name, &a.Sex, &a.Birthday)
		if err != nil {
			e.RowsDataError(err)
		}

		rows2, err2 := db.Query("select title from movie_actors where actor_name=$1 ", a.Name)
		if err2 != nil {
			e.DBGetInfoError(err2)
		}
		defer rows2.Close()

		for rows2.Next() {
			err := rows2.Scan(&m)
			if err != nil {
				e.RowsDataError(err)
			}
			a.Movies = append(a.Movies, m)
		}
		actors = append(actors, a)
	}
	err = json.NewEncoder(w).Encode(actors)
	if err != nil {
		e.EncodeError(err)
	}
}

func CreateActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	a := Actor{}
	_ = json.NewDecoder(r.Body).Decode(&a)
	/*проблема при заходе на страницу сайта с админ правами
	соединение теряется тк формат даты в виде пустой строки выдает ошибку, через постман все ок*/
	if a.Birthday == "" || a.Name == "" || a.Sex == "" {
		return
	}

	conn := bd.Connection()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		e.DBConnError(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into actor (name, sex, birthday) values ($1,$2,$3)", a.Name, a.Sex, a.Birthday)
	if err != nil {
		e.DBPostInfoError(err)
	}
	fmt.Println(result.RowsAffected())
}

func UpdateActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	a := Actor{}
	_ = json.NewDecoder(r.Body).Decode(&a)

	n := r.URL.Query().Get("name")
	s := r.URL.Query().Get("sex")

	conn := bd.Connection()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		e.DBConnError(err)
	}
	defer db.Close()

	result, err := db.Exec("update actor set name=$1, sex=$2, birthday=$3 where name=$4 and sex=$5", a.Name, a.Sex, a.Birthday, n, s)
	if err != nil {
		e.DBPostInfoError(err)
	}
	fmt.Println(result.RowsAffected())
}

func DeleteActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	a := Actor{}
	_ = json.NewDecoder(r.Body).Decode(&a)
	/*проблема при заходе на страницу сайта с админ правами
	соединение теряется тк формат даты в виде пустой строки выдает ошибку, через постман все ок*/
	if a.Birthday == "" {
		return
	}

	conn := bd.Connection()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		e.DBConnError(err)
	}
	defer db.Close()

	result, err := db.Exec("delete from actor where name=$1 and sex=$2 and birthday=$3", a.Name, a.Sex, a.Birthday)
	if err != nil {
		e.DBPostInfoError(err)
	}
	fmt.Println(result.RowsAffected())
}
