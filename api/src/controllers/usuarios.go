package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
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
		if erro != nil {
			log.Fatal(erro)
		}
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		if erro != nil {
			log.Fatal(erro)
		}
	}

	db, erro := banco.Conectar()
	if erro != nil {
		if erro != nil {
			log.Fatal(erro)
		}
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioID, erro := repositorio.Criar(usuario)
	if erro != nil {
		if erro != nil {
			log.Fatal(erro)
		}
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", usuarioID)))
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