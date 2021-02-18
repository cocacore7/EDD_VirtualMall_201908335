package Servidor

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

var cant int
var canta int

//Recursividad Grafo
func reportes(i int, f int, n int){
	//Metodo Recursivo
	if i+5 <= len(vec){
		f = f + 5
		escribir(i,f,n)
	}else if i+4 <= len(vec){
		f = f + 4
		escribir(i,f,n)
	}else if i+3 <= len(vec){
		f = f + 3
		escribir(i,f,n)
	}else if i+2 <= len(vec){
		f = f + 2
		escribir(i,f,n)
	}else if i+1 <= len(vec){
		f = f + 1
		escribir(i,f,n)
	} else{
		return
	}
}

//Crear Grafo
func escribir(i int, f int, n int){
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
		aux := vec[inicio].primero
		for aux != nil{
			cant++
			_, _ = arch.WriteString("nodo" + strconv.Itoa(cant) + `[label="Nombre: `+ aux.tienda.nombre + ". Contacto: " + aux.tienda.contacto + `",fillcolor=red];` + "\n")
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
		aux := vec[inicio].primero
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
	reportes(i,f,n)
}

func retorno(in int, fi int) []byte{
	var reg Datos
	reg = data
	for i:=0;i<len(data.Datos);i++{
		for j:=0;j<len(data.Datos[i].Departamentos);j++{
			if in+5 <= len(vec){
				fi = fi + 5
			}else if in+4 <= len(vec){
				fi = fi + 4
			}else if in+3 <= len(vec){
				fi = fi + 3
			}else if in+2 <= len(vec){
				fi = fi + 2
			}else if in+1 <= len(vec){
				fi = fi + 1
			}
			reg.Datos[i].Departamentos[j].Tiendas = obtenerT(in, fi)
			in = fi
		}
	}
	arch, _ := os.Create("Salida.json")
	crearJson, _ := json.MarshalIndent(reg,"","    ")
	arch.WriteString(string(crearJson))
	//arch.Write(crearJson)
	arch.Close()
	return crearJson
}

func obtenerT(i int, f int) []Tiendas{
	var t []Tiendas
	var aux Tiendas
	for i<f{
		a := vec[i].primero
		for a != nil{
			aux.Tiendas = a.tienda.nombre
			aux.Descripcion = a.tienda.descripcion
			aux.Contacto = a.tienda.contacto
			aux.Calificacion = a.tienda.calif
			t = append(t, aux)
			a = a.sig
		}
		i++
	}
	return t
}
