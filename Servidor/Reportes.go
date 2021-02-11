package Servidor

import (
	"fmt"
	"os"
	"strconv"
)

var ruta = "./Vector.dot"
var cabe string
var cant int
var canta int

func graficoCompleto(v []lista){
	var _, err = os.Stat(ruta)

	//Creando arhivo Si no Existe
	if os.IsNotExist(err){
		var archivo, _ = os.Create(ruta)
		if err != nil{
			fmt.Println(err.Error())
		}
		defer archivo.Close()
	}

	//Escribiendo Grafico En Archivo
	var archivo, err2 = os.OpenFile(ruta, os.O_RDWR, 0644)
	if err2 != nil{
		fmt.Println(err2.Error())
	}
	defer archivo.Close()

	//Estructura Del Archivo
	_, err = archivo.WriteString("Prueba")
	if err != nil{
		fmt.Println(err.Error())
	}

	//Guardar Cambios en Archivo
	cabe = ""
	cant = 1
	canta = 1
	num := 1
	_, err = archivo.WriteString("digraph grafo{" + "\n")
	_, err = archivo.WriteString("rankdir=LR;" + "\n")
	_, err = archivo.WriteString("node [shape = record, style=filled];" + "\n")
	for i := 0; i < len(v); i++{
		_, err = archivo.WriteString("nodo" + strconv.Itoa(cant) + `[label="`+ strconv.Itoa(num) + `"];` + "\n")
		aux := v[i].primero
		for aux != nil{
			cant++
			_, err = archivo.WriteString("nodo" + strconv.Itoa(cant) + `[label="`+ aux.tienda.nombre + `"];` + "\n")
			aux = aux.sig
		}
		num++
		cant++
	}

	num = 1
	cant = 1
	for i := 0; i < len(v); i++{
		cant = canta
		aux := v[i].primero
		for aux != nil{
			cabe = "nodo" + strconv.Itoa(canta) + " -> nodo" + strconv.Itoa(canta + 1) + "; \n"
			_, err = archivo.WriteString(cabe)
			canta++
			aux = aux.sig
		}
		canta++
		if i != len(v) - 1{
			cabe = "nodo" + strconv.Itoa(cant) + " -> nodo" + strconv.Itoa(canta) + "; \n"
			_, err = archivo.WriteString(cabe)
		}
		num++
	}
	_, err = archivo.WriteString("}" + "\n")

	//Guardar Cambios
	err = archivo.Sync()
	if err != nil{
		fmt.Println(err.Error())
	}

	//Abrir Archivo Creado (No Funciona Aun)
	/*grafo := "dot -Tpng Vector.dot -o grafo.png"
	cmd := exec.Command(grafo)
	_, err = cmd.Output()*/
}