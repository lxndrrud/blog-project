package main

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	dbString := "postgres://" + viper.GetString("DB_USER") + ":" + viper.GetString("DB_PASSWORD") + "@" +
		viper.GetString("DB_HOST") + ":" + viper.GetString("DB_PORT") + "/" +
		viper.GetString("DB_NAME") + "?" +
		"sslmode=disable"
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		fmt.Println(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Println(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db//migrations",
		"postgres", driver)
	if err != nil {
		fmt.Println(err)
	}
	m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
}
