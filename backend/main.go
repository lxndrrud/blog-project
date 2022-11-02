package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/lxndrrud/blog-project/routers"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	db, err := sqlx.Connect(
		"postgres",
		"host="+viper.GetString("DB_HOST")+" "+
			"user="+viper.GetString("DB_USER")+" "+
			"password="+viper.GetString("DB_PASSWORD")+" "+
			"dbname="+viper.GetString("DB_NAME")+" "+
			"sslmode=disable")
	if err != nil {
		log.Fatalln("Подключение к базе данных не осуществлено!" + err.Error())
	}
	app := routers.BootstrapRouter(db)

	app.Run(":8001")
}
