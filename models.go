package main

type BancoDeDados struct {
	Host    string `json:"host"`
	Porta   int    `json:"porta"`
	Usuario string `json:"usuario"`
	Senha   string `json:"senha"`
	Banco   string `json:"banco"`
	URL     string `json:"url"`
}

// Struct para a configuração da aplicação
type ConfigApp struct {
	URLFila      string       `json:"url-fila"`
	NomeFila     string       `json:"nome-fila"`
	URLFrontend  string       `json:"url-frontend"`
	BancoDeDados BancoDeDados `json:"banco-de-dados"`
}

type Exames struct {
	IdExame    int    `json:"id_exame"`
	TipoExame  string `json:"tipo_exame"`
	Instrucoes string `json:"instrucoes"`
}

type Agendamentos struct {
	DataHora      string `json:"data_hora"`
	IdExame       int    `json:"id_exame"`
	Instrucoes    string `json:"instrucoes"`
	Paciente      string `json:"nome_paciente"`
	EmailPaciente string `json:"email_paciente"`
	CpfPaciente   string `json:"cpf"`
	CartaoSus     string `json:"cartao_sus"`
}

type AgendamentoEmail struct {
	DataHora      string `json:"data_hora"`
	NomeExame     string `json:"nome_exame"`
	Instrucoes    string `json:"instrucoes"`
	Paciente      string `json:"nome_paciente"`
	EmailPaciente string `json:"email_paciente"`
}

type RecuperaAgendamento struct {
	NomePaciente 		string `json:"nome_paciente"`
	EmailPaciente 		string `json:"email_paciente"` 
	Data 				string `json:"data_hoa"`
	IdExame 			int    `json:"id_exame"` 
	TipoExame 			string `json:"tipo_exame"`
	Instrucoes 			string `json:"instrucoes"`
	Cpf 				string `json:"cpf"`
	CartaoSus 			string `json:"cartao_sus"`
}