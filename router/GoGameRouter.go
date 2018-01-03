package router

import (
	"net/http"
	"fmt"
	"github.com/alexseph/gogame/api"
	"log"
)

func routerGetNextPS4GamesDates(w http.ResponseWriter, r *http.Request) {
	//Aqui já escreveu, nao tem o html completo, mas aqui o navegador ja mostra
	fmt.Fprintf(w, api.GetNextPS4GamesDates(10))
}

//StartRouter
func StartRouter() {
	//a funcao que vai responder a esse path
	http.HandleFunc("/ps4releasedates", routerGetNextPS4GamesDates)

	log.Println("Executando na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
