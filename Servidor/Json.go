package Servidor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var matriz [][]string
var vec []lista

type Datos struct{
	Datos []Indice `json:"Datos"`
}

type Indice struct{
	Indice string `json:"Indice"`
	Departamentos []Departamentos `json:"Departamentos"`
}

type Departamentos struct {
	Departamentos string `json:"Nombre"`
	Tiendas []Tiendas `json:"Tiendas"`
}

type Tiendas struct {
	Tiendas string `json:"Nombre"`
	Descripcion string `json:"Descripcion"`
	Contacto string `json:"Contacto"`
	Calificacion int `json:"Calificacion"`
}

func Leer2(){
	archivo, err := os.Open("C:\\Users\\usuario\\OneDrive\\Escritorio\\primero.json")
	if err != nil{
		fmt.Println(err)
	}
	defer archivo.Close()

	valorB, _ := ioutil.ReadAll(archivo)
	var data Datos

	err = json.Unmarshal(valorB, &data)
	if err != nil{
		log.Fatal(err)
	}

	vec = make([]lista, 0, len(data.Datos[0].Departamentos)*len(data.Datos)*5)
	matriz = make([][]string, 0, len(data.Datos[0].Departamentos)*len(data.Datos))
}

func Leer(data Datos){
	vec = make([]lista, 0, len(data.Datos[0].Departamentos)*len(data.Datos)*5)
	matriz = make([][]string, 0, len(data.Datos[0].Departamentos)*len(data.Datos))
	for i := 0; i < len(data.Datos); i++ {
		fmt.Println("Indice: " + data.Datos[i].Indice)
		for j:=0; j<len(data.Datos[i].Departamentos);j++{
			fmt.Println("Departamento: " + data.Datos[i].Departamentos[j].Departamentos)
			for x:=0;x<len(data.Datos[i].Departamentos[j].Tiendas);x++{
				fmt.Println("Tienda: " + data.Datos[i].Departamentos[j].Tiendas[x].Tiendas)
				fmt.Println("Contacto: " + data.Datos[i].Departamentos[j].Tiendas[x].Descripcion)
				fmt.Println("Descripcion: " + data.Datos[i].Departamentos[j].Tiendas[x].Contacto)
				fmt.Println("Calificacion: " + strconv.Itoa(data.Datos[i].Departamentos[j].Tiendas[x].Calificacion))
			}
		}
	}
}
