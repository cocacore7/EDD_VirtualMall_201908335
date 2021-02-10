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
	fmt.Fprintf(w, "Servidor Funcionando")
}

func cargar(w http.ResponseWriter, r *http.Request){
	var data Datos
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(body, &data)
	json.NewEncoder(w).Encode(data)
	Crear(data)
}

func getArreglo(w http.ResponseWriter, r *http.Request){
	
}

func Iniciar(){
	router := mux.NewRouter()
	router.HandleFunc("/", Inicial).Methods("GET")
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	router.HandleFunc("/cargartienda", cargar).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000",router))
}