package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/GolangFundamentals/database"
)

type user struct {
	ID    int64  `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha na request"))
		return
	}

	var user user

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		w.Write([]byte("Erro ao converter usuario"))
		return
	}

	fmt.Println(user)

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	statement, err := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	insert, err := statement.Exec(user.Nome, user.Email)
	if err != nil {
		w.Write([]byte("Erro ao executar statement"))
	}

	idInsert, err := insert.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter ID"))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário criado, ID: %d", idInsert)))
}
