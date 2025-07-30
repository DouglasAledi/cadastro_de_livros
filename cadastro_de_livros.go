package main

import (
	"fmt"
	"os"
)

type infoLivros struct {
	nome string
}

func main() {
	var nomeLivro string
	var conjuntoLivros []infoLivros
	var opcao string

	escolha(nomeLivro, conjuntoLivros, opcao)
}

func adicionarLivro(conjuntoLivros []infoLivros, nomeLivro string) []infoLivros {
	if nomeLivro != "" {
		novoLivro := infoLivros{nome: nomeLivro}
		conjuntoLivros = append(conjuntoLivros, novoLivro)
	} else {
		fmt.Println("Nome vazio")
	}
	return conjuntoLivros
}

func excluirLivro(conjuntoLivros []infoLivros, nome string) []infoLivros {
	for i, livro := range conjuntoLivros {
		if livro.nome == nome {
			fmt.Println("Livro removido com sucesso:", livro.nome)
			return append(conjuntoLivros[:i], conjuntoLivros[i+1:]...)
		}
	}
	fmt.Println("Livro não encontrado.")
	return conjuntoLivros
}

func escolha(nomeLivro string, conjuntoLivros []infoLivros, opcao string) []infoLivros {
	for {
		fmt.Println("\nEscolha uma das opções: ")
		fmt.Println("1 - Adicionar livro")
		fmt.Println("2 - Ver lista de livros")
		fmt.Println("3 - Excluir livro")
		fmt.Println("4 - Sair")
		fmt.Print("Opção: ")
		fmt.Scan(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Insira o nome do livro: ")
			fmt.Scan(&nomeLivro)
			conjuntoLivros = adicionarLivro(conjuntoLivros, nomeLivro)

		case "2":
			if len(conjuntoLivros) == 0 {
				fmt.Println("Nenhum livro na biblioteca.")
				break
			}
			fmt.Println("Livros na biblioteca:")
			for i, livro := range conjuntoLivros {
				fmt.Printf("%d - %s\n", i+1, livro.nome)
			}

		case "3":
			fmt.Print("Digite o nome do livro a ser removido: ")
			fmt.Scan(&nomeLivro)
			conjuntoLivros = excluirLivro(conjuntoLivros, nomeLivro)

		case "4":
			fmt.Println("Saindo...")
			os.Exit(0)

		default:
			fmt.Println("Erro, comando não identificado.")
		}
	}
}
