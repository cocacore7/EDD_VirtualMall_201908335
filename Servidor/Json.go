package Servidor

import (
	"encoding/json"
	"fmt"
)

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

type varios struct {
	Tiendas []Tiendas `json:"Tiendas"`
}

type unico struct {
	Departamento string `json:"Departamento"`
	Tienda string `json:"Nombre"`
	Calificacion int `json:"Calificacion"`
}

func Crear(data Datos){
	vec = make([]lista, 0, len(data.Datos[0].Departamentos)*len(data.Datos)*5)

	//Creando Vector E Insertando Listas con Tiendas "ROW-MAJOR"
	for i := 0; i < len(data.Datos); i++ {
		for j:=0; j<len(data.Datos[i].Departamentos);j++{
			l1 := newLista()
			l2 := newLista()
			l3 := newLista()
			l4 := newLista()
			l5 := newLista()
			for x:=0;x<len(data.Datos[i].Departamentos[j].Tiendas);x++{
				t := newTienda(data.Datos[i].Departamentos[j].Tiendas[x].Tiendas, data.Datos[i].Departamentos[j].Tiendas[x].Descripcion, data.Datos[i].Departamentos[j].Tiendas[x].Contacto, data.Datos[i].Departamentos[j].Tiendas[x].Calificacion)
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
	reportes(vec,0,0,1)
}

func posicionl(i int) []byte{
	if vec != nil{
		if vec[i-1].Vacio(){
			crear_json, _ := json.Marshal("Lista Vacia")
			return crear_json
		}else{
			var v varios
			a := vec[i-1].primero
			for a != nil{
				t := Tiendas{Tiendas:a.tienda.nombre,Descripcion:a.tienda.descripcion,Contacto:a.tienda.contacto,Calificacion:a.tienda.calif}
				v.Tiendas = append(v.Tiendas, t)
				a = a.sig
			}
			crear_json, _ := json.Marshal(v)
			return crear_json
		}
	}else{
		crear_json, _ := json.Marshal("No Hay Tiendas Cargadas")
		return crear_json
	}
}

func posiciont(t unico) Tiendas{
	var ti Tiendas
	for i := 0; i < len(data.Datos); i++ {
		for j:=0; j<len(data.Datos[i].Departamentos);j++{
			if data.Datos[i].Departamentos[j].Departamentos == t.Departamento{
				for z:=0; z<len(data.Datos[i].Departamentos[j].Tiendas);z++{
					if data.Datos[i].Departamentos[j].Tiendas[z].Tiendas == t.Tienda{
						if data.Datos[i].Departamentos[j].Tiendas[z].Calificacion == t.Calificacion{
							ti = data.Datos[i].Departamentos[j].Tiendas[z]
						}
					}
				}
			}
		}
	}
	return ti
}
