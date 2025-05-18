package main 

import ("fmt"
"os"
"encoding/json"

_ "go-sql-driver/mysql")

type BancoDeDados struct {
	Host     string `json:"host"`
	Porta    int    `json:"porta"`
	Usuario  string `json:"usuario"`
	Senha    string `json:"senha"`
	Banco    string `json:"banco"`
	URL      string `json:"url"`
}

// Struct para a configuração da aplicação
type ConfigApp struct {
	URLFila       string      `json:"url-fila"`
	NomeFila      string      `json:"nome-fila"`
	URLFrontend   string      `json:"url-frontend"`
	BancoDeDados  BancoDeDados `json:"banco-de-dados"`
}

func GetConfig() (*ConfigApp, error) {

	var Config ConfigApp

	// Ler o arquivo de configuração
	data, err := os.ReadFile("config.json")

	if err != nil {
		fmt.Println("Erro ao ler o arquivo de configuração:", err)
		return nil, err
	}

	// Exibir o conteúdo do arquivo de configuração
	fmt.Println("Conteúdo do arquivo de configuração:", string(data))	

	// Fazer o parse do JSON
	err = json.Unmarshal(data, &Config)

	if err != nil {
		fmt.Println("Erro ao fazer o parse do JSON:", err)
		return nil, err
	}
	
	return &Config, nil	
}

func ConectaBanco() {
	// Lê o arquivo de configuração
	config, err := GetConfig()
	if err != nil {
		fmt.Println("Erro ao ler o arquivo de configuração:", err)
		return
	}

	// Conecta ao banco de dados usando as informações do arquivo de configuração
	fmt.Println("Conectando ao banco de dados...")
	fmt.Printf("Host: %s\n", config.BancoDeDados.Host)
	fmt.Printf("Porta: %d\n", config.BancoDeDados.Porta)
	fmt.Printf("Usuário: %s\n", config.BancoDeDados.Usuario)
	fmt.Printf("Senha: %s\n", config.BancoDeDados.Senha)
	fmt.Printf("Banco: %s\n", config.BancoDeDados.Banco)

	//db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")



}
