package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuário da Rede Social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Preparar vai validar e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O Nome é obrigatório!")
	}

	if usuario.Nick == "" {
		return errors.New("O Nick é obrigatório!")
	}

	if usuario.Email == "" {
		return errors.New("O E-mail é obrigatório!")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A Senha é obrigatória!")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
