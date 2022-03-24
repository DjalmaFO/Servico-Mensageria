package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func RodarRotina() {
	for {
		args := new(ArgumentosEmail)
		erro := new(ArgumentosEmail)
		erro.Assunto = "Serviço de Mensageria - Falha ao enviar email"
		erro.TipoTexto = "text"
		erro.Texto = "Falha ao enviar email \n"
		r, err := ObterItemFila()
		if err != nil {
			log.Println(err.Error())
			erro.Texto += fmt.Sprintf("Motivo: Falha ao acessar fila.\nMensagem de erro: %s", err.Error())
			erro.EnviarEmail(config.EmailError)
		}

		if len(r) > 0 {
			if err := json.Unmarshal([]byte(r), args); err != nil {
				log.Println(err.Error())
				erro.Texto += fmt.Sprintf("Motivo: Falha ao montar objeto.\nMensagem de erro: %s", err.Error())
				erro.EnviarEmail(config.EmailError)
			}

			for _, v := range args.Para {
				if err := args.EnviarEmail(v); err != nil {
					log.Println(err.Error())
					erro.Texto += fmt.Sprintf("Motivo: Falha ao montar email.\nMensagem de erro: %s\nDestinatário: %s", err.Error(), v)
					erro.EnviarEmail(config.EmailError)
				}
			}
		} else {
			time.Sleep(1 * time.Minute)
		}
	}
}
