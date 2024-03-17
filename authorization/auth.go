package authorization

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/BlackAWhite8/vkTest/bd"
	e "github.com/BlackAWhite8/vkTest/errors"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

var Flag = false
var Priv = false

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	u := User{}
	_ = json.NewDecoder(r.Body).Decode(&u)

	conn := bd.Connection()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		e.DBConnError(err)
	}
	defer db.Close()

	rows, err := db.Query("select login, pass,is_admin from web_user where login=$1 and pass=$2", u.Login, u.Password)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusBadRequest)
	}
	counter := 0
	check := User{}
	for rows.Next() {
		err = rows.Scan(&check.Login, &check.Password, &Priv)
		if err != nil {
			e.RowsDataError(err)
		}
		counter += 1
	}
	if counter == 1 {
		log.Println("auth complete succesfully")
		Flag = true
		http.Redirect(w, r, "/", http.StatusOK)
	} else {
		fmt.Fprintf(w, "invalid login or password")
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u := User{}
	_ = json.NewDecoder(r.Body).Decode(&u)
	conn := bd.Connection()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		e.DBConnError(err)
	}
	defer db.Close()

	rows, err := db.Query("select login from web_user where login=$1", u.Login)
	if err != nil {
		e.DBGetInfoError(err)
	}

	counter := 0
	for rows.Next() {
		counter += 1
	}
	if counter == 0 {
		result, err := db.Exec("insert into web_user (login, pass) values ($1,$2)", u.Login, u.Password)
		if err != nil {
			e.DBPostInfoError(err)
		}
		fmt.Println(result.RowsAffected())
	} else {
		fmt.Fprintf(w, "this login is already exists, please choose another one")
	}
}
