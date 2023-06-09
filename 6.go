package main

//Crie uma struct chamada Funcionário
//com os campos "nome", "salário" e "idade".
//Escreva funções que permitam aumentar ou diminuir o
//salário do funcionário em uma determinada
//porcentagem e uma função que calcule o tempo
//de serviço do funcionário (considerando que ele começou
//a trabalhar aos 18 anos).

type Funcionario struct {
	nome    string
	salario float64
	idade   int
}

func increment(f Funcionario, percentual float64) Funcionario {
	f.salario = f.salario + (f.salario * percentual)
	return f
}

func decrement(f Funcionario, percentual float64) Funcionario {
	f.salario = f.salario - (f.salario * percentual)
	return f
}

func calcularTempoServico(f Funcionario) int {
	return f.idade - 18
}

func main() {

}
