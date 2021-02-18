package Servidor

import (
	"encoding/json"
)

var vec []lista

//Estructura tiendas json
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

//Estructura complementaria
type varios struct {
	Tiendas []Tiendas `json:"Tiendas"`
}

//Estructura Busqueda
type unico struct {
	Departamento string `json:"Departamento"`
	Tienda string `json:"Nombre"`
	Calificacion int `json:"Calificacion"`
}

//Estructura Busqueda
type unico2 struct {
	Categoria string `json:"Categoria"`
	Tienda string `json:"Nombre"`
	Calificacion int `json:"Calificacion"`
}

//Linealizacion
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
	//Ordenar Valores en Listas
	for i:=0; i<len(vec);i++{
		vec[i] = vec[i].ordenar()
	}
}

//Reporte vector
func grafico1(){
	reportes(0,0,1)
}

func Eliminar(t unico2) []byte{
	aux := newLista()
	var v varios
	i := posicionv2(t)
	if vec!=nil{
		if vec[i].Vacio(){
			crearJson, _ := json.Marshal("Lista Vacia")
			return crearJson
		} else{
			if i < len(vec){
				a := vec[i].primero
				for a != nil{
					if t.Tienda != a.tienda.nombre{
						insertar(a.tienda,aux)
						ti := Tiendas{Tiendas:a.tienda.nombre,Descripcion:a.tienda.descripcion,Contacto:a.tienda.contacto,Calificacion:a.tienda.calif}
						v.Tiendas = append(v.Tiendas, ti)
					}
					a=a.sig
				}
			}
			vec[i] = *aux
			crearJson, _ := json.Marshal(v)
			return crearJson
		}
	}else{
		crearJson, _ := json.Marshal("No Hay Tiendas Cargadas")
		return crearJson
	}
}

//Id Lista
func posicionl(i int) []byte{
	if vec != nil{
		if vec[i-1].Vacio(){
			crearJson, _ := json.Marshal("Lista Vacia")
			return crearJson
		}else{
			var v varios
			a := vec[i-1].primero
			for a != nil{
				t := Tiendas{Tiendas:a.tienda.nombre,Descripcion:a.tienda.descripcion,Contacto:a.tienda.contacto,Calificacion:a.tienda.calif}
				v.Tiendas = append(v.Tiendas, t)
				a = a.sig
			}
			crearJson, _ := json.Marshal(v)
			return crearJson
		}
	}else{
		crearJson, _ := json.Marshal("No Hay Tiendas Cargadas")
		return crearJson
	}
}

//Tienda Especifica
func posiciont(t unico) Tiendas{
	var ti Tiendas
	i := posicionv(t)
	if i < len(vec){
		a := vec[i].primero
		for a != nil{
			if t.Tienda == a.tienda.nombre{
				ti.Tiendas = a.tienda.nombre
				ti.Descripcion = a.tienda.descripcion
				ti.Contacto = a.tienda.contacto
				ti.Calificacion = a.tienda.calif
				break
			}
			a=a.sig
		}
	}
	return ti
}

//Posicion Especifica
func posicionv(t unico) int{
	indice := string(t.Tienda[0])
	var i int
	var c int
	for a:=0;a<len(data.Datos);a++{
		if data.Datos[a].Indice == indice{
			i = a
			break
		}
	}

	for b:=0;b<len(data.Datos[0].Departamentos);b++{
		if t.Departamento == data.Datos[0].Departamentos[b].Departamentos{
			c = b
			break
		}
	}
	seg:=(i*len(data.Datos[0].Departamentos)) + c
	ter:= (seg*5)+ t.Calificacion-1
	return ter
}

//Posicion Especifica
func posicionv2(t unico2) int{
	indice := string(t.Tienda[0])
	var i int
	var c int
	for a:=0;a<len(data.Datos);a++{
		if data.Datos[a].Indice == indice{
			i = a
			break
		}
	}

	for b:=0;b<len(data.Datos[0].Departamentos);b++{
		if t.Categoria == data.Datos[0].Departamentos[b].Departamentos{
			c = b
			break
		}
	}
	seg:=(i*len(data.Datos[0].Departamentos)) + c
	ter:= (seg*5)+ t.Calificacion-1
	return ter
}
