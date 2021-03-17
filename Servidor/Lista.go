package Servidor

import "fmt"

type tienda struct {
	nombre      string
	descripcion string
	contacto    string
	calif 		int
	logo 		string
	productos 	*ArbolProducto
}

type nodo struct {
	tienda  *tienda
	sig 	*nodo
	ant     *nodo
}

type lista struct {
	primero  *nodo
	ultimo   *nodo
}

func newTienda(nombre string, descripcion string, contacto string, calif int,logo string) *tienda{
	return &tienda{nombre,descripcion,contacto, calif,logo, NewArbolProducto()}
}

func newNodo(t *tienda) *nodo {
	return &nodo{t,nil,nil}
}

func newLista() *lista{
	return &lista{nil,nil}
}

func (l *lista) Vacio() bool  {
	return l.primero == nil
}

//Insertar En Lista
func insertar(t *tienda, l *lista){
	var nuevo = newNodo(t)
	if l.Vacio(){
		l.primero =nuevo
		l.ultimo  = nuevo
	} else {
		l.ultimo.ant = l.ultimo
		l.ultimo.sig = nuevo
		l.ultimo 	 = l.ultimo.sig
	}
}

//Ordenar Lista
func (l *lista) ordenar() lista {
	aux := l.primero
	valores := make([]int, 0)
	//Obtenemos valores ascii en slice
	for aux != nil {
		valores = append(valores, ascii(aux.tienda.nombre))
		aux = aux.sig
	}

	//Ordenamos slice ascii
	burbuja(valores)
	aux2 := newLista()

	//Creamos nueva lista con nodos ordenados
	for i:=0;i<len(valores);i++{
		aux = l.primero
		for aux != nil {
			valor := ascii(aux.tienda.nombre)
			if valores[i] == valor{
				insertar(newTienda(aux.tienda.nombre, aux.tienda.descripcion,aux.tienda.contacto, aux.tienda.calif,aux.tienda.logo),aux2)
				break
			}
			aux = aux.sig
		}
	}
	return *aux2
}

func burbuja(arreglo []int) {
	var i, j, aux int
	for i = 0; i < len(arreglo)-1; i++ {
		for j = 0; j < len(arreglo)-i-1; j++ {
			if arreglo[j+1] < arreglo[j] {
				aux = arreglo[j+1]
				arreglo[j+1] = arreglo[j]
				arreglo[j] = aux
			}
		}
	}
	fmt.Println(arreglo)
}

func ascii(a string) int{
	valor := 0
	palabra := []rune(a)
	for i:=0; i < len(palabra); i++{
		valor = valor + int(palabra[i])
	}
	return valor
}

func (l *lista) RestarStockLista(t string, codigo int) lista{
	aux:=l.primero
	aux2:=newLista()
	for aux != nil{
		if aux.tienda.nombre == t{
			aux.tienda.productos.raiz = RestarStock(aux.tienda.productos.raiz,codigo)
			insertar(aux.tienda,aux2)
		}else{insertar(aux.tienda,aux2)}
		aux = aux.sig
	}
	return *aux2
}
