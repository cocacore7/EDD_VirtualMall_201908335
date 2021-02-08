package Servidor

import "fmt"

type tienda struct {
	nombre      string
	descripcion string
	contacto    string
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

func newTienda(nombre string, descripcion string, contacto string) *tienda{
	return &tienda{nombre,descripcion,contacto}
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

func imprimir(l *lista){
	aux := l.primero
	for aux != nil {
		fmt.Println("Nombre: " 		+ aux.tienda.nombre)
		fmt.Println("Descripcion: " + aux.tienda.descripcion)
		fmt.Println("Contacto: " 	+ aux.tienda.contacto)
		aux = aux.sig
	}
}

func Eliminar(l *lista){
	if !l.Vacio(){
		if l.primero  == l.ultimo{
			l.primero = nil
			l.ultimo  = nil
		}else{
			l.primero 	  = l.primero.sig
			l.primero.ant = nil
		}
	}
}
