package genericInheritance

type People struct {
	Name   string
	Age    int64
	Height float64
}

type Student struct {
	Course  string
	College string
	People  People
}
