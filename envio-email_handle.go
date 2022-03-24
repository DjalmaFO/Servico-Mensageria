package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"gopkg.in/gomail.v2"
)

type ArgumentosEmail struct {
	Para      []string `json:"para"`
	Assunto   string   `json:"assunto"`
	TipoTexto string   `json:"tipo_texto"`
	Texto     string   `json:"texto"`
	Anexo     string   `json:"anexo"`
	CC        string   `json:"cc"`
}

func (a *ArgumentosEmail) ValidarEmails() (ok bool) {
	for _, email := range a.Para {
		if !ValidaFormatoEmail(email) {
			return
		}
	}
	return true
}

func (a ArgumentosEmail) ValidarTipoTexto() (ok bool) {
	return a.TipoTexto == "text" || a.TipoTexto == "html"
}

func (a *ArgumentosEmail) ValidarDados() error {
	if !a.ValidarEmails() {
		return fmt.Errorf("Formato de email inválido")
	}

	if !a.ValidarTipoTexto() {
		return fmt.Errorf("Tipo de texto inválido. Valor permitido => (text ou html)")
	}

	return nil
}

func (a *ArgumentosEmail) EnviarEmail(destinatario string) (err error) {
	log.Println("Formatando e-mail")
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", config.MailFrom)

	// Set E-Mail receivers
	m.SetHeader("To", destinatario)

	if len(a.CC) > 0 {
		m.SetHeader("Cc", a.CC)
	}
	// m.SetHeader("Bcc", "office@example.com", "anotheroffice@example.com")

	// Set E-Mail subject
	m.SetHeader("Subject", a.Assunto)

	// Set E-Mail body. You can set plain text or html with text/html
	switch a.TipoTexto {
	case "text":
		m.SetBody("text/plain", a.Texto)
	case "html":
		m.SetBody("text/html", a.Texto)
	}

	if len(a.Anexo) > 0 {
		m.Attach(a.Anexo)
	}

	port, err := strconv.Atoi(config.MailPort)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("Falha ao converter porta para inteiro: Verificar .env")
	}

	// Settings for SMTP server
	d := gomail.NewDialer(config.MailServer, port, config.MailUser, config.MailPass)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Save E-Mail in mymail.txt file

	// Get directory where binary is started
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }

	// Write contents of E-Mail into mymail.txt.
	// This is useful for debuging.
	// var b bytes.Buffer
	// m.WriteTo(&b)
	// err = ioutil.WriteFile(dir+`mymail.txt`, b.Bytes(), 0777)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }

	log.Printf("Enviando email para %s\n", destinatario)

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		return err
	} else {
		log.Println("Email enviado com sucesso")
	}

	return nil
}

func (a *ArgumentosEmail) AdicionarFila() (err error) {
	str, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return AdicionarFila(string(str))
}
