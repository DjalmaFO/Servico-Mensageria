package main

import (
	"fmt"
	"testing"
)

func TestValidarTipoTexto(t *testing.T) {
	arg := new(ArgumentosEmail)

	validador := []Validator{
		{"html", true},
		{"", false},
		{"text", true},
		{"texto", false},
	}

	for _, v := range validador {
		arg.TipoTexto = fmt.Sprintf("%v", v.Valor)
		if v.Esperado != arg.ValidarTipoTexto() {
			t.Fail()
		}
	}
}

func TestValidarEmails(t *testing.T) {
	arg := new(ArgumentosEmail)

	validador := []Validator{
		{"asdfg@adfg.com", true},
		{"", false},
		{"adfsdf@fdsfas.com.br", true},
		{"sadfsadfdfdfkkkkfdsa@nndfsndfna", false},
		{"sadfsdfssdjsjh.com.br", false},
	}

	for _, v := range validador {
		arg.Para = []string{fmt.Sprintf("%v", v.Valor)}
		if v.Esperado != arg.ValidarEmails() {
			t.Fail()
		}
	}
}

func TestEnviarEmail(t *testing.T) {
	arg := new(ArgumentosEmail)

	arg.Assunto = "Não Responda - Teste de Funcionalidade"
	arg.TipoTexto = "text"
	arg.Texto = "Teste de Funcionalidade do serviço de mensageria."

	if err := arg.EnviarEmail(config.MailTest); err != nil {
		t.Fail()
	}
}
