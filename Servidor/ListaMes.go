package Servidor

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

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
func (l *listaMes) IngresarPedido(mes string,pedido *Pedido) *listaMes {
	aux := l.primero
	aux2 := newListaMes()
	for aux != nil {
		if mes == aux.Mes.mes{
			aux.Mes.matriz.Agregar(pedido)
			insertarMes(aux.Mes,aux2)
		}else{
			insertarMes(aux.Mes,aux2)
		}
		aux = aux.sig
	}
	return aux2
}

//Validar Stock Del Producto Solicitado En Pedido
func (l *listaMes) ValidarExistenciasMes(mes string,pedido *Pedido) *listaMes {
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

func (l *listaMes)GraficarMeses(){
	arch, _ := os.Create("GraficoMeses.dot")
	//Encabezado
	_, _ = arch.WriteString("digraph G{" + "\n")
	_, _ = arch.WriteString(`compound=true;` + "\n")
	_, _ = arch.WriteString(`edge[minlen=0.1, dir=fordware]` + "\n")
	aux := l.primero
	contador := 0
	for aux != nil{
		_, _ = arch.WriteString(`nodo`+strconv.Itoa(contador)+` [shape=record,color=red,label="` + aux.Mes.mes + `}"];` + "\n")
		contador++
		aux = aux.sig
	}
	for x:=1;x<contador;x++{
		_, _ = arch.WriteString(`nodo`+strconv.Itoa(x-1)+` -> nodo`+strconv.Itoa(x)+`;` + "\n")
		_, _ = arch.WriteString(`nodo`+strconv.Itoa(x)+` -> nodo`+strconv.Itoa(x-1)+`;` + "\n")
	}
	_, _ = arch.WriteString("}" + "\n")
	arch.Close()

	//Crear Archivo
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./GraficoMeses.dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("GraficoMeses.png", cmd, os.FileMode(mode))
}

func (l *listaMes) GraficarMatrizMes(mes string, bandera bool) bool {
	aux := l.primero
	for aux != nil {
		if mes == aux.Mes.mes{
			if aux.Mes.matriz.CabH != nil{
				aux.Mes.matriz.GraficarMatriz()
				bandera = true
			}
			break
		}
		aux = aux.sig
	}
	return bandera
}

func (l *listaMes) GraficarColaMatrizMes(mes string,dia int, categoria string, bandera bool) bool {
	aux := l.primero
	for aux != nil {
		if mes == aux.Mes.mes{
			bandera = aux.Mes.matriz.GraColaM(dia, categoria,false)
			break
		}
		aux = aux.sig
	}
	return bandera
}
