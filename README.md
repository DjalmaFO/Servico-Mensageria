# Serviço de Mensageria
>Este serviço consiste em receber requisições REST e adicionar uma solicitação de envio de e-mail em fila no Redis em caso de sucesso. Do contrário retorna resposta de erro.

## Autor
>Djalma Freire de Oliveira

### Tecnologias
- Golang
- Redis 
- echo Framwork (https://echo.labstack.com/)

### Configuração do .env
- Copie o arquivo **.env.example** e o renomeie para **.env**
>Variáveis

- SERVER_PORT: porta onde será esperada a requisição
- MAIL_SERVER: servidor da conta de e-mail
- MAIL_PORT: porta do servidor da conta
- MAIL_USER: usuário de login da conta
- MAIL_PASSWORD: senha da conta
- MAIL_FROM= e-mail utilizado para envio de e-mails
- MAIL_ERROR= e-mail que irá receber as notificações de erro
- MAIL_TEST= e-mail que irá receber email de teste de integração da aplicação
- ADDR_REDIS= url servidor Redis
- PASSWORD_REDIS= senha de acesso ao Redis

### End Point
- /mensageria
- Method: POST
- Body: JSON

### Retorno de status
- 406 = JSON inválido
- 400 = Problemas com a validação dos campos recebidos ou Redis
- 200 = Recebido e adicionado à fila

### Rodar testes
>go test
