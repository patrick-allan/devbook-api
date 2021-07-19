package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

/*
func init() {
	chave := make([]byte, 64)
	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	//fmt.Println(chave)

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)
	//usado para gerar o SECRET_KEY
}
*/

func main() {
	config.Carregar()
	r := router.Gerar()
	fmt.Printf("API is ON - Port: %d", config.Porta)
	//fmt.Println(config.SecretKey)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
