package Servidor


type Mes struct {
	mes      	string
}

type nodoMes struct {
	Mes  *Mes
	sig 	*nodoMes
	ant     *nodoMes
}

type listaMes struct {
	primero  *nodoMes
	ultimo   *nodoMes
}

func newMes(mes string) *Mes{
	return &Mes{mes}
}

func newNodoMes(m *Mes) *nodoMes {
	return &nodoMes{m,nil,nil}
}

func newListaMes() *listaMes{
	return &listaMes{nil,nil}
}

func (l *listaMes) VacioMes() bool  {
	return l.primero == nil
}

//Insertar En Lista
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

