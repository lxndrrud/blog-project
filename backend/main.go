package main

import (
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/lxndrrud/blog-project/routers"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	// Подключение к основной базе данных
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
	// Подключение к кэшу
	redisConn := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("CACHE_HOST") + ":" + viper.GetString("CACHE_PORT"),
		Password: viper.GetString("CACHE_PASSWORD"),
		DB:       0,
	})
	_, err = redisConn.Ping(redisConn.Context()).Result()
	if err != nil {
		log.Fatalln("Подключение к кэшу не осуществлено!" + err.Error())
	}
	defer redisConn.Close()
	app := routers.BootstrapRouter(db, redisConn)

	app.Run(":8001")
}
