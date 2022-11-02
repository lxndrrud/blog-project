package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var globalDB *sql.DB

func connectMigrations() *migrate.Migrate {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	dbString := "postgres://" + viper.GetString("DB_USER") + ":" + viper.GetString("DB_PASSWORD") + "@" +
		viper.GetString("DB_HOST") + "/" +
		viper.GetString("DB_NAME") + "?" +
		"sslmode=disable"
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatalln(err)
	}
	globalDB = db
	//defer db.Close()
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalln(err)
	}
	return m
}

func runTestSeeds() {
	_, err := globalDB.Exec(`
			INSERT INTO "public"."roles" (title) VALUES 
				('Администратор');

			INSERT INTO "public"."permissions" (title) VALUES 
				('Создавать посты');

			INSERT INTO "public"."roles_permissions" (id_role, id_permission) VALUES 
				(1, 1);

			INSERT INTO "public"."users" (login, password, id_role) VALUES 
				('admin', '123456', 1);
		`)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// Пропарсить командную строку
	var migrationsFlag string
	var seedsFlag string

	flag.StringVar(&migrationsFlag, "m", "migrations", "Direction of migration (up or down).")
	flag.StringVar(&seedsFlag, "s", "seeds", "Direction of seeds after migrations (up or down).")
	flag.Parse()

	// Подключиться к миграциям
	migrations := connectMigrations()
	defer globalDB.Close()
	if migrationsFlag == "up" {
		fmt.Println("Migrations up!")
		err := migrations.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
		if err != nil {
			fmt.Println(err)
		}
	} else if migrationsFlag == "down" {
		fmt.Println("Migrations down!")
		err := migrations.Down()
		if err != nil {
			fmt.Println(err)
		}
	}

	if seedsFlag == "y" {
		fmt.Println("Run seeds!")
		runTestSeeds()
	}
}
