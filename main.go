// App CLI de tradução de palavras
package main

import (
	front "app/internal"

	"log"

	"github.com/joho/godotenv"
)

// executa antes do 'main()'
func init() {
	// carregar variveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar as variaveis de ambiente: %v", err)
	}
}

func main() {
	front.Run()
}
