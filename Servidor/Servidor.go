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

//Cargar tiendas En Json
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

//Generar Grafos
func getArreglo(w http.ResponseWriter, _ *http.Request){
	grafico1()
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusCreated)
}

//Buscar Tienda Especifica
func tiendaE(w http.ResponseWriter, r *http.Request){
	var t unico
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusFound)
	if vec!=nil{
		_ = json.Unmarshal(body, &t)
		a := posiciont(t)
		if a.Calificacion != 0{
			_ = json.NewEncoder(w).Encode(a)
		}else{
			_ = json.NewEncoder(w).Encode("No Se Encontro Ninguna Tienda Que Cumpla Con Los Parametros")
		}
	}else {
		_ = json.NewEncoder(w).Encode("No Hay Tiendas Cargadas")
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
	var t unico2
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusOK)
	_ = json.Unmarshal(body, &t)
	if vec != nil{
		a:= Eliminar(t)
		var v varios
		_ = json.Unmarshal(a, &v)
		_ = json.NewEncoder(w).Encode(v)
	}else{
		_ = json.NewEncoder(w).Encode("No Hay Tiendas Cargadas")
	}
}

func guardar(w http.ResponseWriter, _ *http.Request){
	if vec != nil{
		var regreso Datos
		a := retorno(0,0)
		w.Header().Set("Content-Type","applicattion/json")
		w.WriteHeader(http.StatusAccepted)
		_ = json.Unmarshal(a, &regreso)
		_ = json.NewEncoder(w).Encode(regreso)
	}else{
		_ = json.NewEncoder(w).Encode("No Hay Tiendas Cargadas")
	}
}

func Iniciar(){
	router := mux.NewRouter()
	router.HandleFunc("/guardar", guardar).Methods("GET")
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	router.HandleFunc("/id/{numero}", tiendaN).Methods("GET")
	router.HandleFunc("/cargartienda", cargar).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", tiendaE).Methods("POST")
	router.HandleFunc("/Eliminar", elim).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000",router))
}