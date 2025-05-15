package main 

import "fmt"

func main() {
	var nome string = "Alice"
	var idade int = 22
	var altura float64 = 1.55
	var casada bool = true
    
	// fmt.Println("Nome:",nome)
	// fmt.Scan(&nome)

	// fmt.Println("Idade: ",idade)
    // fmt.Scan(&idade)

	// fmt.Println("Altura: ",altura)
	// fmt.Scan(&altura)

	// fmt.Println("Casada: ",casada)
	// fmt.Scan(&casada)
    
    // fmt.Println("\n--- Dados informados ---")
	// fmt.Println("Nome:", nome)
	// fmt.Println("Idade:", idade)
	// fmt.Printf("Altura: %.2f m\n", altura)
	// fmt.Println("Casada:", casada)
	fmt.Printf("Nome: %s, idade: %d, Altura: %.2f, Casada: %t\n", nome, idade, altura, casada)
}