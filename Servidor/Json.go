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

	//Creando Vector E Insertando Listas con Tiendas "ROW-MAJOR"
	for i := 0; i < len(data.Datos); i++ {
		for j:=0; j<len(data.Datos[i].Departamentos);j++{
			l1 := newLista()
			l2 := newLista()
			l3 := newLista()
			l4 := newLista()
			l5 := newLista()
			for x:=0;x<len(data.Datos[i].Departamentos[j].Tiendas);x++{
				t := newTienda(data.Datos[i].Departamentos[j].Tiendas[x].Tiendas, data.Datos[i].Departamentos[j].Tiendas[x].Descripcion, data.Datos[i].Departamentos[j].Tiendas[x].Contacto)
				if data.Datos[i].Departamentos[j].Tiendas[x].Calificacion == 1 {
					insertar(t,l1)
				} else if data.Datos[i].Departamentos[j].Tiendas[x].Calificacion == 2{
					insertar(t,l2)
				} else if data.Datos[i].Departamentos[j].Tiendas[x].Calificacion == 3{
					insertar(t,l3)
				} else if data.Datos[i].Departamentos[j].Tiendas[x].Calificacion == 4{
					insertar(t,l4)
				} else if data.Datos[i].Departamentos[j].Tiendas[x].Calificacion == 5{
					insertar(t,l5)
				}
			}
			vec = append(vec, *l1)
			vec = append(vec, *l2)
			vec = append(vec, *l3)
			vec = append(vec, *l4)
			vec = append(vec, *l5)
		}
	}

	fmt.Println(vec)

	//Ordenar Valores en Listas
	for i:=0; i<len(vec);i++{
		vec[i] = vec[i].ordenar()
	}

	graficoCompleto(vec)
}

func Crear(data Datos){
	vec = make([]lista, 0, len(data.Datos[0].Departamentos)*len(data.Datos)*5)
	matriz = make([][]string, 0, len(data.Datos))

	//Creando Matriz
	for i := 0; i < len(data.Datos); i++ {
		var aux = make([]string, 0, len(data.Datos[0].Departamentos))
		for j:=0; j<len(data.Datos[i].Departamentos);j++{
			aux = append(aux, data.Datos[i].Indice + "," + data.Datos[i].Departamentos[j].Departamentos)
		}
		matriz = append(matriz, aux)
	}

	fmt.Println(matriz)

	//Creando Vector E Insertando Listas con Tiendas "ROW-MAJOR"
	for i := 0; i < len(data.Datos); i++ {
		for j:=0; j<len(data.Datos[i].Departamentos);j++{
			l1 := newLista()
			l2 := newLista()
			l3 := newLista()
			l4 := newLista()
			l5 := newLista()
			for x:=0;x<len(data.Datos[i].Departamentos[j].Tiendas);x++{
				t := newTienda(data.Datos[i].Departamentos[j].Tiendas[x].Tiendas, data.Datos[i].Departamentos[j].Tiendas[x].Descripcion, data.Datos[i].Departamentos[j].Tiendas[x].Contacto)
				if data.Datos[i].Departamentos[j].Tiendas[x].Calificacion == 1 {
					insertar(t,l1)
				} else if data.Datos[i].Departamentos[j].Tiendas[x].Calificacion == 2{
					insertar(t,l2)
				} else if data.Datos[i].Departamentos[j].Tiendas[x].Calificacion == 3{
					insertar(t,l3)
				} else if data.Datos[i].Departamentos[j].Tiendas[x].Calificacion == 4{
					insertar(t,l4)
				} else if data.Datos[i].Departamentos[j].Tiendas[x].Calificacion == 5{
					insertar(t,l5)
				}
			}
			vec = append(vec, *l1)
			vec = append(vec, *l2)
			vec = append(vec, *l3)
			vec = append(vec, *l4)
			vec = append(vec, *l5)
		}
	}

	fmt.Println(vec)

	//Ordenar Valores en Listas
	for i:=0; i<len(vec);i++{
		vec[i] = vec[i].ordenar()
	}
}

func grafico1(){
	graficoCompleto(vec)
}
