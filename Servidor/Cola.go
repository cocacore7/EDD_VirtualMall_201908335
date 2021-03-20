package Servidor

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

var InicialesCola []string
var Np int

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
	nodog 				string
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

func newArregloPedidos(nodog string) *ArregloNodoPedido{
	return &ArregloNodoPedido{nil,nil,nil,nil,newCola(),nodog}
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

func (l *Cola) GraficarPedidos() bool{
	Np = 0
	InicialesCola = make([]string,0)
	arch, _ := os.Create("GraficoColaMatriz.dot")
	//Encabezado
	_, _ = arch.WriteString("digraph G{" + "\n")
	_, _ = arch.WriteString(`compound=true;` + "\n")

	//Primer Subgrafo
	_, _ = arch.WriteString(`subgraph cluster0{style=invis;` + "\n")
	_, _ = arch.WriteString(`edge[minlen=0.1, dir=fordware]` + "\n")
	aux:=l.primero
	contador:=0
	for aux != nil{
		_, _ = arch.WriteString("n"+strconv.Itoa(contador)+`[shape=box,color=green,label="`+strconv.Itoa(aux.Pedido.NoPedido)+"|"+aux.Pedido.Tienda+`"];`+"\n")
		contador++
		aux = aux.sig
	}
	for x:=0;x<contador;x++{
		if x != (contador-1){
			_, _ = arch.WriteString("n"+strconv.Itoa(x)+"->n"+strconv.Itoa(x+1)+"\n")
		}
	}
	_, _ = arch.WriteString("}" + "\n")

	//Subgrafos de Codigos
	aux = l.primero
	subg:=1
	for aux != nil{
		_, _ = arch.WriteString(subgrafosCola(aux.Pedido.codigos,subg,Np))
		subg++
		aux = aux.sig
	}

	//ConectarSubgrafos
	for x:=0;x<contador;x++{
		_, _ = arch.WriteString("n"+strconv.Itoa(x)+"->"+InicialesCola[x]+";\n")
	}

	_, _ = arch.WriteString("}" + "\n")
	arch.Close()

	//Crear Archivo
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./GraficoColaMatriz.dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("GraficoColaMatriz.png", cmd, os.FileMode(mode))
	return true
}

func subgrafosCola(codigos []int, subg int, nop int) string{
	eti:=""
	var auxp []int
	auxp = make([]int,0)
	eti+=`subgraph cluster`+strconv.Itoa(subg)+`{style=invis;` + "\n"
	eti+=`edge[dir=fordware]` + "\n"
	for x:=0;x<len(codigos);x++{
		if x == 0{
			eti+="p"+strconv.Itoa(nop)+`[shape=box,color=blue,label="`+strconv.Itoa(codigos[x])+`"];` + "\n"
			auxp = append(auxp, nop)
			InicialesCola = append(InicialesCola, "p"+strconv.Itoa(nop))
			nop++
			Np++

		}else{
			eti+="p"+strconv.Itoa(nop)+`[shape=box,color=blue,label="`+strconv.Itoa(codigos[x])+`"];` + "\n"
			auxp = append(auxp, nop)
			nop++
			Np++
		}
	}
	for x:=0;x<len(auxp);x++{
		if x != (len(auxp)-1){
			eti+="p"+strconv.Itoa(auxp[x])+`->p`+strconv.Itoa(auxp[x+1]) + ";\n"
		}
	}
	eti+="}\n"
	return eti
}
