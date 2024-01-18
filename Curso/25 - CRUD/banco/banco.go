package banco

import (
	
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func criarConexao() (*sql.DB, error) {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	driverName := "mysql"

	db, erro := sql.Open(driverName, stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro := db.Ping(); erro != nil {
		return nil, erro
	}

	fmt.Println("Conexão com o db realizada com sucesso")
	return db, nil
}

func Salvar(usuario model.Usuario) (int64, error) {
	db, erro := criarConexao()
	if erro != nil {
		return 0, errors.New("erro ao conectar ao banco de dados")
	}

	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO usuarios (nome, email) VALUES (?, ?)")
	if erro != nil {
		return 0, errors.New("erro ao criar o statement")
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		return 0, errors.New("erro ao executar o statement")
	}

	idInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, errors.New("erro ao obter o id inserido")
	}

	return idInserido, nil
}

func Buscar() ([]model.Usuario, error) {
	var usuarios []model.Usuario

	db, erro := criarConexao()
	if erro != nil {
		return usuarios, errors.New("erro ao conectar ao banco de dados")
	}

	defer db.Close()

	rows, erro := db.Query("SELECT * FROM usuarios")
	if erro != nil {
		return usuarios, errors.New("erro ao criar o statement")
	}

	defer rows.Close()

	for rows.Next() {
		var usuario model.Usuario

		if erro := rows.Scan(&usuario.Id, &usuario.Nome, &usuario.Email); erro != nil {
			return usuarios, errors.New("erro ao escanear usuários")
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func BuscarUsuarioPorId(id uint64) (model.Usuario, error) {
	var usuario model.Usuario

	db, erro := criarConexao()
	if erro != nil {
		return usuario, errors.New("erro ao conectar ao banco de dados")
	}

	defer db.Close()

	row, erro := db.Query("SELECT * FROM usuarios WHERE id = ?", id)
	if erro != nil {
		return usuario, errors.New("erro ao criar statement")
	}

	defer row.Close()

	if row.Next() {
		if erro := row.Scan(&usuario.Id, &usuario.Nome, &usuario.Email); erro != nil {
			return usuario, errors.New("erro ao escanear usuário")
		}
	}

	return usuario, nil
}

func AtualizarUsuario(id uint64, usuarioAtualizado model.Usuario) error {
	db, erro := criarConexao()
	if erro != nil {
		return errors.New("erro ao criar conexão com DB")
	}
	defer db.Close()

	stmt, erro := db.Prepare("UPDATE usuarios SET nome = ?, email = ? WHERE id = ?")
	if erro != nil {
		return errors.New("erro ao criar statement")
	}
	defer stmt.Close()

	_, erro = stmt.Exec(usuarioAtualizado.Nome, usuarioAtualizado.Email, id)
	if erro != nil {
		return errors.New("erro ao atualizar usuario")
	}

	return erro
}

func DeletarUsuario(id uint64) error {
	db, erro := criarConexao()
	if erro != nil {
		return errors.New("erro ao criar conexão com DB")
	}
	defer db.Close()

	stmt, erro := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if erro != nil {
		return errors.New("erro ao criar statement")
	}
	defer stmt.Close()

	_, erro = stmt.Exec(id)
	if erro != nil {
		return errors.New("erro ao deletar usuario")
	}

	return erro
}
