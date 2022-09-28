package jsons

type Usuario struct {
	Nome  string `json:"nome"`
	Idade int64  `json:"idade"`
}

type Cursos struct {
	Curso     string `json:"curso"`
	Faculdade string `json:"faculdade"`
}
