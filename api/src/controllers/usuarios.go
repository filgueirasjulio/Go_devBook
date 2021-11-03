package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando usuários..."))
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	
		corpoRequest, erro := ioutil.ReadAll(r.Body) 
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	repositorio.Criar(usuario)
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