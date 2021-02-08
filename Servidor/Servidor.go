package Servidor

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Inicial(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Funcionando EDD")
}

func agregar(w http.ResponseWriter, r *http.Request){
	/*body, err := ioutil.ReadAll(r.Body)
	if err != nill{
		fmt.Fprintf(w, "Error al insertar")
	}*/
}

func Iniciar(){
	router := mux.NewRouter()
	router.HandleFunc("/", Inicial).Methods("GET")
	router.HandleFunc("/agregar", agregar).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000",router))
}