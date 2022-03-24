package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	EmailError string
	MailServer string
	MailPort   string
	MailUser   string
	MailPass   string
	MailFrom   string
	MailTest   string
}

func (c *Config) Configurar() {
	var ok bool

	err := godotenv.Load(envFilename)
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Arquivo .env não encontrado!")
	}

	c.ServerPort, ok = os.LookupEnv("SERVER_PORT")
	if !ok {
		log.Fatalf("Variável SERVER_PORT não definida")
	}

	c.EmailError, ok = os.LookupEnv("MAIL_ERROR")
	if !ok {
		log.Fatalf("Variável MAIL_ERROR não definida")
	}

	c.MailServer, ok = os.LookupEnv("MAIL_SERVER")
	if !ok {
		log.Fatalf("Variável MAIL_SERVER não definida")
	}

	c.MailPort, ok = os.LookupEnv("MAIL_PORT")
	if !ok {
		log.Fatalf("Variável MAIL_PORT não definida")
	}
	c.MailUser, ok = os.LookupEnv("MAIL_USER")
	if !ok {
		log.Fatalf("Variável MAIL_USER não definida")
	}
	c.MailPass, ok = os.LookupEnv("MAIL_PASSWORD")
	if !ok {
		log.Fatalf("Variável MAIL_PASSWORD não definida")
	}
	c.MailFrom, ok = os.LookupEnv("MAIL_FROM")
	if !ok {
		log.Fatalf("Variável MAIL_FROM não definida")
	}
	c.MailTest, ok = os.LookupEnv("MAIL_TEST")
	if !ok {
		log.Fatalf("Variável MAIL_TEST não definida")
	}
}
