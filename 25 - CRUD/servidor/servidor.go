package servidor

import (
	"crud/banco"
	"crud/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	}

	w.Header().Add("Content-Type", "application/json")
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		gerarRetorno("erro ao marshal usuarios", http.StatusInternalServerError, w)
	}

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

}
