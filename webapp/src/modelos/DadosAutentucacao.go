package modelos

// DadosAutenticacao contém o id e token do usuário
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
