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
				('Администратор'), ('Блогер');

			INSERT INTO "public"."permissions" (title, code) VALUES 
				('Создавать посты', 'СОЗД_ПОСТЫ'), 
				('Модерировать посты', 'МОДЕР_ПОСТЫ');

			INSERT INTO "public"."roles_permissions" (id_role, id_permission) VALUES 
				(1, 1), (1, 2), (2, 1);

			INSERT INTO "public"."users" (login, password, id_role) VALUES 
				('admin', '$2a$10$sE/CMuBOJ4UP8oqRNYtCu.EvDKcsX.C853Nx2sRTgONzdKMGyAUOu', 1),
				('bloger', '$2a$10$sE/CMuBOJ4UP8oqRNYtCu.EvDKcsX.C853Nx2sRTgONzdKMGyAUOu', 2);

			INSERT INTO "public"."posts" 
			(title, annotation, text, id_author, picture, views, approved, time_start, time_end) VALUES
			('test title', 'test annotation', 'test text', 1, '/diretide.jpeg', DEFAULT, DEFAULT, DEFAULT, '2023-01-01T10:00:00'),
			('test title 2', 'test annotation', 'test text 2', 1, '/diretide.jpeg', DEFAULT, TRUE, DEFAULT, '2023-01-01T10:00:00'),
			('test title 3', 'test annotation', 'test text 3', 1, '/diretide.jpeg', DEFAULT, TRUE, DEFAULT, '2023-01-01T10:00:00');
		`)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// Пропарсить командную строку
	var migrationsFlag string
	var seedsFlag string

	flag.StringVar(&migrationsFlag, "m", "migrations", "Direction of migration ('up' or 'down').")
	flag.StringVar(&seedsFlag, "s", "seeds", "Direction of seeds after migrations (only 'y' if needed).")
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
