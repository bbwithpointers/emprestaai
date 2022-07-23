package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
	"strconv"
)

var (
	// StringConnDB String de conexão com banco mysql
	StringConnDB = ""
	// Port - porta onde a API esta rodando
	Port = 0
	// SecretKey - chave para assinar o token jwt
	SecretKey []byte
)

// Carregar inicializa as variaveis de ambiente
func Carregar() {
	var erro error
	// o .env deve estar no MAIN do pacote, por exemplo, aqui esta dentro de API
	// se colocado fora de API, derá erro, por default.
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 8080
	}
	StringConnDB = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"), os.Getenv("DB_SENHA"), os.Getenv("DB_NOME"))
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
