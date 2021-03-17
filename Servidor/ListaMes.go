package Servidor


type Mes struct {
	mes    string
	matriz *Matriz
}

type nodoMes struct {
	Mes  	*Mes
	sig 	*nodoMes
	ant     *nodoMes
}

type listaMes struct {
	primero  *nodoMes
	ultimo   *nodoMes
}

func newMes(mes string) *Mes{
	return &Mes{mes, newMatriz()}
}

func newNodoMes(m *Mes) *nodoMes {
	return &nodoMes{m,nil,nil}
}

func newListaMes() *listaMes{
	return &listaMes{nil,nil}
}

//VacioMes
func (l *listaMes) VacioMes() bool  {
	return l.primero == nil
}

//Insertar Mes En Lista
func insertarMes(t *Mes, l *listaMes){
	var nuevo = newNodoMes(t)
	if l.VacioMes(){
		l.primero =nuevo
		l.ultimo  = nuevo
	} else {
		l.ultimo.ant = l.ultimo
		l.ultimo.sig = nuevo
		l.ultimo 	 = l.ultimo.sig
	}
}

//Ingresar Pedido En Matriz De Mes
func (l *listaMes) IngresarPedido(mes string,pedido *nodoPedido) *listaMes {
	aux := l.primero
	aux2 := newListaMes()
	for aux != nil {
		if mes == aux.Mes.mes{
			aux.Mes.matriz.Agregar(pedido)
			NoPedido++
			insertarMes(aux.Mes,aux2)
			break
		}else{
			insertarMes(aux.Mes,aux2)
		}
		aux = aux.sig
	}
	return aux2
}

//Validar Stock Del Producto Solicitado En Pedido
func (l *listaMes) ValidarExistenciasMes(mes string,pedido *nodoPedido) *listaMes {
	aux := l.primero
	aux2 := newListaMes()
	for aux != nil {
		if mes == aux.Mes.mes{
			aux.Mes.matriz.Agregar(pedido)
			NoPedido++
			insertarMes(aux.Mes,aux2)
			break
		}else{
			insertarMes(aux.Mes,aux2)
		}
		aux = aux.sig
	}
	return aux2
}
