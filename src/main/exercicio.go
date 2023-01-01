package main

import "fmt"


type Estudante struct {
	matricula string
	nome string
	sobrenome string
	idade int
	turno string
	turma string
}

func (estudante Estudante) detalhes_alunos(){
	fmt.Println("Matricula:",estudante.matricula)
	fmt.Println("Estudante:",estudante.nome,estudante.sobrenome)
	fmt.Println("idade:",estudante.idade)
	fmt.Println("turno:",estudante.turno)
	fmt.Println("turma:",estudante.turma)
} 



func main() {
	estudante1 := Estudante {"001","Junildo","Aroldo",18,"manhã","C"}
	estudante1.detalhes_alunos()
}



