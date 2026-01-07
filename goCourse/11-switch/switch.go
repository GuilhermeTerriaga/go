package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Segunda"
	case 2:
		return "Terça"
	case 3:
		return "Quarta"
	case 4:
		return "Quinta"
	case 5:
		return "Sexta"
	case 6:
		return "Sábado"
	case 7:
		return "Domingo"
	default:
		return "Entrada inválida"
	}
}

// assim também é válido mas principalmente quando se quer avaliar mais de uma variável no switch
func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Segunda"
	case numero == 2:
		return "Terça"
	case numero == 3:
		return "Quarta"
	case numero == 4:
		return "Quinta"
	case numero == 5:
		return "Sexta"
	case numero == 6:
		return "Sábado"
	case numero == 7:
		return "Domingo"
	default:
		return "Entrada inválida"
	}
}
func diaDaSemana3(numero int) string {
	var diaSemana string
	switch {
	case numero == 1:
		diaSemana = "Segunda"
		fallthrough // fallthrough faz com que o Go não avalie a próxima condição, ele aceita a próxima condição como válida
	case numero == 2:
		diaSemana = "Terça"
	case numero == 3:
		diaSemana = "Quarta"
	case numero == 4:
		diaSemana = "Quinta"
	case numero == 5:
		diaSemana = "Sexta"
	case numero == 6:
		diaSemana = "Sábado"
	case numero == 7:
		diaSemana = "Domingo"
	default:
		diaSemana = "Entrada inválida"
	}
	return diaSemana
}
func main() {
	fmt.Println("Switch")
	diaSemana := diaDaSemana(3)
	fmt.Println(diaSemana)
	fmt.Println(diaDaSemana(200))
	fmt.Println(diaDaSemana2(300))
	fmt.Println(diaDaSemana3(1)) // vai retornar terça
}
