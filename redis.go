package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

func abrirRedis() *redis.Client {
	err := godotenv.Load(envFilename)
	if err != nil {
		log.Fatal(err)
	}

	addr, ok := os.LookupEnv("ADDR_REDIS")
	if !ok {
		log.Fatal("Variável 'ADDR_REDIS' não foi definida!")
	}

	pass, ok := os.LookupEnv("PASSWORD_REDIS")
	if !ok {
		log.Fatal("Variável 'PASSWORD_REDIS' não foi definida!")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0, // use default DB
	})

	return rdb
}

func AdicionarFila(value string) error {
	redis := abrirRedis()
	if redis == nil {
		return fmt.Errorf("Falha ao conectar ao redis")
	}
	defer redis.Close()

	redis.LPush("fila-3223", value)
	return nil
}

func ObterItemFila() (string, error) {
	redis := abrirRedis()
	if redis == nil {
		return "", fmt.Errorf("Falha ao conectar ao redis")
	}
	defer redis.Close()

	r := redis.BRPopLPush("fila-3223", "", time.Second*1)

	return r.Val(), nil
}
