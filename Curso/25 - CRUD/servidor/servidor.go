package servidor

import (
	"crud/banco"
	"crud/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		gerarRetorno("Falha ao ler o corpo da requisição", http.StatusBadRequest, w)
		return
	}

	var usuario model.Usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		gerarRetorno("Erro ao converter json para struct", http.StatusInternalServerError, w)
		return
	}

	idInserido, erro := banco.Salvar(usuario)
	if erro != nil {
		gerarRetorno(erro.Error(), http.StatusInternalServerError, w)
		return
	}

	gerarRetorno(fmt.Sprintf("Usuário inserido com sucesso. Id: %d", idInserido), http.StatusCreated, w)
}

func gerarRetorno(mensagem string, status int, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Write([]byte(mensagem))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	usuarios, erro := banco.Buscar()
	if erro != nil {
		gerarRetorno(erro.Error(), http.StatusInternalServerError, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		gerarRetorno("erro ao marshal usuarios", http.StatusInternalServerError, w)
	}

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	id, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	usuario, erro := banco.BuscarUsuarioPorId(id)
	if erro != nil {
		gerarRetorno(erro.Error(), http.StatusInternalServerError, w)
	}

	w.Header().Add("Content-Type", "application/json")
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		gerarRetorno("erro ao marshal usuario", http.StatusInternalServerError, w)
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		gerarRetorno("erro ao ler body", http.StatusBadRequest, w)
		return
	}

	var usuarioAtualizado model.Usuario
	if erro := json.Unmarshal(body, &usuarioAtualizado); erro != nil {
		gerarRetorno("erro ao unmarshall body", http.StatusInternalServerError, w)
		return
	}

	params := mux.Vars(r)

	id, erro := strconv.ParseUint(params["id"], 0, 0)
	if erro != nil {
		gerarRetorno("erro ao converter id", http.StatusBadRequest, w)
		return
	}

	if erro := banco.AtualizarUsuario(id, usuarioAtualizado); erro != nil {
		gerarRetorno("erro ao atualizar usuário", http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, erro := strconv.ParseUint(params["id"], 0, 0)
	if erro != nil {
		gerarRetorno("erro ao ler id", http.StatusBadRequest, w)
		return
	}

	banco.DeletarUsuario(id)
}
