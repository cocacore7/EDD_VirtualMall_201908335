package Servidor

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func cargar(w http.ResponseWriter, r *http.Request){
	var data Datos
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.Unmarshal(body, &data)
	_ = json.NewEncoder(w).Encode(data)
	Crear(data)
}

func getArreglo(w http.ResponseWriter, r *http.Request){
	grafico1()
}

func tiendaE(w http.ResponseWriter, r *http.Request){

}

func tiendaN(w http.ResponseWriter, r *http.Request){

}

func elim(w http.ResponseWriter, r *http.Request){

}

func Iniciar(){
	router := mux.NewRouter()
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	router.HandleFunc("/id/:numero", tiendaN).Methods("GET")
	router.HandleFunc("/cargartienda", cargar).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", tiendaE).Methods("POST")
	router.HandleFunc("/Eliminar", elim).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000",router))
}