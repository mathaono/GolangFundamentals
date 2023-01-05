package crud

import (
	"log"
	"net/http"

	"github.com/GolangFundamentals/server"
	"github.com/gorilla/mux"
)

func Crud() {

	//Contem todas as rotas para interação com o banco de dados
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", server.CreateUser).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":4000", router))
}
