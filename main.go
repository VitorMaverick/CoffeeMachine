package main

import . "fmt"

type Cafe struct {
	qtdAgua     int
	qtdLeite    int
	qtdCafe     int
	tipo        string
	descartavel int
	dinheiro    int
}

func venderCafe(estoque Cafe) Cafe {
	var cafe Cafe
	var command int
	Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu::")
	Scan(&command)
	switch command {
	case 1:
		cafe = geraCafeEspresso()

	case 2:
		cafe = geraCafeLatte()
	case 3:
		cafe = geraCafeCappuccino()
	default:
		return estoque
	}

	if cafe.qtdAgua <= estoque.qtdAgua && cafe.qtdLeite <= estoque.qtdLeite && cafe.qtdCafe <= estoque.qtdCafe && estoque.descartavel > 0 {
		estoque.qtdAgua -= cafe.qtdAgua
		estoque.qtdLeite -= cafe.qtdLeite
		estoque.qtdCafe -= cafe.qtdCafe
		estoque.descartavel -= 1
		estoque.dinheiro += cafe.dinheiro
		Println("I have enough resources, making you a coffee!")
	} else if cafe.qtdAgua > estoque.qtdAgua {
		Println("Sorry, not enough water!")
	} else if cafe.qtdLeite > estoque.qtdLeite {
		Println("Sorry, not enough milk!")
	} else if cafe.qtdCafe > estoque.qtdCafe {
		Println("Sorry, not enough coffee beans!")
	} else if estoque.descartavel < 0 {
		Println("Sorry, not enough disposable cups")
	}
	return estoque
}

func printEstoque(estoque Cafe) {
	Println("The coffee machine has:")
	Printf("%v of water\n", estoque.qtdAgua)
	Printf("%v of milk\n", estoque.qtdLeite)
	Printf("%v of coffee beans\n", estoque.qtdCafe)
	Printf("%v of disposable cups\n", estoque.descartavel)
	Printf("%v of money\n", estoque.dinheiro)
}

func retirada(estoque Cafe) Cafe {
	Printf("I gave you %v", estoque.dinheiro)
	estoque.dinheiro = 0
	return estoque

}

// geradores do objeto café com suas determinadas quantidades
func geraCafeEspresso() Cafe {
	var copoCafe Cafe
	copoCafe.qtdAgua = 250
	copoCafe.qtdLeite = 0
	copoCafe.qtdCafe = 16
	copoCafe.dinheiro = 4
	return copoCafe
}

func geraCafeLatte() Cafe {
	var copoCafe Cafe
	copoCafe.qtdAgua = 350
	copoCafe.qtdLeite = 75
	copoCafe.qtdCafe = 20
	copoCafe.dinheiro = 7
	return copoCafe
}

func geraCafeCappuccino() Cafe {
	var copoCafe Cafe
	copoCafe.qtdAgua = 200
	copoCafe.qtdLeite = 100
	copoCafe.qtdCafe = 12
	copoCafe.dinheiro = 6
	return copoCafe
}

func geraEstoque() Cafe {
	var estoqueCafe Cafe
	estoqueCafe.qtdAgua = 400
	estoqueCafe.qtdLeite = 540
	estoqueCafe.qtdCafe = 120
	estoqueCafe.descartavel = 9
	estoqueCafe.dinheiro = 550
	return estoqueCafe
}

func recarregaEstoque(cafe Cafe) Cafe {
	var agua, leite, grao, descartaveis int
	Println("Write how many ml of water you want to add:")
	Scan(&agua)
	Println("Write how many ml of milk you want to add:")
	Scan(&leite)
	Println("Write how many grams of coffee beans you want to add:")
	Scan(&grao)
	Println("Write how many disposable coffee cups you want to add:")
	Scan(&descartaveis)

	cafe.qtdAgua += agua
	cafe.qtdLeite += leite
	cafe.qtdCafe += grao
	cafe.descartavel += descartaveis

	return cafe
}

//calcula quantos copos consegue fazer
func calculaQtdCopos(estoque Cafe, copoCafe Cafe) int {
	var copos Cafe
	copos.qtdAgua = estoque.qtdAgua / copoCafe.qtdAgua
	copos.qtdLeite = estoque.qtdLeite / copoCafe.qtdLeite
	copos.qtdCafe = estoque.qtdCafe / copoCafe.qtdCafe

	return calculaMenorQuantidadeDeIngridientes(copos.qtdAgua, copos.qtdLeite, copos.qtdCafe)
}

//
func calculaMenorQuantidadeDeIngridientes(num1 int, num2 int, num3 int) int {
	itemComMenorQuantidade := 0
	if num1 < num2 {
		itemComMenorQuantidade = num1
	} else {
		itemComMenorQuantidade = num2
	}
	if itemComMenorQuantidade > num3 {
		itemComMenorQuantidade = num3
	}
	return itemComMenorQuantidade
}

func main() {
	// write your code here
	var input string
	estoque := geraEstoque()
	deveContinuar := 0

	for deveContinuar == 0 {
		Println("Write action (buy, fill, take, remaining, exit):")
		Scan(&input)
		switch input {
		case "buy":
			estoque = venderCafe(estoque)
		case "fill":
			estoque = recarregaEstoque(estoque)
		case "take":
			estoque = retirada(estoque)
		case "remaining":
			printEstoque(estoque)
		case "exit":
			deveContinuar = 1

		}
	}

}
