package Servidor

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var añosArbol *ArbolAño
var vec []lista
var usuarios *ArbolB
var grafo *ListaAdyacencia
var Indi []string
var Depa []string
var NoPedido int
var NodoG int
var idenAño = 0
var regresoComP []NodoHash

//---------------------------------------------------------------------------------------------
//FASE1

//Estructura tiendas json
type Datos struct{
	Datos []Indice `json:"Datos"`
}

type Indice struct{
	Indice 		  string 		  `json:"Indice"`
	Departamentos []Departamentos `json:"Departamentos"`
}

type Departamentos struct {
	Departamentos string 	`json:"Nombre"`
	Tiendas 	  []Tiendas `json:"Tiendas"`
}

type Tiendas struct {
	Tiendas 		string  `json:"Nombre"`
	Descripcion 	string  `json:"Descripcion"`
	Contacto 		string  `json:"Contacto"`
	Calificacion 	int 	`json:"Calificacion"`
	Logo 			string  `json:"Logo"`
}

//Estructura complementaria
type varios struct {
	Tiendas []Tiendas `json:"Tiendas"`
}

//Estructura Busqueda
type unico struct {
	Departamento string `json:"Departamento"`
	Tienda 		 string `json:"Nombre"`
	Calificacion int 	`json:"Calificacion"`
}

//Estructura Busqueda
type unico2 struct {
	Categoria 		string  `json:"Categoria"`
	Tienda 			string  `json:"Nombre"`
	Calificacion 	int 	`json:"Calificacion"`
}

//FASE1
//---------------------------------------------------------------------------------------------

//---------------------------------------------------------------------------------------------
//FASE2

//Estructura Inventarios json
type Inventarios struct {
	Inventarios []TiendasInventario `json:"Inventarios"`
}

type TiendasInventario struct {
	Tienda 			string  	`json:"Tienda"`
	Departamento 	string  	`json:"Departamento"`
	Calificacion 	int 		`json:"Calificacion"`
	Productos 		[]Productos `json:"Productos"`
}

type Productos struct {
	Nombre 			string  `json:"Nombre"`
	Codigo 			int 	`json:"Codigo"`
	Descripcion 	string  `json:"Descripcion"`
	Precio 			int 	`json:"Precio"`
	Cantidad 		int 	`json:"Cantidad"`
	Imagen 			string  `json:"Imagen"`
	Almacenamiento 	string 	`json:"Almacenamiento"`
}

type Pedidos struct {
	Pedidos []tiendaPedido `json:"Pedidos"`
}

type tiendaPedido struct {
	Fecha 			string 				`json:"Fecha"`
	Tienda 			string 				`json:"Tienda"`
	Departamento 	string 				`json:"Departamento"`
	Calificacion 	int    				`json:"Calificacion"`
	Cliente 		int 				`json:"Cliente"`
	Productos 		[]productosPedido   `json:"Productos"`
}

type productosPedido struct {
	Codigo int `json:"Codigo"`
}

//FASE2
//---------------------------------------------------------------------------------------------

//---------------------------------------------------------------------------------------------
//FASE3

type Usuarios struct {
	Usuarios []Usuario `json:"Usuarios"`
}

type Usuario struct {
	Dpi 			int 				`json:"Dpi"`
	Nombre 			string 				`json:"Nombre"`
	Correo 			string 				`json:"Correo"`
	Password 		string 				`json:"Password"`
	Cuenta 			string 				`json:"Cuenta"`
}

type Grafo struct {
	Nodos 					[]Nodo 		`json:"Nodos"`
	PosicionInicialRobot 	string 		`json:"PosicionInicialRobot"`
	Entrega 				string 		`json:"Entrega"`
}

type Nodo struct {
	Nombre 			string 				`json:"Nombre"`
	Enlaces 		[]Enlace 			`json:"Enlaces"`
}

type Enlace struct {
	Nombre 			string 				`json:"Nombre"`
	Distancia 		int 				`json:"Distancia"`
}

//FASE3
//---------------------------------------------------------------------------------------------

//---------------------------------------------------------------------------------------------
//FASE4

type Comentarios struct {
	Comentarios []com	`json:"Comentarios"`
}

type com struct {
	Tipo 			string 				`json:"Tipo"`
	Tienda 			string 				`json:"Tienda"`
	Departamento 	string 				`json:"Departamento"`
	Calificacion 	int    				`json:"Calificacion"`
	Codigo 			int 				`json:"Codigo"`
	Dpi 			int 				`json:"Dpi"`
	Fecha 			string 				`json:"Fecha"`
	Hora 			string				`json:"Hora"`
	Comentario 		string 				`json:"Comentario"`
}

type coment struct {
	Tipo 			string 				`json:"Tipo"`
	Tienda 			string 				`json:"Tienda"`
	Departamento 	string 				`json:"Departamento"`
	Calificacion 	int    				`json:"Calificacion"`
	Codigo 			int 				`json:"Codigo"`
	Dpi 			int 				`json:"Dpi"`
	Comentario 		string 				`json:"Comentario"`
}

//FASE4
//---------------------------------------------------------------------------------------------

//---------------------------------------------------------------------------------------------
//FASE1

//Linealizacion E Inicializacion AVLAños
func Crear(data Datos){
	vec = make([]lista, 0)
	Indi = make([]string, 0)
	Depa = make([]string, 0)
	añosArbol = NewArbolAño()
	NoPedido = 1
	NodoG = 0

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
	indice := strings.ToUpper(string(t.Tienda[0]))
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
	indice := strings.ToUpper(string(t.Tienda[0]))
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
	var regreso Inventarios
	regreso.Inventarios = make([]TiendasInventario,0)
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
								p := NewProducto(t.Inventarios[y].Productos[x].Nombre,t.Inventarios[y].Productos[x].Codigo,t.Inventarios[y].Productos[x].Descripcion,t.Inventarios[y].Productos[x].Precio,t.Inventarios[y].Productos[x].Cantidad,t.Inventarios[y].Productos[x].Imagen,t.Inventarios[y].Productos[x].Almacenamiento)
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

//Regresar Productos De Vector Linealizado

func RegresoProductos() []byte{
	var regreso Inventarios
	regreso.Inventarios = make([]TiendasInventario,0)
	if vec!=nil{
		contador := 0
		for x:=0;x<len(Indi);x++{
			for y:=0;y<len(Depa);y++{
				var tinv TiendasInventario
				tinv.Departamento = Depa[y]
				for z:=0;z<5;z++{
						aux := vec[contador].primero
						for aux!=nil{
							tinv.Tienda = aux.tienda.nombre
							tinv.Calificacion = aux.tienda.calif
							tinv.Productos = ProuctosModificado(aux.tienda.productos.raiz,tinv.Productos)
							aux = aux.sig
							regreso.Inventarios = append(regreso.Inventarios, tinv)
							tinv.Productos = nil
						}
					contador++
				}
			}
		}
		crearJson, _ := json.Marshal(regreso)
		return crearJson
	}else{
		crearJson, _ := json.Marshal("No Hay Tiendas Cargadas")
		return crearJson
	}
}

//Obtener Posicion Especifica En Inventario
func posicionv3(t TiendasInventario) int{
	indice := strings.ToUpper(string(t.Tienda[0]))
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

//Insertar Pedidos En Arbol Años Desde Administrador

func PedidosJson(t Pedidos) []byte{
	if vec!=nil{
		for y:=0;y<len(t.Pedidos);y++{
			i := posicionv4(t.Pedidos[y])
			if vec[i].Vacio(){
				crearJson, _ := json.Marshal("No existen Tiendas Cargadas")
				return crearJson
			} else{
				//Fecha
				fecha := strings.Split(t.Pedidos[y].Fecha,"-")
				año, _ := strconv.Atoi(fecha[2])
				mes := obtenerMes(fecha[1])
				dia, _ := strconv.Atoi(fecha[0])

				//Ingresar Año Y Mes
				añosArbol.InsertarAVLAño(*NewAño(año),idenAño)
				idenAño++
				añosArbol.raiz = insertarMesArbol(añosArbol.raiz,año,mes)

				//Obtener Productos De Tienda solicitada A Consultar
				codigosAux:=make([]int,0)
				aux:=vec[i].primero
				var produtosaux *ArbolProducto
				for aux != nil{
					if aux.tienda.nombre == t.Pedidos[y].Tienda{
						produtosaux = aux.tienda.productos
						break
					}
					aux = aux.sig
				}
				//Obtener Codigos Validos De Pedidos EN Tienda Solicitada
				if produtosaux!=nil{
					if produtosaux.raiz != nil{
						for x:=0;x<len(t.Pedidos[y].Productos);x++{
							if buscarCodigoPedido(produtosaux.raiz,t.Pedidos[y].Productos[x].Codigo,false){
								bandera := true
								for z:=0;z<len(codigosAux);z++{
									if codigosAux[z]== t.Pedidos[y].Productos[x].Codigo{
										fmt.Println("El Codigo: "+ strconv.Itoa(t.Pedidos[y].Productos[x].Codigo)+ " Viene repetido Dentro Del Pedido")
										bandera = false
										break
									}
								}
								if bandera{
									codigosAux = append(codigosAux, t.Pedidos[y].Productos[x].Codigo)
								}
							}
						}

						//Validar Stock Del Pedido Y Restar Stock Solicitado
						bandera := true
						contador:=0
						for x:=0;x<len(t.Pedidos[y].Productos);x++{
							if ValidarExistencias(produtosaux.raiz,t.Pedidos[y].Productos[x].Codigo,false){
								vec[i] = vec[i].RestarStockLista(t.Pedidos[y].Tienda,t.Pedidos[y].Productos[x].Codigo)
								contador++
							}else{
								fmt.Println("Pedido Denegado, Producto con condigo: "+strconv.Itoa(t.Pedidos[y].Productos[x].Codigo)+" No Cuenta Con Existencias Solicitadas")
								bandera = false
								break
							}
						}
						if bandera{
							//Insertar Pedido En Matriz
							ped:=newPedido(dia,NoPedido,t.Pedidos[y].Tienda,t.Pedidos[y].Departamento,t.Pedidos[y].Calificacion,codigosAux)
							NoPedido++
							añosArbol.raiz = insertarPedidoArbol(añosArbol.raiz,año,mes,ped)
						}else {
							for x:=0;x<contador;x++{
								vec[i] = vec[i].SumarStockLista(t.Pedidos[y].Tienda,t.Pedidos[y].Productos[x].Codigo)
							}
						}
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

//Obtener Posicion Especifica En Pedidos
func posicionv4(t tiendaPedido) int{
	indice := strings.ToUpper(string(t.Tienda[0]))
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

//Obtener Mes
func obtenerMes(m string) string{
	switch m {
	case "01":
		return "Enero"
	case "02":
		return "Febrero"
	case "03":
		return "Marzo"
	case "04":
		return "Abril"
	case "05":
		return "Mayo"
	case "06":
		return "Junio"
	case "07":
		return "Julio"
	case "08":
		return "Agosto"
	case "09":
		return "Septiembre"
	case "10":
		return "Octubre"
	case "11":
		return "Noviembre"
	case "12":
		return "Diciembre"
	}
	return "Invalido"
}

//Insertar Pedidos En Arbol Años

func ValidarPedidosJsonCarrito(t Pedidos) bool{
	bandera2:=true
	var Posiciones = make([]int,0)
	var Tiendasaux = make([]string,0)
	var Codigos = make([]int,0)
	if vec!=nil{
		for y:=0;y<len(t.Pedidos);y++{
			i := posicionv4(t.Pedidos[y])
			if vec[i].Vacio(){
				return false
			} else{
				//Obtener Productos De Tienda solicitada A Consultar
				codigosAux:=make([]int,0)
				aux:=vec[i].primero
				var produtosaux *ArbolProducto
				for aux != nil{
					if aux.tienda.nombre == t.Pedidos[y].Tienda{
						produtosaux = aux.tienda.productos
						break
					}
					aux = aux.sig
				}
				//Obtener Codigos Validos De Pedidos EN Tienda Solicitada
				if produtosaux!=nil{
					if produtosaux.raiz != nil{
						for x:=0;x<len(t.Pedidos[y].Productos);x++{
							if buscarCodigoPedido(produtosaux.raiz,t.Pedidos[y].Productos[x].Codigo,false){
								bandera := true
								for z:=0;z<len(codigosAux);z++{
									if codigosAux[z]== t.Pedidos[y].Productos[x].Codigo{
										bandera = false
										break
									}
								}
								if bandera{
									codigosAux = append(codigosAux, t.Pedidos[y].Productos[x].Codigo)
								}
							}
						}

						//Validar Stock Del Pedido Y Restar Stock Solicitado
						bandera := true
						contador:=0
						for x:=0;x<len(t.Pedidos[y].Productos);x++{
							if ValidarExistencias(produtosaux.raiz,t.Pedidos[y].Productos[x].Codigo,false){
								vec[i] = vec[i].RestarStockLista(t.Pedidos[y].Tienda,t.Pedidos[y].Productos[x].Codigo)
								Posiciones = append(Posiciones, i)
								Tiendasaux = append(Tiendasaux, t.Pedidos[y].Tienda)
								Codigos = append(Codigos, t.Pedidos[y].Productos[x].Codigo)
								contador++
							}else{
								Cod = t.Pedidos[y].Productos[x].Codigo
								bandera = false
								bandera2 = false
								break
							}
						}
						if !bandera{
							break
						}
					}
				}
			}
			if !bandera2{break}
		}
		if !bandera2{
			for x:=0;x<len(Tiendasaux);x++{
				vec[Posiciones[x]] = vec[Posiciones[x]].SumarStockLista(Tiendasaux[x],Codigos[x])
			}
		}
		return bandera2
	}else{
		return false
	}
}

//Insertar Pedidos En Arbol Años Desde Carrito

func PedidosJsonCarrito(t Pedidos) []byte{
	if vec!=nil{
		for y:=0;y<len(t.Pedidos);y++{
			i := posicionv4(t.Pedidos[y])
			if vec[i].Vacio(){
				crearJson, _ := json.Marshal("No existen Tiendas Cargadas")
				return crearJson
			} else{
				//Fecha
				fecha := strings.Split(t.Pedidos[y].Fecha,"-")
				año, _ := strconv.Atoi(fecha[2])
				mes := obtenerMesCarrito(fecha[1])
				dia, _ := strconv.Atoi(fecha[0])

				//Ingresar Año Y Mes
				añosArbol.InsertarAVLAño(*NewAño(año),idenAño)
				idenAño++
				añosArbol.raiz = insertarMesArbol(añosArbol.raiz,año,mes)

				//Obtener Productos De Tienda solicitada A Consultar
				codigosAux:=make([]int,0)
				aux:=vec[i].primero
				var produtosaux *ArbolProducto
				for aux != nil{
					if aux.tienda.nombre == t.Pedidos[y].Tienda{
						produtosaux = aux.tienda.productos
						break
					}
					aux = aux.sig
				}
				//Obtener Codigos Validos De Pedidos EN Tienda Solicitada
				if produtosaux!=nil{
					if produtosaux.raiz != nil{
						for x:=0;x<len(t.Pedidos[y].Productos);x++{
							if buscarCodigoPedido(produtosaux.raiz,t.Pedidos[y].Productos[x].Codigo,false){
								bandera := true
								for z:=0;z<len(codigosAux);z++{
									if codigosAux[z]== t.Pedidos[y].Productos[x].Codigo{
										fmt.Println("El Codigo: "+ strconv.Itoa(t.Pedidos[y].Productos[x].Codigo)+ " Viene repetido Dentro Del Pedido")
										bandera = false
										break
									}
								}
								if bandera{
									codigosAux = append(codigosAux, t.Pedidos[y].Productos[x].Codigo)
								}
							}
						}
						//Insertar Pedido En Matriz
						ped:=newPedido(dia,NoPedido,t.Pedidos[y].Tienda,t.Pedidos[y].Departamento,t.Pedidos[y].Calificacion,codigosAux)
						añosArbol.raiz = insertarPedidoArbol(añosArbol.raiz,año,mes,ped)
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

//Obtener Mes
func obtenerMesCarrito(m string) string{
	switch m {
	case "1":
		return "Enero"
	case "2":
		return "Febrero"
	case "3":
		return "Marzo"
	case "4":
		return "Abril"
	case "5":
		return "Mayo"
	case "6":
		return "Junio"
	case "7":
		return "Julio"
	case "8":
		return "Agosto"
	case "9":
		return "Septiembre"
	case "10":
		return "Octubre"
	case "11":
		return "Noviembre"
	case "12":
		return "Diciembre"
	}
	return "Invalido"
}

//Obtener Tienda En Posicion Especifica
func graficarArbolP(t unico, bandera bool) bool{
	i:=posicionv(t)
	if i < len(vec){
		a := vec[i].primero
		for a != nil{
			if t.Tienda == a.tienda.nombre{
				bandera = graficar(a.tienda.productos,false)
				break
			}
			a=a.sig
		}
	}
	return bandera
}

//FASE2
//---------------------------------------------------------------------------------------------

//---------------------------------------------------------------------------------------------
//FASE3

//Crear y Agregar Usuarios Al Arbol B

func AgregarUsuarios(t Usuarios) []byte{
	if usuarios == nil{
		usuarios = NewArbolB(5)

		hash := sha256.Sum256([]byte("1234"))
		a := NewKeyAB(Usuario{Dpi: 1234567890101,Nombre: "EDD2021",Correo: "auxiliar@edd.com",Password: hex.EncodeToString(hash[:]),Cuenta: "Admin"})
		usuarios.InsertarAB(a,true)
	}
	for i:=0;i<len(t.Usuarios);i++{
		contraseña := sha256.Sum256([]byte(t.Usuarios[i].Password))
		a := NewKeyAB(Usuario{Dpi: t.Usuarios[i].Dpi,Nombre: t.Usuarios[i].Nombre,Correo: t.Usuarios[i].Correo,Password: hex.EncodeToString(contraseña[:]),Cuenta: t.Usuarios[i].Cuenta})
		usuarios.InsertarAB(a,true)
	}
	crearJson, _ := json.Marshal(t)
	return crearJson
}

//Agregar Usuario al Arbol B

func CrearUsuario(t Usuario) bool{
	if usuarios == nil{
		usuarios = NewArbolB(5)

		hash := sha256.Sum256([]byte("1234"))
		a := NewKeyAB(Usuario{Dpi: 1234567890101,Nombre: "EDD2021",Correo: "auxiliar@edd.com",Password: hex.EncodeToString(hash[:]),Cuenta: "Admin"})
		usuarios.InsertarAB(a,true)
	}
	a := NewKeyAB(t)
	b:=usuarios.InsertarAB(a,true)
	if b {
		return true
	}else{
		return false
	}
}

//Buscar Usuario al Arbol B

func BuscarUsuario(t Usuario) Usuario{
	if usuarios == nil{
		usuarios = NewArbolB(5)

		hash := sha256.Sum256([]byte("1234"))
		a := NewKeyAB(Usuario{Dpi: 1234567890101,Nombre: "EDD2021",Correo: "auxiliar@edd.com",Password: hex.EncodeToString(hash[:]),Cuenta: "Admin"})
		usuarios.InsertarAB(a,true)
	}
	a:= NewKeyAB(t)
	b:=usuarios.BuscarAB(a)
	return b
}

//Eliminar Usuario Del Arbol B

func EliminarUsuario(t Usuario) []byte{
	//Eliminar Usuario
	crearJson, _ := json.Marshal(t)
	return crearJson
}

//Grafo dirigido
func definirGrafo(t Grafo) []byte{
	grafo = NewListaAdyacencia(t.PosicionInicialRobot,t.Entrega)
	for i:=0;i<len(t.Nodos);i++ {
		grafo.Insertar(t.Nodos[i].Nombre,0)
		for j:=0;j<len(t.Nodos[i].Enlaces);j++{
			grafo.Insertar(t.Nodos[i].Enlaces[j].Nombre,t.Nodos[i].Enlaces[j].Distancia)
			grafo.enlazar(t.Nodos[i].Nombre,t.Nodos[i].Enlaces[j].Nombre)
		}
	}
	crearJson, _ := json.Marshal(t)
	return crearJson
}

func dibujarGrafo() bool{
	if grafo !=nil{
		grafo.dibujar()
		return true
	}
	return false
}

//FASE3
//---------------------------------------------------------------------------------------------

//---------------------------------------------------------------------------------------------
//FASE4

//Ingresar Comentario
func Comenta(t coment) []byte{
	if vec!=nil{
		i := posicionv5(t)
		if vec[i].Vacio(){
			crearJson, _ := json.Marshal("No existen Tiendas Cargadas")
			return crearJson
		} else{
			if i < len(vec){
				a := vec[i].primero
				for a != nil{
					if t.Tienda == a.tienda.nombre{
						if t.Tipo == "Tienda"{
							a.tienda.comentarioT.insertar(t.Dpi,t.Comentario)
						}else{
							agregarComentario(a.tienda.productos.raiz,t.Codigo,t)
						}

					}
					a=a.sig
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

func posicionv5(t coment) int{
	indice := strings.ToUpper(string(t.Tienda[0]))
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

//Ingresar Comentario
func MostrarComenta(t coment) []byte{
	var regreso Comentarios
	regreso.Comentarios = make([]com,0)
	if vec!=nil{
		i := posicionv5(t)
		if vec[i].Vacio(){
			crearJson, _ := json.Marshal("No existen Tiendas Cargadas")
			return crearJson
		} else{
			if i < len(vec){
				a := vec[i].primero
				for a != nil{
					if t.Tienda == a.tienda.nombre{
						if t.Tipo == "Tienda"{
							reg :=a.tienda.comentarioT.Buscar()
							for b:=0;b<len(reg);b++ {
								var c com
								c.Tienda = "Tienda"
								c.Tienda = t.Tienda
								c.Departamento = t.Departamento
								c.Calificacion = t.Calificacion
								c.Codigo = t.Codigo
								c.Comentario = reg[b].valor
								c.Dpi = reg[b].hash
								c.Fecha = strconv.Itoa(reg[b].Año) + "/" + strconv.Itoa(reg[b].Mes) + "/" + strconv.Itoa(reg[b].Dia)
								c.Hora = strconv.Itoa(reg[b].Hora) + ":" + strconv.Itoa(reg[b].Minuto) + ":" + strconv.Itoa(reg[b].Seg)
								regreso.Comentarios = append(regreso.Comentarios, c)
							}
						}else{
							BuscarComentario(a.tienda.productos.raiz,t.Codigo)
							for b:=0;b<len(regresoComP);b++ {
								var c com
								c.Tienda = "Producto"
								c.Tienda = t.Tienda
								c.Departamento = t.Departamento
								c.Calificacion = t.Calificacion
								c.Codigo = t.Codigo
								c.Comentario = regresoComP[b].valor
								c.Dpi = regresoComP[b].hash
								c.Fecha = strconv.Itoa(regresoComP[b].Año) + "/" + strconv.Itoa(regresoComP[b].Mes) + "/" + strconv.Itoa(regresoComP[b].Dia)
								c.Hora = strconv.Itoa(regresoComP[b].Hora) + ":" + strconv.Itoa(regresoComP[b].Minuto) + ":" + strconv.Itoa(regresoComP[b].Seg)
								regreso.Comentarios = append(regreso.Comentarios, c)
							}
						}

					}
					a=a.sig
				}
			}
		}
		crearJson, _ := json.Marshal(regreso)
		return crearJson
	}else{
		crearJson, _ := json.Marshal("No Hay Tiendas Cargadas")
		return crearJson
	}
}

//FASE4
//---------------------------------------------------------------------------------------------
