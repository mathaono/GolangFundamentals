package httpProtocol

import (
	"html/template"
	"log"
	"net/http"
)

var tmpls *template.Template

type Usuario struct {
	Nome  string
	Email string
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Rota de usuários dentro da rota home"))
}

func UpServer() {

	tmpls = template.Must(template.ParseGlob("*.html"))

	user1 := Usuario{"Matheus", "aono.1304@gmail.com"}

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		tmpls.ExecuteTemplate(w, "home.html", user1)
	})
	http.HandleFunc("/home/users", users)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
