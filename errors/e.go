package errors

import "log"

func DBConnError(e error) {
	log.Fatal("can't connect to the database, error occurs:", e, "check db needed")
}

func DBGetInfoError(e error) {
	log.Println("can't get the information from database, error occurs:", e, "check db needed")
}

func DBPostInfoError(e error) {
	log.Println("can't update/delete/insert information due to error", e, "check db and code needed")
}

func RowsDataError(e error) {
	log.Println("something went wrong while extracting data from bd", e, "check code needed")
}

func EncodeError(e error) {
	log.Panicln("can't encode data properly", e)
}
