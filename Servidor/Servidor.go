package Servidor

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func Inicial(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Funcionando EDD")
}

func agregar(w http.ResponseWriter, r *http.Request){
	var data Datos
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(body, &data)
	json.NewEncoder(w).Encode(data)
	Leer(data)
}

func Iniciar(){
	router := mux.NewRouter()
	router.HandleFunc("/", Inicial).Methods("GET")
	router.HandleFunc("/agregar", agregar).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000",router))
}