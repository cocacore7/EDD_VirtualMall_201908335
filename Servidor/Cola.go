package Servidor

type Pedido struct {
	dia 				int
	NoPedido 			int
	Tienda 				string
	Categoria 			string
	Calificacion 		int
	codigos 			[]int
}

type nodoPedido struct {
	Pedido  *Pedido
	sig 	*nodoPedido
}

type Cola struct {
	primero  *nodoPedido
	ultimo   *nodoPedido
}

type ArregloNodoPedido struct {
	Der, Arr, Aba, Izq  interface{}
	Cola *Cola
}

func newPedido(dia int, noPedido int, Tienda string, Categoria string, Calificacion int, codigos []int) *Pedido{
	return &Pedido{dia,noPedido,Tienda,Categoria,Calificacion,codigos}
}

func newNodoPedido(pedido *Pedido) *nodoPedido{
	return &nodoPedido{pedido,nil }
}

func newCola() *Cola{
	return &Cola{nil,nil}
}

func newArregloPedidos() *ArregloNodoPedido{
	return &ArregloNodoPedido{nil,nil,nil,nil,newCola()}
}

func (l *Cola) Vacio() bool  {
	return l.primero == nil
}

//Insertar En Lista
func insertarPedido(t *Pedido, l *Cola) {
	var nuevo = newNodoPedido(t)
	if l.Vacio(){
		l.primero =nuevo
		l.ultimo  = nuevo
	} else {
		l.ultimo.sig = nuevo
		l.ultimo 	 = l.ultimo.sig
	}
}
