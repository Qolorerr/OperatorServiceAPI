package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"operator_text_channel/src/routes"
	"operator_text_channel/src/services"
	"operator_text_channel/src/storage"
	"time"
)

var port = flag.String("port", "8080", "http service port")
var redisAddr = flag.String("redisAddr", "localhost:6379", "redis service address")
var redisUser = flag.String("redisUser", "testUser", "redis user name")
var redisPassword = flag.String("redisPassword", "", "redis service password")
var redisDB = flag.Int("redisDB", 0, "redis db")
var redisMaxRetries = flag.Int("redisMaxRetries", 5, "redis max retries")
var redisDialTimeout = flag.Int("redisDialTimeout", 10, "redis dial timeout in seconds")
var redisTimeout = flag.Int("redisTimeout", 5, "redis timeout in seconds")
var pgAddr = flag.String("pgAddr", "localhost", "db address")
var pgPort = flag.Int("pgPort", 5432, "db port")
var pgUser = flag.String("pgUser", "test_user", "db user name")
var pgPassword = flag.String("pgPassword", "test_password", "db user password")
var pgDbName = flag.String("pgDbName", "test_db", "db path")

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	db, err := storage.InitDB(storage.PGConfig{
		Addr:     *pgAddr,
		Port:     *pgPort,
		User:     *pgUser,
		Password: *pgPassword,
		DbName:   *pgDbName,
	})
	if err != nil {
		log.Panicln("database init error", err)
	}
	defer func() {
		if err = storage.CloseDB(db); err != nil {
			log.Println("database close error", err)
		}
	}()

	ctx := context.Background()
	rdb, err := storage.InitRedis(ctx, storage.RedisConfig{
		Addr:        *redisAddr,
		Password:    *redisPassword,
		User:        *redisUser,
		DB:          *redisDB,
		MaxRetries:  *redisMaxRetries,
		DialTimeout: time.Duration(*redisDialTimeout) * time.Second,
		Timeout:     time.Duration(*redisTimeout) * time.Second,
	})
	if err != nil {
		log.Panicln("redis init error", err)
	}

	service := services.NewDBService(ctx, db, rdb)
	router := routes.RegisterRoutes(service)

	log.Println("Starting server on " + *port)
	log.Fatal(http.ListenAndServe(":"+*port, router))
}
