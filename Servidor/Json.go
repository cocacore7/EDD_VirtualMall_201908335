package Servidor

import (
	"encoding/json"
	"os"
)

var años *ArbolAño
var vec []lista
var Indi []string
var Depa []string

//---------------------------------------------------------------------------------------------
//FASE1

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
	Logo string `json:"Logo"`
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

//FASE1
//---------------------------------------------------------------------------------------------

//---------------------------------------------------------------------------------------------
//FASE2

//Estructura Inventarios json
type Inventarios struct {
	Inventarios []TiendasInventario `json:"Invetarios"`
}

type TiendasInventario struct {
	Tienda string `json:"Tienda"`
	Departamento string `json:"Departamento"`
	Calificacion int `json:"Calificacion"`
	Productos []Productos `json:"Productos"`
}

type Productos struct {
	Nombre string `json:"Nombre"`
	Codigo int `json:"Codigo"`
	Descripcion string `json:"Descripcion"`
	Precio int `json:"Precio"`
	Cantidad int `json:"Cantidad"`
	Imagen string `json:"Imagen"`
}

//FASE2
//---------------------------------------------------------------------------------------------

//---------------------------------------------------------------------------------------------
//FASE1

//Linealizacion
func Crear(data Datos){
	vec = make([]lista, 0)

	for j:=0; j<len(data.Datos[0].Departamentos);j++{
		Depa = append(Depa, data.Datos[0].Departamentos[j].Departamentos)
	}

	//Creando Vector E Insertando Listas con Tiendas "ROW-MAJOR"
	for i := 0; i < len(data.Datos); i++ {
		Indi = append(Indi, data.Datos[i].Indice)
		for j:=0; j<len(data.Datos[i].Departamentos);j++{
			l1 := newLista()
			l2 := newLista()
			l3 := newLista()
			l4 := newLista()
			l5 := newLista()
			for x:=0;x<len(data.Datos[i].Departamentos[j].Tiendas);x++{
				t := newTienda(data.Datos[i].Departamentos[j].Tiendas[x].Tiendas, data.Datos[i].Departamentos[j].Tiendas[x].Descripcion, data.Datos[i].Departamentos[j].Tiendas[x].Contacto, data.Datos[i].Departamentos[j].Tiendas[x].Calificacion,data.Datos[i].Departamentos[j].Tiendas[x].Logo)
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

//Obtener Tienda En Posicion Especifica
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
				ti.Logo = a.tienda.logo
				break
			}
			a=a.sig
		}
	}
	return ti
}

//Obtener Posicion Especifica En posiciont
func posicionv(t unico) int{
	indice := string(t.Tienda[0])
	var i int
	var c int
	for a:=0;a<len(Indi);a++{
		if Indi[a] == indice{
			i = a
			break
		}
	}

	for b:=0;b<len(Depa);b++{
		if t.Departamento == Depa[b]{
			c = b
			break
		}
	}
	seg:=(i*len(Depa)) + c
	ter:= (seg*5)+ t.Calificacion-1
	return ter
}

//Obtener Tiendas Por Id De Lista
func posicionl(i int) []byte{
	if vec != nil{
		if vec[i-1].Vacio(){
			crearJson, _ := json.Marshal("Lista Vacia")
			return crearJson
		}else{
			var v varios
			a := vec[i-1].primero
			for a != nil{
				t := Tiendas{Tiendas:a.tienda.nombre,Descripcion:a.tienda.descripcion,Contacto:a.tienda.contacto,Calificacion:a.tienda.calif,Logo: a.tienda.logo}
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

//Eliminar tieda de linealizado
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
						ti := Tiendas{Tiendas:a.tienda.nombre,Descripcion:a.tienda.descripcion,Contacto:a.tienda.contacto,Calificacion:a.tienda.calif,Logo: a.tienda.logo}
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

//Obtener Posicion Especifica En Eliminar
func posicionv2(t unico2) int{
	indice := string(t.Tienda[0])
	var i int
	var c int
	for a:=0;a<len(Indi);a++{
		if Indi[a] == indice{
			i = a
			break
		}
	}

	for b:=0;b<len(Depa);b++{
		if t.Categoria == Depa[b]{
			c = b
			break
		}
	}
	seg:=(i*len(Depa)) + c
	ter:= (seg*5)+ t.Calificacion-1
	return ter
}

//retorno de linealizacion final en json
func retorno(in int, fi int) []byte{
	var reg Datos
	reg.Datos = make([]Indice,len(Indi))
	for i:=0;i<len(Indi);i++{
		reg.Datos[i].Indice = Indi[i]
		reg.Datos[i].Departamentos = make([]Departamentos,len(Depa))
		for j:=0;j<len(Depa);j++{
			reg.Datos[i].Departamentos[j].Departamentos = Depa[j]
			if in+5 <= len(vec){
				fi = fi + 5
			}else if in+4 <= len(vec){
				fi = fi + 4
			}else if in+3 <= len(vec){
				fi = fi + 3
			}else if in+2 <= len(vec){
				fi = fi + 2
			}else if in+1 <= len(vec){
				fi = fi + 1
			}
			reg.Datos[i].Departamentos[j].Tiendas = obtenerT(in, fi)
			in = fi
		}
	}
	arch, _ := os.Create("Salida.json")
	crearJson, _ := json.MarshalIndent(reg,"","    ")
	_, _ = arch.WriteString(string(crearJson))
	arch.Close()
	return crearJson
}

//Llamada recursiva de tiendas funcion retorno
func obtenerT(i int, f int) []Tiendas{
	var t []Tiendas
	var aux Tiendas
	for i<f{
		a := vec[i].primero
		for a != nil{
			aux.Tiendas = a.tienda.nombre
			aux.Descripcion = a.tienda.descripcion
			aux.Contacto = a.tienda.contacto
			aux.Calificacion = a.tienda.calif
			aux.Logo = a.tienda.logo
			t = append(t, aux)
			a = a.sig
		}
		i++
	}
	return t
}

//FASE1
//---------------------------------------------------------------------------------------------

//---------------------------------------------------------------------------------------------
//FASE2

//Insertar Productos En Tiendas Linealizadas
func Inventario(t Inventarios) []byte{
	if vec!=nil{
		for y:=0;y<len(t.Inventarios);y++{
			i := posicionv3(t.Inventarios[y])
			if vec[i].Vacio(){
				crearJson, _ := json.Marshal("No existen Tiendas Cargadas")
				return crearJson
			} else{
				if i < len(vec){
					a := vec[i].primero
					for a != nil{
						if t.Inventarios[y].Tienda == a.tienda.nombre{
							for x:=0;x<len(t.Inventarios[y].Productos);x++{
								p := NewProducto(t.Inventarios[y].Productos[x].Nombre,t.Inventarios[y].Productos[x].Codigo,t.Inventarios[y].Productos[x].Descripcion,t.Inventarios[y].Productos[x].Precio,t.Inventarios[y].Productos[x].Cantidad,t.Inventarios[y].Productos[x].Imagen)
								a.tienda.productos.InsertarAVLProducto(*p,x+1)
							}
						}
						a=a.sig
					}
				}
			}
		}
		crearJson, _ := json.Marshal(t)
		return crearJson
	}else{
		crearJson, _ := json.Marshal("No Hay Tiendas Cargadas")
		return crearJson
	}
}

//Obtener Posicion Especifica En Inventario
func posicionv3(t TiendasInventario) int{
	indice := string(t.Tienda[0])
	var i int
	var c int
	for a:=0;a<len(Indi);a++{
		if Indi[a] == indice{
			i = a
			break
		}
	}

	for b:=0;b<len(Depa);b++{
		if t.Departamento == Depa[b]{
			c = b
			break
		}
	}
	seg:=(i*len(Depa)) + c
	ter:= (seg*5)+ t.Calificacion-1
	return ter
}

//FASE2
//---------------------------------------------------------------------------------------------
