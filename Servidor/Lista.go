package Servidor

import (
	"sort"
)

type tienda struct {
	nombre      string
	descripcion string
	contacto    string
	calif int
}

type nodo struct {
	tienda  *tienda
	sig 	*nodo
	ant     *nodo
}

type lista struct {
	primero  *nodo
	ultimo   *nodo
	contador int
}

func newTienda(nombre string, descripcion string, contacto string, calif int) *tienda{
	return &tienda{nombre,descripcion,contacto, calif}
}

func newNodo(t *tienda) *nodo {
	return &nodo{t,nil,nil}
}

func newLista() *lista{
	return &lista{nil,nil,0}
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
	l.contador++
}

func (l *lista) ordenar() lista {
	aux := l.primero
	valores := make([]int, 0, l.contador)
	//Obtenemos valores ascii en slice
	for aux != nil {
		valor := 0
		palabra := []rune(aux.tienda.nombre)
		for i:=0; i < len(aux.tienda.nombre); i++{
			valor = valor + int(palabra[i])
		}
		valores = append(valores, valor)
		aux = aux.sig
	}

	//Ordenamos slice ascii
	sort.Ints(valores)
	aux2 := newLista()

	//Creamos nueva lista con nodos ordenados
	for i:=0;i<len(valores);i++{
		aux = l.primero
		for aux != nil {
			valor := 0
			palabra := []rune(aux.tienda.nombre)
			for i:=0; i < len(aux.tienda.nombre); i++{
				valor = valor + int(palabra[i])
			}
			if valores[i] == valor{
				insertar(newTienda(aux.tienda.nombre, aux.tienda.descripcion,aux.tienda.contacto, aux.tienda.calif),aux2)
				break
			}
			aux = aux.sig
		}
	}
	return *aux2
}

/*func Eliminar(l *lista){
	if !l.Vacio(){
		if l.primero  == l.ultimo{
			l.primero = nil
			l.ultimo  = nil
		}else{
			l.primero 	  = l.primero.sig
			l.primero.ant = nil
		}
		l.contador--
	}
}*/
