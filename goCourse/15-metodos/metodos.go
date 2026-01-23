package main

import "fmt"
type usuario struct{
	nome string
	sobrenome string
	idade uint8
}

func (u usuario) salvar(){ //isso é um método de usuário. é importante que seja nessa estrutura.
	fmt.Printf("Usuário %s, salvo! ", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() { // utilizando um ponteiro do usuário
	u.idade ++
}

func main(){
	usuario1:= usuario{"Zé", "Sobrenome do zé", 42}
	fmt.Println(usuario1)
	usuario1.salvar()
	fmt.Println(usuario1.maiorDeIdade())
	fmt.Printf("usuario1.idade: %v\n", usuario1.idade)
	usuario1.fazerAniversario()
	fmt.Printf("usuario1.idade: %v\n", usuario1.idade)
}
