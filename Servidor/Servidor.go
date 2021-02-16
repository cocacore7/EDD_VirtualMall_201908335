package Servidor

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var data Datos

func cargar(w http.ResponseWriter, r *http.Request){
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
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusCreated)
}

func tiendaE(w http.ResponseWriter, r *http.Request){
	var t unico
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusFound)
	_ = json.Unmarshal(body, &t)
	a := posiciont(t)
	if a.Calificacion != 0{
		_ = json.NewEncoder(w).Encode(a)
	}else{
		_ = json.NewEncoder(w).Encode("No Se Encontro Ninguna Tienda Que Cumpla Con Los Parametros")
	}

}

func tiendaN(w http.ResponseWriter, r *http.Request){
	var v varios
	vars:=mux.Vars(r)
	b,_:=strconv.Atoi(vars["numero"])
	a := posicionl(b)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusFound)
	_ = json.Unmarshal(a, &v)
	_ = json.NewEncoder(w).Encode(v)
}

func elim(w http.ResponseWriter, r *http.Request){

}

func Iniciar(){
	router := mux.NewRouter()
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	router.HandleFunc("/id/{numero}", tiendaN).Methods("GET")
	router.HandleFunc("/cargartienda", cargar).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", tiendaE).Methods("POST")
	router.HandleFunc("/Eliminar", elim).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000",router))
}