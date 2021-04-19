package Servidor

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Vertice struct{
	Nombre string
	Distancia int
	sig 	*Vertice
	adyacentes *ColaLA
}

type ColaLA struct {
	primero  *Vertice
	ultimo   *Vertice
}

type ListaAdyacencia struct{
	PosicionInicialRobot string
	Entrega string
	lista *ColaLA
}

func newNodoLA2(n string,i int) *Vertice{
	return &Vertice{n,i,nil, newCola2()}
}

func newCola2() *ColaLA{
	return &ColaLA{nil,nil}
}

func NewListaAdyacencia(p string, e string)* ListaAdyacencia{
	a:=newCola2()
	return &ListaAdyacencia{p,e,a}
}

func (this *ListaAdyacencia)getVertice(i string)*Vertice{
	e:=this.lista.primero
	for e!=nil{
		if e.Nombre==i{
			return e
		}
		e=e.sig
	}
	return nil
}

func (l *ColaLA) Vacio() bool  {
	return l.primero == nil
}

func insertarCLA(n string,t int, l *ColaLA) {
	var nuevo = newNodoLA2(n,t)
	if l.Vacio(){
		l.primero =nuevo
		l.ultimo  = nuevo
	} else {
		l.ultimo.sig = nuevo
		l.ultimo 	 = l.ultimo.sig
	}
}

func (this *ListaAdyacencia)Insertar(n string,i int){
	if this.getVertice(n)==nil{
		insertarCLA(n,i,this.lista)
	}else{
		fmt.Println("Elemento ya agregado")
	}
}

func (this *ListaAdyacencia)enlazar(a string, b string){
	var origen *Vertice
	var destino *Vertice
	origen = this.getVertice(a)
	destino = this.getVertice(b)
	if origen == nil || destino == nil{
		fmt.Println("No se encontro el vertice")
		return
	}
	insertarCLA(destino.Nombre,destino.Distancia,origen.adyacentes)
}

func contiene(buscando *ColaLA, elemento *Vertice)bool{
	e:=buscando.primero
	for e!=nil{
		if e==elemento{
			return true
		}
		e=e.sig
	}
	return false
}

func (this *ListaAdyacencia) dibujar(){
	arch, _ := os.Create("Grafo.dot")
	aux:=newCola2()
	var sc strings.Builder
	fmt.Fprintf(&sc,"digraph G{\n")
	fmt.Fprintf(&sc,`layout="circo"`+"\n")
	fmt.Fprintf(&sc,`edge [dir="both"]`+"\n")
	e:=this.lista.primero
	for e!=nil{
		palabras := strings.Split(e.Nombre," ")
		nodoaux := ""
		for j:=0;j<len(palabras);j++ {
			nodoaux = nodoaux + palabras[j]
		}
		if contiene(aux,e)==false{
			insertarCLA(e.Nombre,e.Distancia,aux)
			if e.Nombre == this.PosicionInicialRobot{
				fmt.Fprintf(&sc,"node"+nodoaux+`[label="`+e.Nombre+`",fillcolor="green",style="filled"]`+"\n")
			}else if e.Nombre == this.Entrega{
				fmt.Fprintf(&sc,"node"+nodoaux+`[label="`+e.Nombre+`",fillcolor="red",style="filled"]`+"\n")
			}else{
				fmt.Fprintf(&sc,"node"+nodoaux+`[label="`+e.Nombre+`"]`+"\n")
			}
		}
		temporal:=e.adyacentes.primero
		for temporal!=nil{
			palabras2 := strings.Split(temporal.Nombre," ")
			nodoaux2 := ""
			for j:=0;j<len(palabras2);j++ {
				nodoaux2 = nodoaux2 + palabras2[j]
			}

			fmt.Fprintf(&sc,"node"+nodoaux+" -> node"+nodoaux2+` [label="`+strconv.Itoa(temporal.Distancia)+`"]`+"\n")
			if contiene(aux, temporal)== false{
				insertarCLA(temporal.Nombre,temporal.Distancia,aux)
				if temporal.Nombre == this.PosicionInicialRobot{
					fmt.Fprintf(&sc,"node"+nodoaux2+`[label="`+temporal.Nombre+`",fillcolor="green",style="filled"]`+"\n")
				}else if temporal.Nombre == this.Entrega{
					fmt.Fprintf(&sc,"node"+nodoaux2+`[label="`+temporal.Nombre+`",fillcolor="red",style="filled"]`+"\n")
				}else{
					fmt.Fprintf(&sc,"node"+nodoaux2+`[label="`+temporal.Nombre+`"]`+"\n")
				}

			}
			temporal = temporal.sig
		}
		e=e.sig
	}
	fmt.Fprintf(&sc,"}")
	_, _ = arch.WriteString(sc.String())
	arch.Close()
	//Crear Archivo
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./Grafo.dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("Grafo.png", cmd, os.FileMode(mode))
}
