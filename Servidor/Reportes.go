package Servidor

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

var cant int
var canta int

func reportes(v []lista, i int, f int, n int){
	//Metodo Recursivo
	if i+5 <= len(v){
		f = f + 5
		escribir(i,f,n,v)
	}else if i+4 <= len(v){
		f = f + 4
		escribir(i,f,n,v)
	}else if i+3 <= len(v){
		f = f + 3
		escribir(i,f,n,v)
	}else if i+2 <= len(v){
		f = f + 2
		escribir(i,f,n,v)
	}else if i+1 <= len(v){
		f = f + 1
		escribir(i,f,n,v)
	} else{
		return
	}
}

func escribir(i int, f int, n int, v []lista){
	arch, _ := os.Create("archivo" + strconv.Itoa(n) + ".dot")

	//Guardar Cambios en Archivo
	cant = 1
	canta = 1
	_, _ = arch.WriteString("digraph G{" + "\n")
	_, _ = arch.WriteString(`graph[splines="ortho"];` + "\n")
	_, _ = arch.WriteString("node [shape = record, style=filled];" + "\n")

	cant = i+1
	canta = i+1
	num := i+1
	inicio := i
	final := f
	//Nombres y Colores Nodos
	for inicio < final{
		_, _ = arch.WriteString("nodo" + strconv.Itoa(cant) + `[label="`+ strconv.Itoa(num) + `",fillcolor=green];` + "\n")
		aux := v[inicio].primero
		for aux != nil{
			cant++
			_, _ = arch.WriteString("nodo" + strconv.Itoa(cant) + `[label="Nombre: `+ aux.tienda.nombre + ". Contacto: " + aux.tienda.contacto + ". Descripcion: " + aux.tienda.descripcion + `",fillcolor=red];` + "\n")
			aux = aux.sig
		}
		num++
		cant++
		inicio++
	}
	num = i+1
	cant = i+1
	inicio = i
	final = f
	//Conexiones Nodos
	for inicio < final{
		cant = canta
		aux := v[inicio].primero
		for aux != nil{
			_, _ = arch.WriteString("nodo" + strconv.Itoa(canta) + " -> nodo" + strconv.Itoa(canta + 1)  + "; \n")
			canta++
			aux = aux.sig
		}
		canta++
		if inicio != f - 1{
			_, _ = arch.WriteString("nodo" + strconv.Itoa(cant) + " -> nodo" + strconv.Itoa(canta)  + "; \n")
			_, _ = arch.WriteString("nodo" + strconv.Itoa(canta) + " -> nodo" + strconv.Itoa(cant) + "; \n")
		}
		num++
		inicio++
	}

	_, _ = arch.WriteString("}" + "\n")
	arch.Close()

	//Crear Archivo
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./archivo" + strconv.Itoa(n) + ".dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("outfile" + strconv.Itoa(n) + ".png", cmd, os.FileMode(mode))
	i = f
	n++
	//Recursividad
	reportes(v,i,f,n)
}
