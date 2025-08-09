package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Nome  string
	Idade int
}

type Profissao struct {
	Nome    string
	Area    string
	Salario float64
}

type Funcionario struct {
	User      User
	Profissao Profissao
}

func main() {
	reader := bufio.NewReader(os.Stdin)  //Cria um leitor de entrada padrão - input - e é mais flexivel que o fmt.Scan()

	// Cria um slice que é tipo um array dinâmico
	var funcionarios []Funcionario

	fmt.Print("Quantos funcionários deseja cadastrar? ")
	qtdStr, _ := reader.ReadString('\n')
	qtdStr = strings.TrimSpace(qtdStr)
	qtd, _ := strconv.Atoi(qtdStr)

	for i := 0; i < qtd; i++ {
		fmt.Printf("\nCadastro do funcionário %d:\n", i+1)

		fmt.Print("Nome: ")
		nome, _ := reader.ReadString('\n')
		nome = strings.TrimSpace(nome)

		fmt.Print("Idade: ")
		idadeStr, _ := reader.ReadString('\n') //Lê até a quebra de linha
		idadeStr = strings.TrimSpace(idadeStr) //remove espaços e quebras de linha
		idade, _ := strconv.Atoi(idadeStr)   //parecido com ParseInt

		fmt.Print("Cargo: ")
		cargo, _ := reader.ReadString('\n')
		cargo = strings.TrimSpace(cargo)

		fmt.Print("Área: ")
		area, _ := reader.ReadString('\n')
		area = strings.TrimSpace(area)

		fmt.Print("Salário: ")
		salarioStr, _ := reader.ReadString('\n')
		salarioStr = strings.TrimSpace(salarioStr)
		salario, _ := strconv.ParseFloat(salarioStr, 64)

		funcionario := Funcionario{
			User:      User{Nome: nome, Idade: idade},
			Profissao: Profissao{Nome: cargo, Area: area, Salario: salario},
		}

		funcionarios = append(funcionarios, funcionario)
	}

	fmt.Println("\n--- Relatório de Funcionários ---")

	// Sempre tipar o índice e o valor no for range para evitar problemas
	for _, f := range funcionarios {
		fmt.Printf("Nome: %s | Idade: %d | Cargo: %s | Área: %s | Salário: R$ %.2f\n",
			f.User.Nome, f.User.Idade, f.Profissao.Nome, f.Profissao.Area, f.Profissao.Salario)
	}
}
