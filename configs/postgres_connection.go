package configs

import (
	"log"
	"os"
)

/*
PostgresConnection struct
Storing postgres connection informations.
*/
type PostgresConnection struct {
	host     string
	user     string
	password string
	dbName   string
	port     string
}

/*
LoadPostgres function
Loading postgres connection information's from environment variables.
*/
func LoadPostgres() PostgresConnection {
	return PostgresConnection{
		host:     os.Getenv("DB_HOST"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbName:   os.Getenv("DB_NAME"),
		port:     os.Getenv("DB_PORT"),
	}
}

/*
GetDSN function
Returning postgres connection string.
*/
func (p PostgresConnection) GetDSN() string {
	dsn := "host=" + p.host + " user=" + p.user + " password=" + p.password + " dbname=" + p.dbName + " port=" + p.port
	log.Println("Postgres DSN:" + dsn)
	return dsn
}
