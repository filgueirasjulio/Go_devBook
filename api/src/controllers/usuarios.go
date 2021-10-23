package controllers

import (
	"net/http"
)

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando usuários..."))
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário..."))
}

func ExibirUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Exibindo usuário..."))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário..."))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário..."))
}