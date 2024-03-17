package bd

import (
	"os"

	"github.com/joho/godotenv"
)

func Connection() string {
	err := godotenv.Load("config/.env")
	if err != nil {
		panic(err)
	}
	user, _ := os.LookupEnv("POSTGRES_USER")
	db, _ := os.LookupEnv("POSTGRES_DB")
	pass, _ := os.LookupEnv("POSTGRES_PASSWORD")
	mode, _ := os.LookupEnv("MODE")
	conn := "user=" + user + " password=" + pass + " dbname=" + db + " sslmode=" + mode
	return conn
}
